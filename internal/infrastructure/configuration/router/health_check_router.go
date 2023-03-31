package router

import (
	handler "azl-vote-api/internal/domain/handlers"

	"github.com/go-chi/chi/v5"
)

func SetHealthCheckRoutes(router *chi.Mux, healthCheckHandlers handler.HealthCheckHandlers) {
	router.Get("/health", healthCheckHandlers.CheckAPIHealth)
}
