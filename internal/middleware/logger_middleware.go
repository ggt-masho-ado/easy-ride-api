package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		log.Printf("Method: %s Path: %s TimeTaken: %v", r.Method, r.URL.Path, time.Since(start))
	}
}
