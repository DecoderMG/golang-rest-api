package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/decodermg/golang-rest-api/internal/transport/http"
)

// App - the struct which contains things like points to database connections
type App struct{}

// Run - sets up our application
func (app *App) Run() error {
	fmt.Println("Setting Up Our App")

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}
	return nil
}

// Our main entrypoint for the application
func main() {
	fmt.Println("Go REST API Server")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error Starting Up")
		fmt.Println(err)
	}
}
