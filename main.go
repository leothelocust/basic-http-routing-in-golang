package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
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
	userID, err := getID(w, req)
	if err != nil {
		return
	}
	fmt.Fprintf(w, "User Profile: %q", userID)
}

var validPath = regexp.MustCompile("^/(user)/([0-9]+)$")

func getID(w http.ResponseWriter, req *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(req.URL.Path)
	if m == nil {
		http.NotFound(w, req)
		return "", errors.New("Invalid ID")
	}
	return m[2], nil // The ID is the second subexpression.
}
