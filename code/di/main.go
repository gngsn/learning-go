package di

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

/*
	아래의 함수는 어떻게 테스트할 수 있을까?
	Printf는 stdout으로 인쇄하기 때문에 테스트 프레임워크를 사용하기 어려움.

	이때 간단히 의존성 주입을 해서 테스트를 해볼 수 있음! -> main_test 확인
*/
func Greet(writer io.Writer, name string) {
    fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
    Greet(w, "world")
}
func main() {
    log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler)))
}