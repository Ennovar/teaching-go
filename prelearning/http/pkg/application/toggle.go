package application

import (
	"context"
	"fmt"
)

func (app *Application) Start() error {
	app.Status = 1
	fmt.Printf("Application listening on localhost:%d\n", app.Port)
	return server.ListenAndServe()
}

func (app *Application) Stop(graceful bool) error {
	if graceful {
		context, cancel := context.WithTimeout(context.Background(), 5)
		defer cancel()

		err := server.Shutdown(context)
		if err == nil {
			app.Status = 0
			fmt.Println("Application has stopped listening")
			return nil
		}

		fmt.Printf("Graceful shutdown failed attempting forced: %v\n", err)
	}

	if err := server.Close(); err != nil {
		return err
	}

	app.Status = 0
	fmt.Println("Application has stopped listening")
	return nil
}
