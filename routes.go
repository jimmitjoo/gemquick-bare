package main

import (
	"fmt"
	"myapp/data"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (route *application) routes() *chi.Mux {
	// middleware must come before any routes
	route.use(route.Middleware.CheckRemember)

	// add routes here
	route.get("/", route.Handlers.Home)

	// static routes
	fileServer := http.FileServer(http.Dir("./public"))
	route.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileServer))

	return route.App.Routes
}
