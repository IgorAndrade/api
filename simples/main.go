package main

import (
	"log"
	"net/http"
	"time"
)

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "hello world"}`))
}

func main() {
	s := &server{}
	log.Fatal(http.ListenAndServe(":8080", s))
}

func main2() {
	s := &server{}
	ss := http.Server{
		Handler:     s,
		Addr:        ":8081",
		IdleTimeout: 2 * time.Second,
		ReadTimeout: 3 * time.Second,
	}
	ss.ListenAndServe()
}
