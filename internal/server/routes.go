package server

import (
	"net/http"

	"github.com/evlbit/notesmd/internal/handlers"
)

func newRouter() http.Handler {
	router := http.NewServeMux()

	notesHandler := handlers.NewNotesHandler()
	notesHandler.RegisterRoutes(router)

	return router
}
