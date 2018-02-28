package main

import (
	"net/http"
	"fmt"
	"log"

	"github.com/Ennovar/teaching-go/web/databases/postgres/handlers"
)

const serverPort = 3000

func main() {
	http.HandleFunc("/api/user/get", nil) // ?id={insert_user_id_here}
	http.HandleFunc("/api/user/register", nil)
	http.HandleFunc("/api/user/update", nil)
	http.HandleFunc("/api/user/login", nil)
	http.HandleFunc("/api/user/verify", handlers.VerifyUser)

	if err := http.ListenAndServe(fmt.Sprintf("localhost:%v", serverPort), nil); err != nil {
		log.Fatalf("A fatal error has occurred: %v\n", err.Error())
	}
}
