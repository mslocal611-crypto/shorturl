package main

import (
	"log/slog"
	"net"
	"net/http"
	"os"

	"github.com/mslocal611-crypto/shorturl/internal/config"
	"github.com/mslocal611-crypto/shorturl/internal/handler"
)

func main() {
	cfg := config.Load()

	addr := net.JoinHostPort(cfg.Host, cfg.Port)
	baseURL := "http://" + addr

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	logger.Info("server started", "baseURL", baseURL)
	err := http.ListenAndServe(addr, handler.Route(baseURL))
	if err != nil {
		logger.Error("server stopped", "error", err)
		os.Exit(1)
	}
}
