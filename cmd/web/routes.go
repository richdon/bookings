package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/richdon/bookings/pkg/handlers"
)

func routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}
