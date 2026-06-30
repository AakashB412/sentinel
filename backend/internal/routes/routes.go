package routes

import (
	"github.com/go-chi/chi/v5"

	"github.com/AakashB412/sentinel/backend/internal/handlers"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", handlers.Health)

	return r
}
