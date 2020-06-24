package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))

	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", fs))
}
