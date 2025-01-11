package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.GetHead)
	r.Use(middleware.Logger)

	r.Get("/version", app.GetVersionHandler)
	return r
}
