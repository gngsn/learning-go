# gRPC

<small>๐ย <a href="https://grpc.io/docs/languages/go/basics/">gRPC Basics Tutorial</a></small>

<small>์ด ํ๋ก์ ํธ๋ gRPC ๊ณต์ ๋ฌธ์์ ํํ ๋ฆฌ์ผ์ ๋ณํํ์ฌ ์์ฑํ์ต๋๋ค.</small>

<br /><br /><br />

## 1. Project ์์ฑ

``` bash
$ mkdir grpc
$ cd grpc

$ go get -u google.golang.org/grpc
$ go get -u github.com/golang/protobuf/protoc-gen-go
```

<br />

### protoc PATH ์ถ๊ฐ

``` bash
$ export PATH="$PATH:$(go env GOPATH)/bin"
```

<br /><br /><br />

## 2. proto ์ ์

`routeguide` ๋๋ ํฐ๋ฆฌ๋ฅผ ์์ฑํ ํ, `route-guide.proto` ํ์ผ์ ์์ฑํจ

message์ service๋ฅผ ์ ์ํ๋ค.

<br /><br />

### message ์ ์

``` protobuf
// ์์น ์ขํ ๊ฐ
message Point {
    int32 latitude = 1;
    int32 longitude = 2;
}
// ํน์  ์์ญ์ ์์น ์ขํ ๊ฐ ๋ฒ์
message Rectangle {
  Point lo = 1;
  Point hi = 2;
}
// ์์น - ์ฅ์ ์ด๋ฆ & ์ขํ
message Feature {
    string name = 1;
    Point location = 2;
}

// ์ด ์์น ์, ์์น์ ํด๋นํ๋ ์ฅ์ ๊ฐ์, ์ฐ์๋ ์์น๋ค์ ์ด ๊ฑฐ๋ฆฌ
message RouteSummary {
  int32 point_count = 1;
  int32 feature_count = 2;
  int32 distance = 3;
  int32 elapsed_time = 4;
}

// ์์น๊ฐ๊ณผ ๋ฉ์ธ์ง
message RouteNote {
  Point location = 1;
  string message = 2;
}
```

<br /><br />

### Service ์ ์

``` protobuf
service RouteGuide {
    /*
        Simple RPC
        : point ๋ฐ์์ Feature ๋ฐํ
    */
    rpc GetFeature(Point) returns (Feature) {}

    /*
        Server-side Streaming RPC 
        : Rectangle ๋ฒ์ ๋ด์ ์๋ ๋ชจ๋  Feature๋ฅผ stream์ผ๋ก ๋ฐํ.
    */
    rpc ListFeatures(Rectangle) returns (stream Feature) {}

    /*
        Client-side Streaming RPC
        : Point stream์ ๋ฐ์์ RouteSummary ๋ฐํ
    */
    rpc RecordRoute(stream Point) returns (RouteSummary) {}

    /*
        Bidirectional Streaming RPC
        : RouteNote stream์ ๋ฐ์์ RouteNote stream์ ๋ฐํ
    */
    rpc RouteChat(stream RouteNote) returns (stream RouteNote) {}
}
```

<br />

#### Server-side Streaming

Client๊ฐ Response ๋ฐ๊ณ  ๋๋ด๋ ๊ฒ ์๋๋ผ, Stream์ ๋ฐ์ผ๋ฉด์ ๋ฉ์ธ์ง๊ฐ ๋ ์์ ๋๊น์ง ๊ณ์ ๊ตฌ๋.

ํ๋ฒ์ ํฐ ๋ฐ์ดํฐ๋ฅผ ๋ฆฌํดํ๊ฒ ๋๋ฉด client๋ ๋ฐ์ดํฐ๋ฅผ ๋ฐ๊ธฐ ๊น์ง ๊ณ์ blocking์ด ๋์ด์์ด์ ๋ค๋ฅธ ์์๋ค์ ํ์ง ๋ชปํ๊ฒ ๋จ

<br />

#### Client-side Streaming

Client๊ฐ stream ๋ฉ์์ง๋ฅผ ์์ฑํด์ ์๋ฒ์ ์์ฒญํ๊ณ  ์ ๊ณต๋ ์คํธ๋ฆผ์ ๋ค์ ์ฌ์ฉ

ํด๋ผ์ด์ธํธ๋ ๋ฉ์์ง ์์ฑ์ ์๋ฃํ ํ ์๋ฒ๊ฐ ๋ฉ์์ง๋ฅผ ๋ชจ๋ ์ฝ๊ณ  ์๋ต์ ๋ฐํํ  ๋๊น์ง ๊ธฐ๋ค๋ฆผ

<br />

#### Bidirectional Streaming

client์ server๊ฐ ๋๋ค stream๋ฐฉ์์ผ๋ก ์๋ก ์ฃผ๊ณ  ๋ฐ๋ ๋ฐฉ์

๊ฐ ์คํธ๋ฆผ์ ๋ฉ์์ง ์์๊ฐ ์ ์ง

2๊ฐ์ stream์ ๊ฐ๊ฐ ๋๋ฆฝ์  - server๋ client๊ฐ stream์ผ๋ก request๋ฅผ ๋ค ๋ณด๋ผ๋๊น์ง ๊ธฐ๋ค๋ฆฌ๊ณ  ๋์ response๋ฅผ ์ฃผ๋์ง request๊ฐ ์ฌ ๋๋ง๋ค ๋ฐ๋ก response๋ฅผ ๋ณด๋ผ ๊ฒ์ธ์ง ๊ฒฐ์ ํ  ์ ์์

<br /><br />

### Compile

``` bash
$ protoc --go_out=. --go_opt=paths=source_relative \
					--go-grpc_out=. --go-grpc_opt=paths=source_relative 
					routeguide/route-guide.proto
```

<br />

####  `--go_out`, `--go-grpc_out`

: ๊ฐ๊ฐ `<file>.pb.go`, `<file>_grpc.pb.go` ํ์ผ์ ์ ์ฅํ  ์์น๋ฅผ ์ง์ 

<br />

#### `--go_opt` , `--go-grpc_opt`

: ๊ฐ๊ฐ `<file>.pb.go`, `<file>_grpc.pb.go`  ํ์ผ์ ์ ์ฅํ  ๋์ ์ต์.

<br />

โ๏ธ `paths=source_relative` :ํ๋๊ทธ๊ฐ ์ง์ ๋ ๊ฒฝ์ฐ ์ถ๋ ฅ ํ์ผ์ ์๋ ฅ **ํ์ผ๊ณผ ๋์ผํ ์๋ ๋๋ ํฐ๋ฆฌ์ ๋ฐฐ์น**๋จ.

โ๏ธ `paths=import` : proto ํ์ผ์ import ๊ฒฝ๋ก๋ก ์ ์ฅ๋จ.  <small>EX) ๋ช๋ น์ด๋ฅผ ์๋ ฅํ ์ง์ ์์ `github.com/gngsn/learning-go/apps/grpc/routeguide` ๋๋ ํฐ๋ฆฌ๋ฅผ ์์ฑํด์ ์ ์ฅ. ์์์ธ๋ฏ..</small>

โ๏ธ ` module=$PREFIX` : proto ํ์ผ์ import ๊ฒฝ๋ก๊ฐ ์ง์ ๋์ง๋ง, ์ง์ ๋ ๋๋ ํฐ๋ฆฌ ์ ๋์ฌ๊ฐ ์ถ๋ ฅ ํ์ผ ์ด๋ฆ์์ ์ ๊ฑฐ. <small>EX) ๋ช๋ น์ด๋ฅผ ์๋ ฅํ ์ง์ ์์ `github.com/gngsn/apps/grpc/routeguide` ๋๋ ํฐ๋ฆฌ ๊ฒฝ๋ก์์ `module=github.com/gngsn`ย ๋ฅผ ์ค์ ํ๋ฉด `/apps/grpc/routeguide`์ ์์ฑํด์ ์ ์ฅ.</small>

<br /><br />

์์ ๋ช๋ น์ด๋ฅผ ์๋ ฅํ๋ฉด ์๋์ ๋ ํ์ผ์ด ์์ฑ๋จ. 

#### \<File\>.pb.go

๋ฐ์ดํฐ๋ฅผ ์ฑ์ฐ๊ณ , serialize์ํค๊ณ , ์์ฒญ๊ณผ ์๋ต ๋ฉ์ธ์ง๋ฅผ ํ์ํ๊ธฐ ์ํ **๋ชจ๋  ํ๋กํ ์ฝ ๋ฒํผ ์ฝ๋**๋ฅผ ํฌํจ

๊ฐ ๋ฉ์ธ์ง ํ์์ ๋ํ `Reset()`, `String()`, `ProtoMessage()`, `ProtoReflect()`, `Descriptor()`, ๊ทธ๋ฆฌ๊ณ  ๋ชจ๋  ํ๋์ ๋ํ `Getter()` ๋ฅผ ์์ฑ.

<br />

#### \<File\>_grpc.pb.go

์๋ฒ๊ฐ ๊ตฌํํด์ผํ  **์ธํฐํ์ด์ค ์ ํ ๋ฉ์๋**์ Client๊ฐ ํธ์ถํ  ์ ์๋ ๋ฉ์๋์ **์ธํฐํ์ด์ค ์ ํ(๋๋ ์คํ)** ๋ฅผ ์์ฑ.

input์ stream์ด ์์ผ๋ฉด `Send()` ๊ฐ ๊ตฌํ๋๊ณ , return์ stream์ด ์์ผ๋ฉด `Revc()` ๊ฐ ๊ตฌํ๋์ด ์์.

Client์ Server์์ ์ฌ์ฉํ  interface๋ฅผ ๋ฐ๋ก ์์ฑํจ.

<br />

EX) proto์์ ์ ์ํ `RouteChat` ์ ํ์ธํด๋ณด์ ~

``` protobuf
rpc RouteChat(stream RouteNote) returns (stream RouteNote) {}
```

<br />

**Client**

``` go
type RouteGuideClient interface {
  // ...
  RouteChat(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RouteChatClient, error)
}

func (c *routeGuideClient) RouteChat(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RouteChatClient, error) { /* ... */ }

type RouteGuide_RouteChatClient interface {
	Send(*RouteNote) error
	Recv() (*RouteNote, error)
	grpc.ClientStream
}

type routeGuideRouteChatClient struct {
  grpc.ClientStream
}

func (x *routeGuideRouteChatClient) Send(m *RouteNote) error { /* ... */ }

func (x *routeGuideRouteChatClient) Recv() (*RouteNote, error) { /* ... */ }

```

<br />

**Server**

``` go
type RouteGuideServer interface {
  // ...
  RouteChat(RouteGuide_RouteChatServer) error
}
func (UnimplementedRouteGuideServer) RouteChat(RouteGuide_RouteChatServer) error { /* ... */ }

func _RouteGuide_RouteChat_Handler(srv interface{}, stream grpc.ServerStream) error { /* ... */ }

type RouteGuide_RouteChatServer interface {
	Send(*RouteNote) error
	Recv() (*RouteNote, error)
	grpc.ServerStream
}

type routeGuideRouteChatServer struct {
	grpc.ServerStream
}

func (x *routeGuideRouteChatServer) Send(m *RouteNote) error { /* ... */ }

func (x *routeGuideRouteChatServer) Recv() (*RouteNote, error) { /* ... */ }

var RouteGuide_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "RouteGuide",
	HandlerType: (*RouteGuideServer)(nil),
	Methods: []grpc.MethodDesc{
    //...
		,{
			StreamName:    "RouteChat",
			Handler:       _RouteGuide_RouteChat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "routeguide/route-guide.proto",
}
```

์์ ๊ฐ์ด ์์ฑ๋๋ค.

<br /><br />

## 3. Create Server

<br />

### ์๋น์ค ๊ตฌํ

์ ์ํ ์๋น์ค๋ค์ด ๋์ํ๊ฒ๋ ๋ด๋ถ ๋ก์ง์ ์ ์ด์ฃผ์

``` go
package server

import (
	"context"

	pb "github.com/gngsn/learning-go/apps/grpc/routeguide"
)

type routeGuideServer struct { /* ... */ }

func (s *routeGuideServer) GetFeature(ctx context.Context, point *pb.Point) { /* ... */ }

func (s *routeGuideServer) ListFeatures(rect *pb.Rectangle, stream pb.RouteGuide_ListFeaturesServer) error { /* ... */ }

func (s *routeGuideServer) RecordRoute(stream pb.RouteGuide_RecordRouteServer) error { /* ... */ }

func (s *routeGuideServer) RouteChat(stream pb.RouteGuide_RouteChatServer) error { /* ... */ }
```

<br />

#### sream 

service์ input์ด๋ output์ stream์ ๋ฐ๊ฑฐ๋ ์ฃผ๊ฒ ์ ์ํ๋ค๋ฉด, `stream pb.<Service>_<Method>Server` ๋ฅผ ์ธ์๋ก ๋ฐ๊ฒ ๋๋ค.

๋ฐ์ดํฐ๋ฅผ ์ ๋ฌํ  ๋์๋ `sream.send(Data)` ๋ฅผ ํตํด ์ ๋ฌ.

๋ฐ์ดํฐ๋ฅผ ๋ฐ์ ๋์๋ `Data, err := stream.Recv()` ์ ๊ฐ์ด ๋ฐ์ ์ ์๋ค.

Data๋ proto์ ์ ์๋ ๋ฐ์ดํฐ ํํ ~

<br /><br />

### ์๋ฒ ์ค์ 

์๋ฒ๋ฅผ ์ค์ ํด๋ณด์

``` go
// ์๋น์ค์ ํด๋นํ๋ Server struct ์์ฑ
type routeGuideServer struct {
	pb.UnimplementedRouteGuideServer
	savedFeatures []*pb.Feature

	mu         sync.Mutex
	routeNotes map[string][]*pb.RouteNote
}

// Server struct ์ธ์คํด์ค ์์ฑ
func newServer() *routeGuideServer {
	s := &routeGuideServer{routeNotes: make(map[string][]*pb.RouteNote)}
	s.loadFeatures(jsonDBFile)
	return s
}

func main() {
  // tcp๋ก net.Listen
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

  // grpc ์๋ฒ ์์ฑ ํ Serve
	grpcServer := grpc.NewServer()
	pb.RegisterRouteGuideServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
```

<br /><br />

## 4. Run Server

``` bash
$ go run server.go
```

<br /><br />

## 5. Test Server

<small>๐ <a href="https://pkg.go.dev/testing">Goย Testing</a></small>

์ ์ฌ์ค ํํ ๋ฆฌ์ผ์์๋ `client/client.go` ๋ฅผ ์ฌ์ฉํ์ง๋ง, Test ๊ธฐ๋ฐ ํ๋ก์ ํธ๋ฅผ ๋ง๋ค๊ณ  ์ถ์ด์ ์ง๊ธ๊น์ง ๋ฐฐ์์จ ๊ฑธ ๋ณต์ต๋ ํด๋ณผ๊ฒธ ํ์คํธ ์ฝ๋๋ฅผ ์งฐ๋ค.

``` go
// connection ์์ฑ
func getConn() *grpc.ClientConn {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	return conn
}

// Unit Test
func TestGrpc(t *testing.T) {
	conn := getConn()
	defer conn.Close()
	client := pb.NewRouteGuideClient(conn)
	
	// ์๋: 409146138, ๊ฒฝ๋: -746188906์ธ ์์น์ ์ฅ์ ์ฐพ๊ธฐ
	t.Run("T1-GetFeature", func(t *testing.T) {
		printFeature(client, &pb.Point{Latitude: 409146138, Longitude: -746188906})
	})

	/*
		์๋: [400000000, 420000000], 
		๊ฒฝ๋: [-750000000, -730000000] ์์น์ ์ฅ์ ์ฐพ๊ธฐ
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
```

<br />

### All Test

``` bash
$ go test -run TestGrpc 
```

<br /><br />

### Subtests

``` bash
$ go test -run TestGrpc/T1-GetFeature
// or
$ go test -run TestGrpc/T1
```

<br /><br />

### Test Result

<br />

#### GetFeature

``` bash
$ go test -run TestGrpc/T1
2021/11/22 22:53:18 Getting feature for point (409146138, -746188906)
2021/11/22 22:53:18 name:"Berkshire Valley Management Area Trail, Jefferson, NJ, USA" location:{latitude:409146138 longitude:-746188906}
PASS
ok      github.com/gngsn/learning-go/apps/grpc  0.228s
```

<br />

#### ListFeatures

``` bash
$ go test -run TestGrpc/T2
2021/11/22 22:53:42 Looking for features within lo:{latitude:400000000 longitude:-750000000} hi:{latitude:420000000 longitude:-730000000}
2021/11/22 22:53:42 Feature: name: "Patriots Path, Mendham, NJ 07945, USA", point:(407838351, -746143763)
2021/11/22 22:53:42 Feature: name: "101 New Jersey 10, Whippany, NJ 07981, USA", point:(408122808, -743999179)
2021/11/22 22:53:42 Feature: name: "U.S. 6, Shohola, PA 18458, USA", point:(413628156, -749015468)
...
PASS
ok      github.com/gngsn/learning-go/apps/grpc  0.164s
```

<br />

#### RecordRoute

``` bash
$ go test -run TestGrpc/T3
2021/11/22 22:54:30 Traversing 80 points.
2021/11/22 22:54:30 Route summary: point_count:80 distance:828725706
PASS
ok      github.com/gngsn/learning-go/apps/grpc  0.297s
```

<br />

#### RouteChat

``` bash
$ go test -run TestGrpc/T4
2021/11/22 22:54:44 Got message First message at point(0, 1)
2021/11/22 22:54:44 Got message Fourth message at point(0, 1)
2021/11/22 22:54:44 Got message First message at point(0, 1)
2021/11/22 22:54:44 Got message Second message at point(0, 2)
2021/11/22 22:54:44 Got message Fifth message at point(0, 2)
2021/11/22 22:54:44 Got message Second message at point(0, 2)
...
PASS
ok      github.com/gngsn/learning-go/apps/grpc  0.165s
```

<br /><br /><br />

์ฌ๋ฐ๋ค

<br /><br />