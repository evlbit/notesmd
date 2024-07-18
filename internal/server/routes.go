package server

import (
	"net/http"

	"github.com/evlbit/notesmd/internal/handlers"
	"github.com/evlbit/notesmd/internal/middleware"
)

func newRouter() http.Handler {
	router := http.NewServeMux()

	notesHandler := handlers.NewNotesHandler()
	notesHandler.RegisterRoutes(router)

	usersHandler := handlers.NewUsersHandler()
	usersHandler.RegisterRoutes(router)

	middlewareStack := middleware.CreateStack(
		middleware.Logging,
	)

	return middlewareStack(router)
}
