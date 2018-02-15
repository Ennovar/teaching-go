package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Ennovar/teaching-go/web/templates_complex/pkg/handlers/about"
	"github.com/Ennovar/teaching-go/web/templates_complex/pkg/handlers/api/test"
	"github.com/Ennovar/teaching-go/web/templates_complex/pkg/handlers/index"
	"github.com/Ennovar/teaching-go/web/templates_complex/pkg/handlers/pages"
)

const (
	publicPort = 3000
)

func main() {
	// Implement asset handler, serving out of public/
	assets := http.FileServer(http.Dir("public"))
	http.Handle("/assets/", http.StripPrefix("/assets/", assets))

	// Index Handler
	http.HandleFunc("/", index.Index)

	// Handlers for /about
	http.HandleFunc("/about", about.Index)
	http.HandleFunc("/about/contact", about.Contact)

	// Handlers for /pages/* (static HTML files, no need for templates)
	http.HandleFunc("/pages/foo", pages.Foo)
	http.HandleFunc("/pages/bar", pages.Bar)

	// Handlers for Test APIs
	http.HandleFunc("/api/test/get", test.Get)
	http.HandleFunc("/api/test/post", test.Post)

	log.Printf("The server is now listening at localhost:%v. CTRL+C to Stop Listening.\n", publicPort)

	// Start the server
	err := http.ListenAndServe(fmt.Sprintf("localhost:%v", publicPort), nil)
	if err != nil {
		log.Fatalf("A fatal error has occurred: %v\n", err.Error())
	}
}
