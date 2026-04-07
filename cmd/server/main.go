package main

import (
	"ListUsers/internal/handlers"
	"ListUsers/internal/repository"
	"fmt"
	"log"
	"net/http"
)

func main() {
	repo := &repository.InMemUserRepository{}
	handler := &handlers.UserHandler{Repo: repo}

	http.HandleFunc("/users", handler.GetUsers)

	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
