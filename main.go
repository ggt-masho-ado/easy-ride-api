package main

import (
	"fmt"
	"log"
	"net/http"

	"easy-ride-api/db"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	db.Connect()
	defer db.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello world")
	})

	log.Println("Server starting on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
