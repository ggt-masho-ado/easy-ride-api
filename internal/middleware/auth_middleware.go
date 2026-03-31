package middleware

import (
	"context"
	"easy-ride-api/internal/services"
	"easy-ride-api/pkg/response"
	"net/http"
	"strings"
)

type contextKey string

const UserIDCtxKey contextKey = "auth_user"

func AuthMiddleware(sessionService services.SessionService) func(next http.Handler) http.HandlerFunc {
	return func(next http.Handler) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			bearerToken := r.Header.Get("Authorization")

			parts := strings.Split(bearerToken, " ")

			if len(parts) != 2 || parts[0] != "Bearer" {
				response.WriteJsonResponse(w, http.StatusUnauthorized, response.Response{
					Message: "Unauthorized",
					Success: false,
				})
				return
			}

			token := parts[1]

			session, err := sessionService.GetSession(r.Context(), token)
			if err != nil {
				response.WriteJsonResponse(w, http.StatusUnauthorized, response.Response{
					Message: "Unauthorized",
					Success: false,
				})
				return
			}

			ctx := context.WithValue(r.Context(), UserIDCtxKey, session.UserId)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	}
}
