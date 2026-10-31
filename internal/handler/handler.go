package handler

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/mslocal611-crypto/shorturl/internal/errors"
	"github.com/mslocal611-crypto/shorturl/internal/service"
)

type Handler struct {
	service *service.Service
	baseURL string
}

func New(service *service.Service, baseURL string) *Handler {
	return &Handler{service: service, baseURL: baseURL}
}

func (h *Handler) Shorten(w http.ResponseWriter, r *http.Request) {
	original := r.FormValue("url")

	code, err := h.service.Put(original)
	if err != nil {
		slog.Error("failed to shorten URL", "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)

	_, err = fmt.Fprintf(w, "%s/%s", h.baseURL, code)
	if err != nil {
		slog.Error("failed to write response", "error", err)
		return
	}
}

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("shortUrl")
	if code == "" {
		slog.Error("invalid short URL", "code", code)
		http.Error(w, errors.ErrInvalidURL.Error(), http.StatusBadRequest)
		return
	}

	original, err := h.service.Get(code)
	if err != nil {
		slog.Error("failed to get original URL", "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, original, http.StatusTemporaryRedirect)
}
