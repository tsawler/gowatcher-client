package main

import (
	"github.com/go-chi/chi"
	"net/http"
)

func routes() http.Handler {

	mux := chi.NewRouter()

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("api"))
	})

	return mux
}
