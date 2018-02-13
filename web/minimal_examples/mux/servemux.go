package main

import (
	"net/http"
	"fmt"
	"log"
)

func index(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/" {
		fmt.Fprintln(res, "You're on the index!")
		return
	}

	http.Error(res, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func test(res http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(res, "You're on the test page!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/test", test)

	log.Println("Server is listening on localhost:3000")
	if err := http.ListenAndServe("localhost:3000", mux); err != nil {
		log.Fatalf("An error has occurred: %v", err.Error())
	}
}