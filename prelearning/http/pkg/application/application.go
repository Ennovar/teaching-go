package application

import (
	"net/http"
	"strconv"
	"time"
)

type Application struct {
	Directory    string
	DocumentRoot string
	Port         int
	Status       int
}

var server http.Server

func New(dir string, root string, port int) *Application {
	app := &Application{
		Directory:    dir,
		DocumentRoot: root,
		Port:         port,
		Status:       0,
	}

	server = http.Server{
		Addr:         "localhost:" + strconv.Itoa(app.Port),
		Handler:      app,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	return app
}
