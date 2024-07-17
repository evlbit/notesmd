package handlers

import "net/http"

type NotesHandler struct {
}

func NewNotesHandler() *NotesHandler {
	return &NotesHandler{}
}

func (h *NotesHandler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /notes", h.handleGetNotes)
}

func (h *NotesHandler) handleGetNotes(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("A list of all notes"))
}
