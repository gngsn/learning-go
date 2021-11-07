package main

import (
	"net/http"

	"github.com/gngsn/learning-go/myapp"
)


func main() {
	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}