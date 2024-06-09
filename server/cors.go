package main

import (
	"net/http"

	"github.com/go-chi/cors"
)

func corsHandler() func(http.Handler) http.Handler {
	return cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://localhost:5173"},
		AllowedMethods: []string{"POST", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type"},
	})
}
