package server

import "net/http"

func newRouter() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("GET /notes", handleGetNotes)

	return router
}

func handleGetNotes(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("A list of all notes"))
}
