package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Ennovar/teaching-go/web/databases/postgresql/pkg/handlers/index"
	userHandler "github.com/Ennovar/teaching-go/web/databases/postgresql/pkg/handlers/user"
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
	http.HandleFunc("/user", userHandler.Login)
	http.HandleFunc("/user/login", userHandler.Login)
	http.HandleFunc("/user/register", userHandler.Register)
	http.HandleFunc("/user/dashboard", userHandler.Dashboard)
	http.HandleFunc("/user/logout", userHandler.Logout)

	log.Printf("Server is listening at localhost:%v. CTRL+C to stop.\n", publicPort)
	if err := http.ListenAndServe(fmt.Sprintf("localhost:%v", publicPort), nil); err != nil {
		log.Fatalf("A fatal error has occurred: %v\n", err.Error())
	}
}
