package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (route *application) routes() *chi.Mux {

	// middleware must come before any routes

	// add routes here
	route.get("/", route.Handlers.Home)

	// static routes for images and other static files
	fileServer := http.FileServer(http.Dir("./public"))
	route.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileServer))

	return route.App.Routes
}
