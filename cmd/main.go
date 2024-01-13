package main

import (
	"api-crud/internal/shared/routes"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	mainRouter := chi.NewRouter();
	mainRouter.Use(middleware.Logger)
	routes.Router(mainRouter)
	http.ListenAndServe(":5555", mainRouter)
}