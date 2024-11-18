package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><body><h1>Hello, %s.</h1></body></html>", r.URL.Path[1:])
}

func main() {
	fmt.Println("Serving http://localhost:8080/world")
	err := http.ListenAndServe(":8080", http.HandlerFunc(handler))
	if err != nil {
		log.Fatal(err)
	}
}
