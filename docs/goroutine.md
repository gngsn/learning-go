## Goroutine

이 문서는 [고루틴 강의](https://youtu.be/tRdODUXV3ik) 를 보고 정리한 문서입니다!



### Thread

CPU는 명령어 실행 주체

명령어 : `연산자 - 피연산자`만을 확인함

시작 시점 `main()` 으로 해서 **Instruction pointer**를 main()위에 위치시키고 한 칸씩 내려가면서 실행시킴

CPU는 IP 포인터가 가르키고 있는 것만 처리하면 됨!



#### Multi-Thread

각 thread는 각자마다의 IP Pointer가 있음. 각각의 Thread의 IP Pointer를 실행시키다가 다음 Thread의 IP Pointer를 실행하는데, 코어가 이 과정을 반복해서 Thread를 바꿔줌

-> OS는 단순히 계산. OS가 다 해주는 것.



Thread 전환에는 성능상의 비용이 발생 -> Context Switching

Core 1개인데 Thread가 3개면 Context Switching이 많이 발생

Core 3개인데 Thread가 3개면 Context Switching이 발생하지 않음



#### Multi-Thread VS Multi-Process

프로그램을 실행하면 메모리에 로드된다. 메모리에 로드된 그 프로그램 인스턴스가 실행 주체. 

하나의 프로그램이 더 들어온다면 실행 인스턴스가 두 개가 되는 거고 서로 다른 데이터, state를 가진다. -> 이것이 멀티 프로세스.





### Go Routine

Go에서 만든 경량 스레드 (가벼운 스레드).

메인함수도 고루틴 - 메인 고루틴 (`main()`)이라고 부른다. 일반 스레드와 마찬가지로, 반드시 하나의 고루틴을 가진다.

새로운 고루틴은 아래와 같이 만듦



``` go
// /code/goroutine/goroutine.go

func PrintHangul() {
  hanguls := []rune{'가', '나', '다', '라', '마', '바', '사'}
  for _, v := range hanguls {
    time.Sleep(300 * time.Millisecond)  // 1sec = 1000ms, 0.3s
    fmt.Printf("%c ", v)
  }
}

func PrintNumbers() {
  for i := 1; i <= 5; i++ {
    time.Sleep(400 * time.Millisecond)  // 0.4s
    fmt.Printf("%d ", i)
  }
}

func main() {
  go  PrintHangul()
  go  PrintNumbers()
  
  time.Sleep(3 * time.Second)
}


// $ go run ex1.go
// 가 1 나 2 다 3 라 마 4 바 5 사
```



`main()`, `PrintHangul()`, `PrintNumbers()` 이라는 세 개의 고루틴이 진행중임



<img src="/Users/gyeongseon/Library/Application Support/typora-user-images/스크린샷 2021-11-17 오후 6.49.46.png" alt="스크린샷 2021-11-17 오후 6.49.46" style="zoom:33%;" />



`main()` 의 `time.Sleep(3 * time.Second)` 는 언제 끝날지 모르니까 3초 기다림 ~

안기다리면 `main()` 이 종료되기 때문에 프로그램이 종료됨



매번 기다려줘야 하나? NO



#### 서브 고루틴이 종료될 때까지 대기

``` go
var wg sync.WaitGroup

wg.Add(3)
wg.Done()
wg.Wait()
```



``` go
// /code/goroutine/wg.go
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func SumAtoB(a, b int) {
	sum := 0
  	for i := a; i <= b; i++ {
  		 sum += i
  	}
  	fmt.Printf("%d부터 %d까지 합계는 %d입니다\n", a, b, sum)
  	wg.Done()			// 10번의 Done을 실행하면 Wait()가 실행이 됨
}

func main() {
	wg.Add(10)
	for i := 0; i <= 10; i++ {
		go SumAtoB(1, 1000)
	}
	wg.Wait()
}
```



#### 고루틴 동작원리

고루틴은 OS 쓰레드를 이용하는 경량 쓰레드(light weight thread)임.

고루틴 != 쓰레드. 고루틴이 쓰레드를 이용함.



<img src="/Users/gyeongseon/Library/Application Support/typora-user-images/스크린샷 2021-11-17 오후 9.30.22.png" alt="스크린샷 2021-11-17 오후 9.30.22" style="zoom:50%;" />



코어 2개가 두 개인 머신이 있을 때 아래의 상황을 확인해보자

**고루틴 1개일 때,**

Go 프로그램에서 OS 스레드를 만들어서 코어와 매칭하고, 이 코어와 벗어나지 않게 묶어둠.

그리고 고루틴을 만들어서 OS스레드를 통해 코어와 연결되게(실행되게) 만듦.



**고루틴 2개일 때,**

OS 스레드를 하나 더 만들고 고루틴 하나를 더 만들어서 새로운 OS 스레드 2와 코어 2와 연결하게 만든다. 



**고루틴 3개일 때,**

새로운 스레드를 만들지 않음. 실행되던 고루틴 중 완료가 되는 것이 있으면 그 자리로 들어가고, 고루틴 하나에서 시스템콜이 발생 하면 대기하는 고루틴과 교체됨.

그럼 대기하는 고루틴이 많지 않을까? 사실 시스템에서는 시스템콜이 발생하는 경우(파일 읽기/쓰기, 네트워크 읽기/쓰기 등등 대기되는 상태가 많음)가 많고 수시로 교체되는데 우리는 잘 모름. 그 만큼 대기하는 고루틴들이 교체되는 과정도 이와같이 느껴지지 않을 만큼 빠르게 진행이 된다 ~



**왜 좋을까?**

OS에서 Context Switching이 일어나지 않음. 물론, Goroutine이 교체되는 것도 Context Switching임. IP Pointer와 Stack 메모리가 있음. Goroutine의 Context Switching는 Stack사이즈(늘어날 순 있지만 기본적으로 작음)가 굉장히 경량이기 때문에 아무리 많아도 그 비용이 적다.

```
고루틴은 생성하는데에 많은 메모리를 필요로 하지 않습니다. 오직 **2kB**의 스택 공간만 필요로 합니다. 고루틴을 할당하고 필요에 따라 힙 저장 공간을 확보하여 사용합니다. 반대로 쓰레드는 쓰레드의 메모리와 다른 쓰레드의 메모리 간의 경비 역할을 하는 Guard page라고 불리는 메모리 영역과 함께 **1Mb(500배 더 큼)**로 시작합니다.

따라서 수신 요청을 처리하는 서버는 문제 없이 요청 한 건 당 하나의 고루틴을 만들 수 있지만, 요청 한 건 당 하나의 쓰레드는 결과적으로 OutOfMemoryError가 일어나게 될 것 입니다. 이런 일은 자바에 한정되어 있는 것이 아닙니다. 동시성의 주요 수단으로 OS 쓰레드를 사용하는 언어라면 언젠가는 이 문제에 대면할 것입니다.

// https://stonzeteam.github.io/How-Goroutines-Work/
```







#### 동시성 프로그래밍의 주의점

동시성의 문제

``` go
package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	Balance int
}

func DepositAndWithdraw(account *Account) {
	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance should not be negative value: %d", account.Balance))
	}
	account.Balance += 1000
	time.Sleep(time.Millisecond)
	account.Balance -= 1000
}


func main() {
	var wg sync.WaitGroup

	account := &Account{0}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for {
				DepositAndWithdraw(account)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
```



위의 코드를 실행해보면 `panic: Balance should not be negative value: -2000` 와 같은 에러로 `panic` 이 발생함

`account` 의 `Balance` 에 동시에 접근하면서 발생하는 동시성 문제. 



`account.Balance += 1000` 는 아래와 같은 두 가지 연산임

```assembly
ADD Balance 1000
MOV Balance Rx1
```

많은 고루틴이 동시에 접근하는 레지스터에 접근할 때 변경되기 전 ADD를 위해 데이터를 가져다 MOV로 대입하는 문제 ~



### Mutex (Mutial Exclusion, 상호배제)

Lock을 걸어서 다른 실행 단위가 접근하지 못하는 거임. 

칠판에 그림을 그리려고 할 때, 하나의 분필을 먼저 잡은 사람이 그림을 그릴 수 있는 거임.



``` go
package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex

type Account struct {
	Balance int
}

func DepositAndWithdraw(account *Account) {
	mutex.Lock()         // Mutex Key 획득
	defer mutex.Unlock() // defer를 사용한 Unlock()
	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance should not be negative value: %d", account.Balance))
	}
	account.Balance += 1000
	time.Sleep(time.Millisecond)
	account.Balance -= 1000
}

func main() {
	var wg sync.WaitGroup

	account := &Account{0}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for {
				DepositAndWithdraw(account)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
```

`Lock()` 을 얻은 고루틴 하나만 실행하고 `Lock()` 을 반환



#### Mutex 문제점

✔️ 동시성 프로그래밍으로 인한 성능 향상을 얻을 수 없음. 

✔️ 과도한 락킹으로 인해 성능이 하락될 수도 있음 -> 락을 획득하고 반납하는 시간이 성능 문제 발생

✔️ DEADLOCK. 고루틴을 완전히 멈추게 함. 

<small>Deadlock: 두 개 이상의 작업이 서로 상대방의 작업이 끝나기 만을 기다리고 있기 때문에 결과적으로 아무것도 완료되지 못하는 상태</small>



