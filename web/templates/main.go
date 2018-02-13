package main

import (
	"net/http"
	"log"

	"github.com/Ennovar/teaching-go/web/templates/pkg/handlers/index"
	"github.com/Ennovar/teaching-go/web/templates/pkg/handlers/about"
)

func main() {
	// http.FileServer returns a Handler interface
	assets := http.FileServer(http.Dir("public"))

	// To access, say public/css/main.css, you would go to the URL
	// /assets/css/main.css. This Handle function will handle all
	// of the static assets found in the /public folder.
	http.Handle("/assets/", http.StripPrefix("/assets/", assets))

	http.HandleFunc("/", index.Index)
	http.HandleFunc("/about", about.Index)
	http.HandleFunc("/about/contact", about.Contact)

	log.Println("The server is listening on localhost:3000. CTRL+C to stop listening...")

	err := http.ListenAndServe("localhost:3000", nil)
	if err != nil {
		log.Fatalf("A fatal error has occurred: %v\n", err.Error())
	}
}