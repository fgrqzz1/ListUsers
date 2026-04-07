package handlers

import (
	"ListUsers/internal/repository"
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	Repo repository.UserRepository
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, err := h.Repo.GetUsers()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}
