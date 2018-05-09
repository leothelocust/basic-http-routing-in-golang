package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloHTTP)
	http.HandleFunc("/user/", userProfile)
	http.HandleFunc("/", notFound)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello HTTP")
}

func notFound(w http.ResponseWriter, req *http.Request) {
	http.NotFound(w, req)
}

func userProfile(w http.ResponseWriter, req *http.Request) {
	userID := req.URL.Path[len("/user/"):]
	fmt.Fprintf(w, "User Profile: %q", userID)
}
