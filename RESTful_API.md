## Go Framework ?

https://golangkorea.github.io/post/go-start/getting-start/

Go 언어는 **동시성(concurrency)**를 잘 지원하는 것으로 유명하다. Go는 **고루틴(goroutine)**라는 경량스레드(lightweight thread)를 제공하는데, 고루틴간 메시지를 주고 받을 수 있는 채널(channel)을 이용하면, 아주 쉽게(정말 쉽다) 동시성 프로그램을 개발 할 수 있다. 고루틴은 얼랑(Erlang)의 경량 쓰레드와 매우 유사한데, **2k** 정도로 그 크기가 매우 작다. 많은 수의 고루틴을 시스템 부담을 최소화 하면서 만들 수 있다.

Go 언어를 사용하다보면, 웹 애플리케이션을 만들기가 매우 편하다는 느낌을 받게 된다. 특히 **MSA(Microservice Architecture)**와 **REST(Representational State Transfer)** 모델의 애플리케이션을 쉽게 만들 수 있다. 루비나 파이선 같은 언어의 경우 다양한 **웹 프레임워크**중에서 선택을 고민하게 마련인데, Go 언어는 기본으로 제공하는 **net/http** 패키지로 충분하다. 물론 Go 언어도 다양한 마이크로 프레임워크와 풀 프레임워크를 제공하긴 하지만 이런 프레임워크를 쓰면, “왜 프레임워크를 쓰세요 ? 그냥 기본(net/http) 패키지 쓰세요”라는 말을 들을 정도로 강력하다.

대규모의 분산 시스템을 유지해야 하는 구글의 요구를 위해서 웹 개발 관련 패키지가 강력해진 것 같다.






<!-- 
# Developing a RESTful API with Go and Gin

https://golang.org/doc/tutorial/web-service-gin


<!-- ### Create Project

Using the command prompt, create a directory for your code called web-service-gin.

```
$ mkdir web-service-gin
$ cd web-service-gin
```



Create a module in which you can manage dependencies.

Run the `go mod init` command, giving it the path of the module your code will be in.

```
$ go mod init example/web-service-gin
go: creating new go.mod: module example/web-service-gin
```

This command creates a **go.mod** file in which **dependencies you add will be listed for tracking**. For more about naming a module with a module path, see [Managing dependencies](https://golang.org/doc/modules/managing-dependencies#naming_module).





### Create the data -->
