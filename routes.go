package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

// routes handles the application routes
func routes(app App) http.Handler {

	mux := chi.NewRouter()

	mux.Use(middleware.RealIP)
	mux.Use(middleware.Recoverer)

	mux.Get("/{action}", ReportStatus(app))
	mux.Post("/{action}", ReportStatus(app))

	return mux
}
