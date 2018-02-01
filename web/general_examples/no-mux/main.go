package no_mux

import (
	"net/http"
	"fmt"
	"strconv"
	"log"
)

type Application struct {
	Port int
}

func (a *Application) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	fmt.Fprintf(res, "From here, without a multiplexer, you can analyze the path and do whatever you want to do based off of that. (path:%v)", path)
}

func main() {
	app := Application{
		Port: 3000,
	}

	log.Printf("Server is listening on localhost:%v\n", app.Port)
	if err := http.ListenAndServe("localhost:"+strconv.Itoa(app.Port), &app); err != nil {
		log.Fatalf("An error has occurred: %v\n", err.Error())
	}
}