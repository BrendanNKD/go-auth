package service

import (
	"auth-server/internal/auth/model"
	"database/sql"
	"encoding/json"
	"net/http"
)

type AuthService struct {
	db *sql.DB
}

func NewAuthService(db *sql.DB) *AuthService {
	return &AuthService{db: db}
}

func (s *AuthService) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Here, you would verify the user credentials (check username and password).
	// This is just a placeholder response.
	json.NewEncoder(w).Encode("Login successful")
}

func (s *AuthService) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Here, you would save the user to the database.
	// This is just a placeholder response.
	json.NewEncoder(w).Encode("Registration successful")
}
