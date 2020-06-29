package main

import (
	"github.com/go-chi/chi"
	"net/http"
)

// routes handles the application routes
func routes(app App) http.Handler {

	mux := chi.NewRouter()

	mux.Get("/{action}", ReportStatus(app))
	mux.Post("/{action}", ReportStatus(app))

	return mux
}
