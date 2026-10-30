package main

import (
	"net/http"

	"github.com/mslocal611-crypto/shorturl/internal/handler"
	"github.com/mslocal611-crypto/shorturl/internal/repository/inmemory"
	"github.com/mslocal611-crypto/shorturl/internal/service"
)

func main() {
	mux := http.NewServeMux()

	service := service.New(inmemory.New())
	handler := handler.New(service)

	mux.HandleFunc("POST /", handler.Shorten)
	mux.HandleFunc("GET /{shortUrl}", handler.Redirect)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
