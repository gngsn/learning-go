package main

import (
	"context"
	"io"
	"log"
	"math/rand"
	"testing"
	"time"

	pb "github.com/gngsn/learning-go/apps/grpc/routeguide"
	"google.golang.org/grpc"
)

var (
	serverAddr = "localhost:50051"
)

func printFeature(client pb.RouteGuideClient, point *pb.Point) {
	log.Printf("Getting feature for point (%d, %d)", point.Latitude, point.Longitude)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	feature, err := client.GetFeature(ctx, point)
	if err != nil {
		log.Fatalf("%v.GetFeatures(_) = _, %v: ", client, err)
	}
	log.Print(feature)
}

func printFeatures(client pb.RouteGuideClient, rect *pb.Rectangle) {
	log.Printf("Looking for features within %v", rect)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.ListFeatures(ctx, rect)
	if err != nil {
		log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
	}
	for {
		feature, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
		}
		log.Printf("Feature: name: %q, point:(%v, %v)", feature.GetName(),
			feature.GetLocation().GetLatitude(), feature.GetLocation().GetLongitude())
	}
}

func randomPoint(r *rand.Rand) *pb.Point {
	lat := (r.Int31n(180) - 90) * 1e7
	long := (r.Int31n(360) - 180) * 1e7
	return &pb.Point{Latitude: lat, Longitude: long}
}

func runRecordRoute(client pb.RouteGuideClient) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	pointCount := int(r.Int31n(100)) + 2
	var points []*pb.Point

	for i := 0; i < pointCount; i++ {
		points = append(points, randomPoint(r))
	}
	log.Printf("Traversing %d points.", len(points))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.RecordRoute(ctx)
	if err != nil {
		log.Fatalf("%v.RecordRoute(_) = _, %v", client, err)
	}

	for _, point := range points {
		if err := stream.Send(point); err != nil {
			log.Fatalf("%v.Send(%v) = %v", stream, point, err)
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
	}
	log.Printf("Route summary: %v", reply)
}

func runRouteChat(client pb.RouteGuideClient) {
	notes := []*pb.RouteNote{
		{Location: &pb.Point{Latitude: 0, Longitude: 1}, Message: "First message"},
		{Location: &pb.Point{Latitude: 0, Longitude: 2}, Message: "Second message"},
		{Location: &pb.Point{Latitude: 0, Longitude: 3}, Message: "Third message"},
		{Location: &pb.Point{Latitude: 0, Longitude: 1}, Message: "Fourth message"},
		{Location: &pb.Point{Latitude: 0, Longitude: 2}, Message: "Fifth message"},
		{Location: &pb.Point{Latitude: 0, Longitude: 3}, Message: "Sixth message"},
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.RouteChat(ctx)
	if err != nil {
		log.Fatalf("%v.RouteChat(_) = _, %v", client, err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a note : %v", err)
			}
			log.Printf("Got message %s at point(%d, %d)", in.Message, in.Location.Latitude, in.Location.Longitude)
		}
	}()
	for _, note := range notes {
		if err := stream.Send(note); err != nil {
			log.Fatalf("Failed to send a note: %v", err)
		}
	}
	stream.CloseSend()
	<-waitc
}


func getConn() *grpc.ClientConn {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	
	return conn
}

func TestGrpc(t *testing.T) {
	conn := getConn()
	defer conn.Close()
	client := pb.NewRouteGuideClient(conn)
	
	// ??????: 409146138, ??????: -746188906??? ????????? ?????? ??????
	t.Run("T1-GetFeature", func(t *testing.T) {
		printFeature(client, &pb.Point{Latitude: 409146138, Longitude: -746188906})
	})

	/*
		??????: [400000000, 420000000], 
		??????: [-750000000, -730000000] ????????? ?????? ??????
	*/
	t.Run("T2-GetFeatures", func(t *testing.T) {
		printFeatures(client, &pb.Rectangle{
			Lo: &pb.Point{Latitude: 400000000, Longitude: -750000000},
			Hi: &pb.Point{Latitude: 420000000, Longitude: -730000000},
		})
	})

	t.Run("T3-RecordRoute", func(t *testing.T) {
		runRecordRoute(client)
	})

	t.Run("T4-RouteChat", func(t *testing.T) {
		runRouteChat(client)
	})
}
