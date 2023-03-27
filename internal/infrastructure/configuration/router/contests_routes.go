package router

import (
	handler "azl-vote-api/internal/domain/handlers"

	"github.com/go-chi/chi/v5"
)

func SetContestsRoutes(router *chi.Mux, contestsHandlers handler.ContestsHandlers) {
	router.Get("/contests", contestsHandlers.ListContestHandler)
	router.Post("/contests", contestsHandlers.CreateContestHandler)
}
