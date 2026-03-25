package response

import (
	"easy-ride-api/pkg/validate"
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool                  `json:"success"`
	Message string                `json:"message"`
	Data    any                   `json:"data,omitempty"`
	Errors  []validate.FieldError `json:"errors,omitempty"`
}

func WriteJsonResponse(w http.ResponseWriter, statusCode int, data Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
