# gRPC

<small>🔗 <a href="https://grpc.io/docs/languages/go/basics/">gRPC Basics Tutorial</a></small>

<small>이 프로젝트는 gRPC 공식 문서의 튜토리얼을 변형하여 작성했습니다.</small>

<br /><br /><br />

## 1. Project 생성

``` bash
$ mkdir grpc
$ cd grpc

$ go get -u google.golang.org/grpc
$ go get -u github.com/golang/protobuf/protoc-gen-go
```

<br />

### protoc PATH 추가

``` bash
$ export PATH="$PATH:$(go env GOPATH)/bin"
```

<br /><br /><br />

## 2. proto 정의

`routeguide` 디렉터리를 생성한 후, `route-guide.proto` 파일을 생성함

message와 service를 정의한다.

<br /><br />

### message 정의

``` protobuf
// 위치 좌표 값
message Point {
    int32 latitude = 1;
    int32 longitude = 2;
}
// 특정 영역의 위치 좌표 값 범위
message Rectangle {
  Point lo = 1;
  Point hi = 2;
}
// 위치 - 장소 이름 & 좌표
message Feature {
    string name = 1;
    Point location = 2;
}

// 총 위치 수, 위치에 해당하는 장소 개수, 연속된 위치들의 총 거리
message RouteSummary {
  int32 point_count = 1;
  int32 feature_count = 2;
  int32 distance = 3;
  int32 elapsed_time = 4;
}

// 위치값과 메세지
message RouteNote {
  Point location = 1;
  string message = 2;
}
```

<br /><br />

### Service 정의

``` protobuf
service RouteGuide {
    /*
        Simple RPC
        : point 받아서 Feature 반환
    */
    rpc GetFeature(Point) returns (Feature) {}

    /*
        Server-side Streaming RPC 
        : Rectangle 범위 내에 있는 모든 Feature를 stream으로 반환.
    */
    rpc ListFeatures(Rectangle) returns (stream Feature) {}

    /*
        Client-side Streaming RPC
        : Point stream을 받아서 RouteSummary 반환
    */
    rpc RecordRoute(stream Point) returns (RouteSummary) {}

    /*
        Bidirectional Streaming RPC
        : RouteNote stream을 받아서 RouteNote stream을 반환
    */
    rpc RouteChat(stream RouteNote) returns (stream RouteNote) {}
}
```

<br />

#### Server-side Streaming

Client가 Response 받고 끝내는 게 아니라, Stream을 받으면서 메세지가 더 없을 때까지 계속 구독.

한번에 큰 데이터를 리턴하게 되면 client는 데이터를 받기 까지 계속 blocking이 되어있어서 다른 작업들을 하지 못하게 됨

<br />

#### Client-side Streaming

Client가 stream 메시지를 작성해서 서버에 요청하고 제공된 스트림을 다시 사용

클라이언트는 메시지 작성을 완료한 후 서버가 메시지를 모두 읽고 응답을 반환할 때까지 기다림

<br />

#### Bidirectional Streaming

client와 server가 둘다 stream방식으로 서로 주고 받는 방식

각 스트림의 메시지 순서가 유지

2개의 stream은 각각 독립적 - server는 client가 stream으로 request를 다 보낼때까지 기다리고 나서 response를 주던지 request가 올 때마다 바로 response를 보낼 것인지 결정할 수 있음

<br /><br />

### Compile

``` bash
$ protoc --go_out=. --go_opt=paths=source_relative \
					--go-grpc_out=. --go-grpc_opt=paths=source_relative 
					routeguide/route-guide.proto
```

<br />

####  `--go_out`, `--go-grpc_out`

: 각각 `<file>.pb.go`, `<file>_grpc.pb.go` 파일을 저장할 위치를 지정

<br />

#### `--go_opt` , `--go-grpc_opt`

: 각각 `<file>.pb.go`, `<file>_grpc.pb.go`  파일을 저장할 때의 옵션.

<br />

✔️ `paths=source_relative` :플래그가 지정된 경우 출력 파일은 입력 **파일과 동일한 상대 디렉터리에 배치**됨.

✔️ `paths=import` : proto 파일의 import 경로로 저장됨.  <small>EX) 명령어를 입력한 지점에서 `github.com/gngsn/learning-go/apps/grpc/routeguide` 디렉터리를 생성해서 저장. 잘안쓸듯..</small>

✔️ ` module=$PREFIX` : proto 파일의 import 경로가 지정되지만, 지정된 디렉터리 접두사가 출력 파일 이름에서 제거. <small>EX) 명령어를 입력한 지점에서 `github.com/gngsn/apps/grpc/routeguide` 디렉터리 경로에서 `module=github.com/gngsn` 를 설정하면 `/apps/grpc/routeguide`에 생성해서 저장.</small>

<br /><br />

위의 명령어를 입력하면 아래의 두 파일이 생성됨. 

#### \<File\>.pb.go

데이터를 채우고, serialize시키고, 요청과 응답 메세지를 회수하기 위한 **모든 프로토콜 버퍼 코드**를 포함

각 메세지 타입에 대한 `Reset()`, `String()`, `ProtoMessage()`, `ProtoReflect()`, `Descriptor()`, 그리고 모든 필드에 대한 `Getter()` 를 생성.

<br />

#### \<File\>_grpc.pb.go

서버가 구현해야할 **인터페이스 유형 메서드**와 Client가 호출할 수 있는 메서드의 **인터페이스 유형(또는 스텁)** 를 생성.

input에 stream이 있으면 `Send()` 가 구현되고, return에 stream이 있으면 `Revc()` 가 구현되어 있음.

Client와 Server에서 사용할 interface를 따로 생성함.

<br />

EX) proto에서 정의한 `RouteChat` 을 확인해보자 ~

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

위와 같이 생성된다.

<br /><br />

## 3. Create Server

<br />

### 서비스 구현

정의한 서비스들이 동작하게끔 내부 로직을 적어주자

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

service의 input이나 output에 stream을 받거나 주게 정의했다면, `stream pb.<Service>_<Method>Server` 를 인자로 받게 된다.

데이터를 전달할 때에는 `sream.send(Data)` 를 통해 전달.

데이터를 받을 때에는 `Data, err := stream.Recv()` 와 같이 받을 수 있다.

Data는 proto에 정의된 데이터 형태 ~

<br /><br />

### 서버 설정

서버를 설정해보자

``` go
// 서비스에 해당하는 Server struct 생성
type routeGuideServer struct {
	pb.UnimplementedRouteGuideServer
	savedFeatures []*pb.Feature

	mu         sync.Mutex
	routeNotes map[string][]*pb.RouteNote
}

// Server struct 인스턴스 생성
func newServer() *routeGuideServer {
	s := &routeGuideServer{routeNotes: make(map[string][]*pb.RouteNote)}
	s.loadFeatures(jsonDBFile)
	return s
}

func main() {
  // tcp로 net.Listen
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

  // grpc 서버 생성 후 Serve
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

<small>🔗 <a href="https://pkg.go.dev/testing">Go Testing</a></small>

음 사실 튜토리얼에서는 `client/client.go` 를 사용했지만, Test 기반 프로젝트를 만들고 싶어서 지금까지 배워온 걸 복습도 해볼겸 테스트 코드를 짰다.

``` go
// connection 생성
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
	
	// 위도: 409146138, 경도: -746188906인 위치의 장소 찾기
	t.Run("T1-GetFeature", func(t *testing.T) {
		printFeature(client, &pb.Point{Latitude: 409146138, Longitude: -746188906})
	})

	/*
		위도: [400000000, 420000000], 
		경도: [-750000000, -730000000] 위치의 장소 찾기
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

재밌다

<br /><br />