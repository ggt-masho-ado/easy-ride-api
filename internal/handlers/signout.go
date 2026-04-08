package handlers

import (
	"easy-ride-api/internal/services"
	"easy-ride-api/pkg/logger"
	"easy-ride-api/pkg/response"
	"net/http"
	"strings"
)

func SignOut(userService services.UserService) http.HandlerFunc {
	return (func(w http.ResponseWriter, r *http.Request) {

		bearerToken := r.Header.Get("Authorization")

		parts := strings.Split(bearerToken, " ")

		token := parts[1]

		err := userService.InvalidateUserSession(r.Context(), token)

		if err != nil {
			logger.Log(err)

			response.WriteJsonResponse(w, http.StatusInternalServerError, response.Response{
				Success: false,
				Message: "An error occurred while logging out, please try again",
			})
			return
		}

		response.WriteJsonResponse(w, http.StatusOK, response.Response{
			Success: true,
			Message: "You have been successfully logged out",
		})

	})
}
