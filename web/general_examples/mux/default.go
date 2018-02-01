package mux

import (
	"net/http"
	"fmt"
	"log"
)

func indexHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/" {
		fmt.Fprintln(res, "You're on the index!")
		return
	}

	http.Error(res, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func testHandler(res http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(res, "You're on the test page!")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/test", testHandler)

	log.Println("Server is listening on localhost:3000")
	if err := http.ListenAndServe("localhost:3000", nil); err != nil {
		log.Fatalf("An error has occurred: %v", err.Error())
	}
}
