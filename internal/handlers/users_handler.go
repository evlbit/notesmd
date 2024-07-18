package handlers

import (
	"net/http"

	"github.com/evlbit/notesmd/internal/types/request"
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
	var payload request.RegisterUserPayload
	if err := parseJSON(r, &payload); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	//TODO: handle user registration

	writeJSON(w, http.StatusCreated, nil)
}

func (h *UsersHandler) handleLogin(w http.ResponseWriter, r *http.Request) {
	var payload request.LoginUserPayload
	if err := parseJSON(r, &payload); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	//TODO: handle user login

	writeJSON(w, http.StatusOK, nil)
}
