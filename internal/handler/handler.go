package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mslocal611-crypto/shorturl/internal/errors"
	"github.com/mslocal611-crypto/shorturl/internal/service"
)

type Handler struct {
	service *service.Service
}

func New(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Shorten(w http.ResponseWriter, r *http.Request) {
	original := r.FormValue("url")

	code, err := h.service.Put(original)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)

	_, err = fmt.Fprintf(w, "http://localhost:8080/%s", code)
	if err != nil {
		log.Fatalf("failed to write response: %v", err)
	}
}

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("shortUrl")
	if code == "" {
		http.Error(w, errors.ErrInvalidURL.Error(), http.StatusBadRequest)
		return
	}

	original, err := h.service.Get(code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, original, http.StatusTemporaryRedirect)
}
