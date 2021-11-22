package main

import (
	"context"
	"log"
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

func getConn() *grpc.ClientConn {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	
	return conn
}

func TestGetFeature(t *testing.T) {
	conn := getConn()
	defer conn.Close()
	client := pb.NewRouteGuideClient(conn)
	
	// 위도: 409146138, 경도: -746188906인 위치의 장소 찾기
	t.Run("T1-GetFeature", func(t *testing.T) {
		printFeature(client, &pb.Point{Latitude: 409146138, Longitude: -746188906})
	})
	
	// 위도: 0, 경도: 0인 위치의 장소 찾기
	t.Run("T2-GetFeature", func(t *testing.T) {
		printFeature(client, &pb.Point{Latitude: 0, Longitude: 0})
	})

	t.Run("PrintSomething", func(t *testing.T) {
		log.Print("로그 테스트")
	})
}