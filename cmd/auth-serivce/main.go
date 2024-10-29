package main

import (
	"auth-server/internal/database"
	"auth-server/internal/user/service"
	"context"
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	// Connect to MongoDB
	mongoClient, err := database.ConnectMongoDB()
	userService := service.NewUserService()

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		users, err := userService.GetAllUsers()
		if err != nil {
			http.Error(w, "Error retrieving users", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	})

	if err != nil {
		log.Fatalf("Could not connect to MongoDB: %v", err)
	}
	defer func() {
		if err := mongoClient.Disconnect(context.Background()); err != nil {
			log.Fatalf("Could not disconnect MongoDB: %v", err)
		}
	}()

	// Set up routes, services, etc.
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
