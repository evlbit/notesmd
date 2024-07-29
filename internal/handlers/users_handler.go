package handlers

import (
	"net/http"

	"github.com/evlbit/notesmd/internal/types"
)

type UsersHandler struct {
}

func NewUsersHandler() *UsersHandler {
	return &UsersHandler{}
}

func (h *UsersHandler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /register", h.handleRegister)
	router.HandleFunc("POST /login", h.handleLogin)
}

func (h *UsersHandler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterUserPayload
	if err := parseJSON(r, &payload); err != nil {
		WriteError(w, http.StatusBadRequest, err)
		return
	}

	//TODO: handle user registration

	writeJSON(w, http.StatusCreated, nil)
}

func (h *UsersHandler) handleLogin(w http.ResponseWriter, r *http.Request) {
	var payload types.LoginUserPayload
	if err := parseJSON(r, &payload); err != nil {
		WriteError(w, http.StatusBadRequest, err)
		return
	}

	//TODO: handle user login

	writeJSON(w, http.StatusOK, nil)
}
