package handlers

import (
	"easy-ride-api/internal/actions"
	"easy-ride-api/internal/services"
	"easy-ride-api/pkg/response"
	"easy-ride-api/pkg/validate"
	"encoding/json"
	"net/http"
)

func SignInHandler(userService services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var body actions.UserSignIn

		//decode request json data to a struct
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {

			data := response.Response{
				Success: false,
				Message: "Invalid request body",
			}

			response.WriteJsonResponse(w, http.StatusBadRequest, data)

			return
		}

		//validate struct
		if err := validate.ValidateStruct(body); err != nil {

			data := response.Response{
				Success: false,
				Message: "Login failed",
				Errors:  err,
			}

			response.WriteJsonResponse(w, http.StatusUnprocessableEntity, data)

			return
		}

		//user session creation and wiring up response
		session, err := userService.CreateUserSession(r.Context(), body.Email, body.Password)

		if err != nil {
			data := response.Response{
				Success: false,
				Message: err.Error(),
			}

			response.WriteJsonResponse(w, http.StatusInternalServerError, data)

			return
		}

		data := response.Response{
			Success: true,
			Message: "Login is succesfull",
			Data:    session,
		}

		response.WriteJsonResponse(w, http.StatusCreated, data)

	}
}
