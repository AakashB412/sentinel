package server

import (
	"net/http"

	"github.com/AakashB412/sentinel/backend/internal/routes"
)

func New() *http.Server {
	return &http.Server{
		Addr:    ":8080",
		Handler: routes.NewRouter(),
	}
}
