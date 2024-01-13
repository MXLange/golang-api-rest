package userRoutes

import (
	userController "api-crud/internal/entities/user/controller"

	"github.com/go-chi/chi/v5"
)

func UserRouter(r chi.Router) {
	r.Get("/{id}/{name}", userController.GetUserById);
}