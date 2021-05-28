package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/decodermg/golang-rest-api/internal/comment"
	"github.com/gorilla/mux"
)

// GetComment - retrieve comment by ID
func (handler *Handler) GetComment(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.WriteHeader(http.StatusOK)

	vars := mux.Vars(request)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(writer, "Unable to parse UINT from ID", err)
	}

	comment, err := handler.Service.GetComment(uint(i))
	if err != nil {
		sendErrorResponse(writer, "Error Retrieving Comment by ID", err)
	}

	if err := json.NewEncoder(writer).Encode(comment); err != nil {
		panic(err)
	}
}

// GetAllComments - retrives all comments from the comment server
func (handler *Handler) GetAllComments(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.WriteHeader(http.StatusOK)
	comments, err := handler.Service.GetAllComments()
	if err != nil {
		sendErrorResponse(writer, "Failed to retrieve all comments", err)
	}

	if err := json.NewEncoder(writer).Encode(comments); err != nil {
		panic(err)
	}
}

// PostComment - adds new comment
func (handler *Handler) PostComment(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.WriteHeader(http.StatusOK)

	var comment comment.Comment
	if err := json.NewDecoder(request.Body).Decode(&comment); err != nil {
		sendErrorResponse(writer, "Failed to decode JSON body", err)
	}
	comment, err := handler.Service.PostComment(comment)

	if err != nil {
		sendErrorResponse(writer, "Failed to post new comment", err)
	}
	if err := json.NewEncoder(writer).Encode(comment); err != nil {
		panic(err)
	}
}

// UpdateComment - updates a comment by ID
func (handler *Handler) UpdateComment(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.WriteHeader(http.StatusOK)

	var comment comment.Comment
	if err := json.NewDecoder(request.Body).Decode(&comment); err != nil {
		sendErrorResponse(writer, "Failed to decode JSON body", err)
	}

	vars := mux.Vars(request)
	id := vars["id"]
	commentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(writer, "Failed to parse uint from ID", err)
	}

	comment, err = handler.Service.UpdateComment(uint(commentID), comment)

	if err != nil {
		sendErrorResponse(writer, "Failed to update comment", err)
	}
	if err := json.NewEncoder(writer).Encode(comment); err != nil {
		panic(err)
	}
}

// DeleteComment - deletes a comment by ID
func (handler *Handler) DeleteComment(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.WriteHeader(http.StatusOK)

	vars := mux.Vars(request)
	id := vars["id"]
	commentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(writer, "Failed to parse uint from ID", err)
	}

	err = handler.Service.DeleteComment(uint(commentID))
	if err != nil {
		sendErrorResponse(writer, "Failed to delete comment by ID", err)
	}
	if err := json.NewEncoder(writer).Encode(Response{Message: "Comment successfully deleted"}); err != nil {
		panic(err)
	}
}
