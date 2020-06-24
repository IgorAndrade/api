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

//https://astaxie.gitbooks.io/build-web-application-with-golang/en/07.4.html
//https://www.alexedwards.net/blog/serving-static-sites-with-go
