package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"restaurant_management/internal/controllers"
)

func Routes(categoryController controllers.CategoryController) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/api", func(r chi.Router) {
		r.Route("/categories", func(r chi.Router) {
			r.Get("/", categoryController.GetAll)
			r.Get("/{categoryId}", categoryController.GetById)
		})
	})

	return r
}
