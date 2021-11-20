## Channel

이 문서는 [Channel 강의](https://youtu.be/F6T9x-M7GNE) 를 보고 정리한 문서입니다!

<br/>

채널은 고루틴끼리 메세지를 전달할 수 있는 메세지 큐

<br/>

### 생성

make() 로 채널 인스턴스 생성

``` go
var messages chan string = make(chan string)
// var <instance_name> <channel_type> = make(chan <message_type>)
```

<br/>

### 데이터 넣기

``` go
messages <- "This is Message"
// <chan_instance> <- data
```

<br/>

### 데이터 빼기

``` go
var msg string  <-  messages
// variable <- <chan_instance>
```





### Channel 크기

기본 크기 0 : 보관함이 없음

빈자리가 있으면 데이터를 두고 간다 ~



#### 크기 설정하기

``` go
func main() {
  ch := make(chan int) 				// 크기가 0인 채널 생성
  ch <- 21             				// fatal error: all goroutines are asleep - deadlock!

  fmt.Println("Never print") 	// 실행되지 안됨
}
```

``` go
func main() {
  ch := make(chan int, 2) 				// 크기가 2인 채널 생성
  ch <- 21

  fmt.Println("Never print")    	// "Never print" 출력!
}
```





### 채널에서 데이터 대기













