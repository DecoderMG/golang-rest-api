package main

import "fmt"

// App - the struct which contains things like points to database connections
type App struct{}

// Run - sets up our application
func (app *App) Run() error {
	fmt.Println("Setting Up Our App")
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
