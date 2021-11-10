package main

import (
	"net/http"

	"github.com/gngsn/learning-go/practice/restapi/app"
)

func main() {
	http.ListenAndServe(":3000", app.NewHandler())
}