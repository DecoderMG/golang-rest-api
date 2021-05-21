package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler - stores pointer to our comments service
type Handler struct {
	Router *mux.Router
}

// NewHandler - returns a pointer to a new Handler
func NewHandler() *Handler {
	return &Handler{}
}

// SetupRouters - sets up all the routes for our application
func (handler *Handler) SetupRoutes() {
	fmt.Println("Setting Up Routes")
	handler.Router = mux.NewRouter()
	handler.Router.HandleFunc("/api/health", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "I am alive!")
	})
}
