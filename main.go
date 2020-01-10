package main

import (
	"github.com/go-learning-test/ioc"
	"net/http"
)

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	ioc.Greet(w, "world")
}

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}
