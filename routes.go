package main

import (
	"github.com/go-chi/chi"
	"net/http"
)

func routes(app App) http.Handler {

	mux := chi.NewRouter()

	mux.Get("/{action}", ReportStatus(app))

	return mux
}
