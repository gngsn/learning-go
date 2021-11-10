package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gngsn/learning-go/practice/logging/app"
	"github.com/gngsn/learning-go/practice/logging/handler"
)

// decorator function
func logger(w http.ResponseWriter, r *http.Request, h http.Handler) {
	start := time.Now()
	log.Println("[LOGGER1] Started")
	h.ServeHTTP(w, r)
	log.Println("[LOGGER1] Completed time:", time.Since(start).Milliseconds())
}

func logger2(w http.ResponseWriter, r *http.Request, h http.Handler) {
	start := time.Now()
	log.Println("[LOGGER2] Started")
	h.ServeHTTP(w, r)
	log.Println("[LOGGER2] Completed time:", time.Since(start).Milliseconds())
}

func NewHandler() http.Handler {
	h := app.NewHandler()
	
	/*
		decorator로 감싸기
		
		DecoratorHandler(DecoratorHandler(mux, logger), logger2)
		: 여기서 mux는 app.NewHandler()
		: DecoratorHandler에서 fn(두번째 인자)를 먼저 실행하게 해서 순서는 43~47lines에 해당
		-> 이렇게 decorator는 암호화 encrypt, 인증, 마케팅 정보 보내는 등으로 확장시킬 수 있음.
		-> 새로운 코드를 한 곳에서 결합시켜주기만 하면 됨.
	*/
	h = handler.NewDecoratorHandler(h, logger)
	h = handler.NewDecoratorHandler(h, logger2)
	return h
}

/*
	2021/11/10 18:50:01 [LOGGER2] Started
	2021/11/10 18:50:01 [LOGGER1] Started
	2021/11/10 18:50:01 [LOGGER1] Completed time: 0
	2021/11/10 18:50:01 [LOGGER2] Completed time: 0
*/

func main() {
	mux := NewHandler()

	http.ListenAndServe(":3000", mux)
}