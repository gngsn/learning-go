package handler

import "net/http"

type DecoratorFunc func (http.ResponseWriter, *http.Request, http.Handler)

type DecoratorHandler struct {
	fn DecoratorFunc
	h http.Handler
}

/*
	HTTP handler
	decorator 함수를 먼저 실행 (이 예제에서는 logger)
*/
func (self *DecoratorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	self.fn(w, r, self.h)
}

/*
	decorator handler mapping
	member variable로 handler를 가짐
*/
func NewDecoratorHandler(h http.Handler, fn DecoratorFunc) http.Handler {
	return &DecoratorHandler{
		fn: fn,
		h: h,
	}
}