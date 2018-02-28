package main

import (
	"net/http"
	"fmt"
	"log"

	"github.com/Ennovar/teaching-go/web/databases/postgres/handlers/user"
)

const serverPort = 3000

func main() {
	// Application Handler
	http.Handle("/", http.FileServer(http.Dir("approot")))

	http.HandleFunc("/api/user/get", nil) // ?id={insert_user_id_here}
	http.HandleFunc("/api/user/register", nil)
	http.HandleFunc("/api/user/update", nil)
	http.HandleFunc("/api/user/login", nil)
	http.HandleFunc("/api/user/verify", user.Verify)

	log.Printf("Server is listening on localhost:%v. CTRL+C to stop.\n", serverPort)
	if err := http.ListenAndServe(fmt.Sprintf("localhost:%v", serverPort), nil); err != nil {
		log.Fatalf("A fatal error has occurred: %v\n", err.Error())
	}
}
