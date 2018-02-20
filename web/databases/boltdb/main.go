package main

import (
	"net/http"
	"fmt"
	"log"
)

const (
	publicPort = 3000
)

func main() {
	// Index Handler
	http.HandleFunc("/", nil)

	// Posts Handlers
	http.HandleFunc("/posts", nil)
	http.HandleFunc("/posts/view", nil)
	http.HandleFunc("/posts/new", nil)
	http.HandleFunc("/posts/delete", nil)

	err := http.ListenAndServe(fmt.Sprintf("localhost:%v", publicPort), nil)
	if err != nil {
		log.Fatalf("A fatal error has occurred: %v\n", err.Error())
	}
}
