package handler

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoute(t *testing.T) {
	mux := Route(testBaseURL)
	assert.NotNil(t, mux)

	request := httptest.NewRequest(http.MethodGet, "/missing", nil)
	writer := httptest.NewRecorder()

	mux.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusBadRequest, writer.Code)

	request = httptest.NewRequest(http.MethodPost, "/", nil)
	writer = httptest.NewRecorder()

	mux.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusBadRequest, writer.Code)

	body := url.Values{
		"url": {"https://www.google.com"},
	}
	request = httptest.NewRequest(http.MethodPost, "/", strings.NewReader(body.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	writer = httptest.NewRecorder()

	mux.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusCreated, writer.Code)
}
