package selects

import (
	"fmt"
	"net/http"
	"time"
)


var tenSecondTimeout = 10 * time.Second

/*
	Synchronising processes

	time.After는 select 를 사용할 때 매우 편리한 기능
	수신 중인 채널이 값을 반환하지 않을 경우 영원히 차단되는 코드를 작성할 수 있음
	time.After는 chan을 반환하고 사용자가 정의한 시간 후에 신호를 보냄
*/
func Racer(a, b string) (winner string, error error) {
    return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
    select {
    case <-ping(a):
        return a, nil
    case <-ping(b):
        return b, nil
    case <-time.After(timeout):
        return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
    }
}

func measureResponseTime(url string) time.Duration {
    start := time.Now()
    http.Get(url)
    return time.Since(start)
}

func ping(url string) chan struct{} {
    ch := make(chan struct{})
    go func() {
        http.Get(url)
        close(ch)
    }()
    return ch
}