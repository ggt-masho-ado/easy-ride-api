package main

import (
	"easy-ride-api/db"
	"easy-ride-api/internal/handlers"
	"easy-ride-api/internal/respositories"
	"easy-ride-api/internal/services"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	db.Connect()
	defer db.Close()

	userRepo := respositories.NewUserRepository(db.Pool)
	userService := services.NewUserService(userRepo)

	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Server is up and running")
	})

	http.HandleFunc("POST /signin", handlers.SignInHandler(userService))
	http.HandleFunc("POST /signup", handlers.SignUpHandler(userService))

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
