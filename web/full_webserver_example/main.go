package main

import (
	"fmt"

	"github.com/Ennovar/teaching-go/web/full_webserver_example/pkg/application"
)

func main() {
	app := application.New("application/", "document_root/", 3000)

	fmt.Println("CTRL+C to stop application")

	err := app.Start()
	if err != nil {
		fmt.Errorf("Error starting application: %v\n", err)
	}
}
