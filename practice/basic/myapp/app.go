package myapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type testHandler struct {}
type jsonHandler struct {}

type User struct {
	FirstName string	`json:"first_name"`
	LastName  string	`json:"last_name"`
	Email	  string	`json:"email"`
	CreatedAt time.Time	`json:"created_at"`
}

func (t *testHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Test")
}

func queryHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "world"
	}
	fmt.Fprintf(w, "Hello %s", name)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func (j *jsonHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	user.CreatedAt = time.Now()
	data, _ := json.Marshal(user)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
}

func NewHttpHandler() http.Handler {
	mux := http.NewServeMux()

	
	mux.HandleFunc("/", indexHandler)
	
	mux.Handle("/test", &testHandler{})
	
	mux.HandleFunc("/query", queryHandler)
	
	mux.Handle("/json", &jsonHandler{})
	
	mux.Handle("/file", http.FileServer((http.Dir("public"))))

	return mux
}
