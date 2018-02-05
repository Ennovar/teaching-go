package server

import (
	"net/http"
	"strconv"
	"log"
	"fmt"
)

type Application struct {
	Port int
}

func (a *Application) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "The server is up and running!")
}

func main() {
	app := Application{
		Port: 3000,
	}

	server := http.Server{
		Addr: "localhost:" + strconv.Itoa(app.Port),
		Handler: &app,
	}

	log.Printf("Server is listening on localhost:%v\n", app.Port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("An error has occurred: %v", err.Error())
	}
}