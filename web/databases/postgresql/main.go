package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Ennovar/teaching-go/web/databases/postgresql/pkg/handlers/index"
)

const (
	publicPort = 3000
)

func main() {
	assets := http.FileServer(http.Dir("public"))
	http.Handle("/assets/", http.StripPrefix("/assets/", assets))

	// Index Handler
	http.HandleFunc("/", index.Index)

	// User Handlers
	http.HandleFunc("/user", nil)
	http.HandleFunc("/user/login", nil)
	http.HandleFunc("/user/register", nil)
	http.HandleFunc("/user/dashboard", nil)
	http.HandleFunc("/user/logout", nil)

	log.Printf("Server is listening at localhost:%v. CTRL+C to stop.\n", publicPort)
	if err := http.ListenAndServe(fmt.Sprintf("localhost:%v", publicPort), nil); err != nil {
		log.Fatalf("A fatal error has occurred: %v\n", err.Error())
	}
}
