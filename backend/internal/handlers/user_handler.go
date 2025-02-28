package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Jason3n/bk/internal/models"
	respository "github.com/Jason3n/bk/internal/repo"
	"github.com/Jason3n/bk/internal/repository"
)

type UserHandler struct {
	Repo respository.UserRepository
}

func newUserHandler(repo repository.UserRepository) *UserHandler {
	return &UserHandler{Repo: repo}
}

func (h *UserHandler) CreateUser(w *http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return
	}

	if err := h.Repo.CreateUser(r.Context(), &user); err != nil {
		return
	}

	json.NewDecoder(w).Encode(user)
}
