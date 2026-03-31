package handlers

import (
	"easy-ride-api/pkg/response"
	"net/http"
)

func ListUsersHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.WriteJsonResponse(w, http.StatusOK, response.Response{
			Success: true,
			Message: "List users",
			Data:    make([]int, 0),
		})
	}
}
