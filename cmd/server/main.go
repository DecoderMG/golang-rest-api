package main

import (
	"fmt"
	"net/http"

	"github.com/decodermg/golang-rest-api/internal/comment"
	"github.com/decodermg/golang-rest-api/internal/database"
	transportHTTP "github.com/decodermg/golang-rest-api/internal/transport/http"

	log "github.com/sirupsen/logrus"
)

// App - the struct which contains things like points to database connections
type App struct {
	Name    string
	Version string
}

// Run - sets up our application
func (app *App) Run() error {
	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(
		log.Fields{
			"AppName":    app.Name,
			"AppVersion": app.Version,
		}).Info("Setting Up Our App")

	var err error

	db, err := database.NewDatabase()
	if err != nil {
		return err
	}

	err = database.MigrateDB(db)
	if err != nil {
		log.Error("Failed to setup database")
		return err
	}

	commentService := comment.NewService(db)

	handler := transportHTTP.NewHandler(commentService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		log.Error("Failed to set up server")
		return err
	}

	log.Info("App startup Successful")
	return nil
}

// Our main entrypoint for the application
func main() {
	fmt.Println("Go REST API Server")
	app := App{
		Name:    "Comment API",
		Version: "1.0",
	}
	if err := app.Run(); err != nil {
		log.Error(err)
		log.Fatal("Error Starting Up Our REST API")
	}
}
