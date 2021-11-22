package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"sync"
	"time"

	modules "github.com/gngsn/learning-go/apps/grpc/modules"
	pb "github.com/gngsn/learning-go/apps/grpc/routeguide"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

var (
	jsonDBFile = "./public/dump.json"
	port       = flag.Int("port", 50051, "The server port")
)

type routeGuideServer struct {
	pb.UnimplementedRouteGuideServer
	savedFeatures []*pb.Feature

	mu         sync.Mutex
	routeNotes map[string][]*pb.RouteNote
}

func (s *routeGuideServer) GetFeature(ctx context.Context, point *pb.Point) (*pb.Feature, error) {
	for _, feature := range s.savedFeatures {
		if proto.Equal(feature.Location, point) {
			return feature, nil
		}
	}

	return &pb.Feature{Location: point}, nil
}

func (s *routeGuideServer) ListFeatures(rect *pb.Rectangle, stream pb.RouteGuide_ListFeaturesServer) error {
	for _, feature := range s.savedFeatures {
		if modules.InRange(feature.Location, rect) {
			if err := stream.Send(feature); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *routeGuideServer) RecordRoute(stream pb.RouteGuide_RecordRouteServer) error {
  var pointCount, featureCount, distance int32
  var lastPoint *pb.Point
  startTime := time.Now()
  for {
    point, err := stream.Recv()
    if err == io.EOF {
      endTime := time.Now()
      return stream.SendAndClose(&pb.RouteSummary{
        PointCount:   pointCount,
        FeatureCount: featureCount,
        Distance:     distance,
        ElapsedTime:  int32(endTime.Sub(startTime).Seconds()),
      })
    }
    if err != nil {
      return err
    }
    pointCount++
    for _, feature := range s.savedFeatures {
      if proto.Equal(feature.Location, point) {
        featureCount++
      }
    }
    if lastPoint != nil {
      distance += modules.CalcDistance(lastPoint, point)
    }
    lastPoint = point
  }
}

func serialize(point *pb.Point) string {
	return fmt.Sprintf("%d %d", point.Latitude, point.Longitude)
}

func (s *routeGuideServer) RouteChat(stream pb.RouteGuide_RouteChatServer) error {
  for {
    in, err := stream.Recv()
    if err == io.EOF {
      return nil
    }
    if err != nil {
      return err
    }

	log.Println("in message : ", in.Message)
    key := serialize(in.Location)

	s.mu.Lock()
	s.routeNotes[key] = append(s.routeNotes[key], in)
	rn := make([]*pb.RouteNote, len(s.routeNotes[key]))
	copy(rn, s.routeNotes[key])
	s.mu.Unlock()

    for _, note := range s.routeNotes[key] {
		log.Println("note : ", note)
		if err := stream.Send(note); err != nil {
			return err
		}
    }
  }
}

func newServer() *routeGuideServer {
	s := &routeGuideServer{routeNotes: make(map[string][]*pb.RouteNote)}
	s.loadFeatures(jsonDBFile)
	return s
}

func (s *routeGuideServer) loadFeatures(filePath string) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to load: %v", err)
	}
	if err = json.Unmarshal(data, &s.savedFeatures); err != nil {
		log.Fatalf("Failed to load: %v", err)
	}
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterRouteGuideServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}