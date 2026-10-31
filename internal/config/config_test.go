package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	serverAddress := "localhost:8080"
	baseURL := "http://localhost:8080"

	t.Setenv("SERVER_ADDRESS", serverAddress)
	t.Setenv("BASE_URL", baseURL)

	cfg, err := Load([]string{})

	assert.NoError(t, err)
	assert.Equal(t, serverAddress, cfg.HTTPAddr)
	assert.Equal(t, baseURL, cfg.BaseURL)

	serverAddress = "localhost:8081"
	baseURL = "http://localhost:8081"

	cfg, err = Load([]string{"-a", serverAddress, "-b", baseURL})

	assert.NoError(t, err)
	assert.Equal(t, serverAddress, cfg.HTTPAddr)
	assert.Equal(t, baseURL, cfg.BaseURL)
}
