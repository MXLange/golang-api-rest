package routes

import (
	userRoutes "api-crud/internal/entities/user/routes"

	"github.com/go-chi/chi/v5"
)

func Router(r *chi.Mux) {
	r.Route("/user", userRoutes.UserRouter)
}