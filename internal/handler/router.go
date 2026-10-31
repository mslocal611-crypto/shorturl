package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mslocal611-crypto/shorturl/internal/repository/inmemory"
	"github.com/mslocal611-crypto/shorturl/internal/service"
)

func Route(baseURL string) http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	handler := New(service.New(inmemory.New()), baseURL)

	router.Post("/", handler.Shorten)
	router.Get("/{shortUrl}", handler.Redirect)

	return router
}
