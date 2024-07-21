package server

import (
	"database/sql"
	"log"
	"net/http"
)

func StartServer(db *sql.DB) {
	router := newRouter()

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Starting server at address :8080")
	server.ListenAndServe()
}
