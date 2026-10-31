package handler

import (
	"net/http"

	"github.com/mslocal611-crypto/shorturl/internal/repository/inmemory"
	"github.com/mslocal611-crypto/shorturl/internal/service"
)

func Route() *http.ServeMux {
	mux := http.NewServeMux()

	service := service.New(inmemory.New())
	handler := New(service)

	mux.HandleFunc("POST /", handler.Shorten)
	mux.HandleFunc("GET /{shortUrl}", handler.Redirect)

	return mux
}
