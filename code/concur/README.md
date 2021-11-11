## Concurrency



'한 가지 이상의 진행 중인 것'을 의미하는 동시성에 대한 코드

<br/>

여기서 말하는 동시성을 아래의 예시를 통해 확인해보자

> 예를 들어, 오늘 아침에 차 한 잔을 만들었다. 주전자를 올려놓고 끓기를 기다리는 동안, 냉장고에서 우유를 꺼내고 찬장에서 차와 내가 좋아하는 머그잔을 찾아 꺼낸다. 그리고는 주전자가 끓으면 물을 컵에 넣는다.
>
> 주전자를 올려놓고 멍하니 주전자가 끓을 때까지 주전자를 바라보고 있다가 주전자가 끓으면 다른 모든 것을 하지는 않음.

<br/>

기본적으로, Go에서는 함수를 호출하면 그 함수가 반환할 때까지 기다리며, 반환값이 없더라도 그 함수가 끝날 때까지 기다린다.

이렇게 작업이 끝날 때까지 기다리게 하는 것을 Blocking 되었다고 한다

⭐️ Go에서 위와같이 Blocking 되지 않게끔 독자적인 프로세스를 실행하는 것을 고루틴 <small>goroutine</small> 이라고 한다.

<br/>

Go의 코드를 페이지에 따라 읽어내려갈 때 함수가 호출되면 각각의 함수 '내부'를 읽고 내려간다. 이 때, 독자적인 프로세스는 페이지를 내려 읽어가는 기존의 리더에서 빠져나온 다른 리더가 함수 내부를 읽어가고, 기존의 리더가 계속해서 페이지를 읽어가는 것과 같다.

<br/>

### goroutine - go *function*()

> go *doSomething*()

<br/>

<br/>

고루틴을 시작하는 유일한 방법은 함수를 호출할 때 앞에 **go** 키워드를 적는데, 익명 함수 <small>Anonymous functions</small> 를 사용하는 경우가 많다. 익명 함수 리터럴은 일반 함수 선언과 동일하게 보이지만 이름은 없다

<br/>

<br/>

### Code - main.go

이 코드에서 구현하는 `CheckWebsite` 는 여러개의 URL을 체크(하는 척..ㅎ)하는 함수임

다른 결과가 나오기 전에 여러를 동시에 





### fatal error: concurrent map writes

위의 에러는 Maps를 두 개 이상의 고루틴이 동시에 접근해서 나오는 **Race Condition**

각 고루틴이 Maps에 접근하는 시기를 정확히 제어할 수 없기 때문에 Maps에 고루틴을 쓰는게 취약함.

테스트할 때 race flag를 사용해보자  `test -race`



``` bash
$ go test -race
==================
WARNING: DATA RACE
Write at 0x00c000012360 by goroutine 10:
...

Previous write at 0x00c000012360 by goroutine 8:
...

Goroutine 10 (running) created at:
```





### Channel : Resolve Race Condition

Channel을 사용해서 고루틴을 조정해서 Race Condition을 해결할 수 있음.

Channel는 value를 주고 받을 수 있는 Go의 데이터 구조로, 서로 다른 프로세스 사이에 자세한 내용과 함께 통신을 주고 받음

``` go
type result struct {
    string
    bool
}
resultChannel := make(chan result)

// channel 추가
resultChannel <- result{u, wc(u)}

/*
	channel 데이터를 변수에 담기
	r.string으로 접근 가능
*/
r := <-resultChannel
```





