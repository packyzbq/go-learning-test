package main

import (
	"github.com/go-learning-test/example"
	"log"
	"net/http"
)

func main() {
	server := &example.PlayerServer{Store: example.NewInMemoryPlayerStore()}
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
