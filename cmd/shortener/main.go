package main

import (
	"net/http"

	"github.com/mslocal611-crypto/shorturl/internal/handler"
)

func main() {
	err := http.ListenAndServe(":8080", handler.Route())
	if err != nil {
		panic(err)
	}
}
