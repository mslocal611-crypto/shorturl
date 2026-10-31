package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	host := "localhost"
	port := "8080"

	t.Setenv("HOST", host)
	t.Setenv("PORT", port)

	cfg := Load()
	assert.Equal(t, host, cfg.Host)
	assert.Equal(t, port, cfg.Port)
}
