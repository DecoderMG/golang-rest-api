package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/decodermg/golang-rest-api/internal/comment"
	"github.com/gorilla/mux"

	log "github.com/sirupsen/logrus"
)

// Handler - stores pointer to our comments service
type Handler struct {
	Router  *mux.Router
	Service *comment.Service
}

// Response - object to store responses from our API
type Response struct {
	Message string
	Error   string
}

// NewHandler - returns a pointer to a new Handler
func NewHandler(service *comment.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

// BasicAuth - Adds basic authentication to endpoints
func BasicAuth(original func(writer http.ResponseWriter, request *http.Request)) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		user, pass, ok := request.BasicAuth()
		if user == "admin" && pass == "password" && ok {
			original(writer, request)
		} else {
			writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
			sendErrorResponse(writer, "Not Authorized", errors.New("Not Authorized"))
		}
	}
}

// SetupRouters - sets up all the routes for our application
func (handler *Handler) SetupRoutes() {
	fmt.Println("Setting Up Routes")
	handler.Router = mux.NewRouter()

	handler.Router.Use(LoggingMiddleware)

	handler.Router.HandleFunc("/api/comment/{id}", handler.GetAllComments).Methods("GET")
	handler.Router.HandleFunc("/api/comment/{id}", BasicAuth(handler.PostComment)).Methods("POST")
	handler.Router.HandleFunc("/api/comment/{id}", handler.GetComment).Methods("GET")
	handler.Router.HandleFunc("/api/comment/{id}", BasicAuth(handler.UpdateComment)).Methods("PUT")
	handler.Router.HandleFunc("/api/comment/{id}", BasicAuth(handler.DeleteComment)).Methods("DELETE")

	handler.Router.HandleFunc("/api/health", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
		writer.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(writer).Encode(Response{Message: "I am Alive"}); err != nil {
			panic(err)
		}
	})
}

// sendErrorResponse - generalized error handler for all endpoints
func sendErrorResponse(writer http.ResponseWriter, message string, err error) {
	writer.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(writer).Encode(Response{Message: message, Error: err.Error()}); err != nil {
		panic(err)
	}
}

// LoggingMiddleware - Logs all requests to the rest APIs
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.WithFields(
			log.Fields{
				"Method": request.Method,
				"Path":   request.URL.Path,
			}).Info("Handled request")
		next.ServeHTTP(writer, request)
	})
}
