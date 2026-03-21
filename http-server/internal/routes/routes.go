package routes

import (
	"github.com/aakash19here/the-go-journey/internal/app"
	"github.com/go-chi/chi/v5"
)

// orchestrate all our routes at one single place

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", app.HealthCheck)

	return r
}
