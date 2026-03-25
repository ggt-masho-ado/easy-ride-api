package handlers

import (
	"easy-ride-api/internal/actions"
	"easy-ride-api/internal/services"
	"easy-ride-api/pkg/logger"
	"easy-ride-api/pkg/response"
	"easy-ride-api/pkg/validate"
	"encoding/json"
	"net/http"
)

func SignUpHandler(userService services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var body actions.UserSignUp

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			response.WriteJsonResponse(w, http.StatusBadRequest, response.Response{
				Success: false,
				Message: "Invalid request body",
			})
			return
		}

		if validationErrs := validate.ValidateStruct(body); validationErrs != nil {
			response.WriteJsonResponse(w, http.StatusUnprocessableEntity, response.Response{
				Success: false,
				Message: "Signup failed",
				Errors:  validationErrs,
			})
			return
		}

		user, err := userService.CreateNewUser(r.Context(), body.FullName, body.Email, body.Password, body.ConfirmPassword)
		if err != nil {
			response.WriteJsonResponse(w, http.StatusInternalServerError, response.Response{
				Success: false,
				Message: err.Error(),
			})
			logger.Log(err)
			return
		}

		response.WriteJsonResponse(w, http.StatusCreated, response.Response{
			Success: true,
			Message: "User created successfully",
			Data:    user,
		})
	}
}
