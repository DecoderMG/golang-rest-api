package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/decodermg/golang-rest-api/internal/comment"
	"github.com/gorilla/mux"
)

// Handler - stores pointer to our comments service
type Handler struct {
	Router  *mux.Router
	Service *comment.Service
}

// NewHandler - returns a pointer to a new Handler
func NewHandler(service *comment.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

// SetupRouters - sets up all the routes for our application
func (handler *Handler) SetupRoutes() {
	fmt.Println("Setting Up Routes")
	handler.Router = mux.NewRouter()

	handler.Router.HandleFunc("/api/comment/{id}", handler.GetAllComments).Methods("GET")
	handler.Router.HandleFunc("/api/comment/{id}", handler.PostComment).Methods("POST")
	handler.Router.HandleFunc("/api/comment/{id}", handler.GetComment).Methods("GET")
	handler.Router.HandleFunc("/api/comment/{id}", handler.UpdateComment).Methods("PUT")
	handler.Router.HandleFunc("/api/comment/{id}", handler.DeleteComment).Methods("DELETE")

	handler.Router.HandleFunc("/api/health", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "I am alive!")
	})
}

// GetComment - retrieve comment by ID
func (handler *Handler) GetComment(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Fprintf(writer, "Unable to parse UINT from ID")
	}

	comment, err := handler.Service.GetComment(uint(i))
	if err != nil {
		fmt.Fprintf(writer, "Error Retrieving Comment by ID")
	}

	fmt.Fprintf(writer, "%+v", comment)
}

// GetAllComments - retrives all comments from the comment server
func (handler *Handler) GetAllComments(writer http.ResponseWriter, request *http.Request) {
	comments, err := handler.Service.GetAllComments()
	if err != nil {
		fmt.Fprintf(writer, "Failed to retrieve all comments")
	}
	fmt.Fprintf(writer, "%+v", comments)
}

// PostComment - adds new comment
func (handler *Handler) PostComment(writer http.ResponseWriter, request *http.Request) {
	comment, err := handler.Service.PostComment(comment.Comment{
		Slug: "/",
	})

	if err != nil {
		fmt.Fprintf(writer, "Failed to post new comment")
	}
	fmt.Fprintf(writer, "%+v", comment)
}

// UpdateComment - updates a comment by ID
func (handler *Handler) UpdateComment(writer http.ResponseWriter, request *http.Request) {
	comment, err := handler.Service.UpdateComment(1, comment.Comment{
		Slug: "/",
	})

	if err != nil {
		fmt.Fprintf(writer, "Failed to update comment")
	}
	fmt.Fprintf(writer, "%+v", comment)
}

// DeleteComment - deletes a comment by ID
func (handler *Handler) DeleteComment(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	commentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Fprintf(writer, "Failed to parse uint from ID")
	}

	err = handler.Service.DeleteComment(uint(commentID))
	if err != nil {
		fmt.Fprintf(writer, "Failed to delete comment by ID")
	}
	fmt.Fprintf(writer, "Successfully deleted comment")
}
