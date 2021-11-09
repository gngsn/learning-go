## Why use gRPC?
Our example is a simple route mapping application that lets clients get information about features on their route, create a summary of their route, and exchange route information such as traffic updates with the server and other clients.

With gRPC we can define our service once in a `.proto` file and generate clients and servers in any of gRPC’s supported languages, which in turn can be run in environments ranging from servers inside a large data center to your own tablet — all the complexity of communication between different languages and environments is handled for you by gRPC. We also get all the advantages of working with protocol buffers, including efficient serialization, a simple IDL, and easy interface updating.

우리의 샘플 코드는 클라이언트가 경로의 기능에 대한 정보를 얻고, 경로 요약을 만들고, 서버 및 기타 클라이언트와 트래픽 업데이트와 같은 경로 정보를 교환할 수 있는 간단한 경로 매핑 응용 프로그램이 있습니다.

gRPC를 사용하면 서비스를 .proto' 파일에 한 번 정의하고 gRPC가 지원하는 언어로 클라이언트와 서버를 생성할 수 있으며, 이는 대규모 데이터 센터 내의 서버에서부터 사용자 자신의 태블릿에 이르기까지 다양한 환경에서 실행될 수 있습니다. 다양한 언어와 환경 간의 모든 복잡한 커뮤니케이션을 GRPC에서 처리할 수 있습니다. 또한 효율적인 직렬화, 간단한 IDL, 쉬운 인터페이스 업데이트를 포함하여 프로토콜 버퍼를 사용하는 모든 이점을 누릴 수 있습니다.



<!-- 
### Defining the service

To define a service, you specify a named service in your `.proto` file:

``` go
service RouteGuide {
   ...
}
```


 -->
