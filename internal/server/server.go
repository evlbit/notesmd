package server

import (
	"log"
	"net/http"
)

func StartServer() {
	router := newRouter()

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Starting server at address :8080")
	server.ListenAndServe()
}
