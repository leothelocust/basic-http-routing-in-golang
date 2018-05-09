package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", HelloHTTP)
	http.HandleFunc("/", NotFound)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HelloHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello HTTP")
}

func NotFound(w http.ResponseWriter, req *http.Request) {
	http.NotFound(w, req)
}
