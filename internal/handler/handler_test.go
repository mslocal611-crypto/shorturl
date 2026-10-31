package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/mslocal611-crypto/shorturl/internal/repository/inmemory"
	"github.com/mslocal611-crypto/shorturl/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testBaseURL = "http://localhost:8080"

func TestHandler_Shorten(t *testing.T) {
	original := "http://example.com"
	code, err := service.New(inmemory.New()).Put(original)
	require.NoError(t, err)

	tests := []struct {
		name            string
		body            string
		wantStatusCode  int
		wantContentType string
		wantBody        string
	}{
		{
			name:            "valid url",
			body:            original,
			wantStatusCode:  http.StatusCreated,
			wantContentType: "text/plain",
			wantBody:        fmt.Sprintf("%s/%s", testBaseURL, code),
		},
		{
			name:           "invalid url",
			body:           "ftp://example.com/path/to/file",
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name:           "empty url",
			body:           "",
			wantStatusCode: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			srvc := service.New(inmemory.New())

			request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(test.body))
			request.Header.Set("Content-Type", "text/plain")

			writer := httptest.NewRecorder()

			New(srvc, testBaseURL).Shorten(writer, request)

			assert.Equal(t, test.wantStatusCode, writer.Code)

			if test.wantBody != "" && test.wantContentType != "" {
				assert.Equal(t, test.wantContentType, writer.Header().Get("Content-Type"))
				assert.Equal(t, test.wantBody, writer.Body.String())
			}
		})
	}
}

func TestHandler_Redirect(t *testing.T) {
	tests := []struct {
		name     string
		original string
		wantCode int
	}{
		{
			name:     "valid short url",
			original: "https://www.google.com",
			wantCode: http.StatusTemporaryRedirect,
		},
		{
			name:     "invalid short url",
			wantCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srvc := service.New(inmemory.New())

			code := "missing"
			if tt.original != "" {
				var err error
				code, err = srvc.Put(tt.original)
				assert.NoError(t, err)
			}

			request := httptest.NewRequest(http.MethodGet, "/"+code, nil)
			request.SetPathValue("shortUrl", code)

			writer := httptest.NewRecorder()

			New(srvc, testBaseURL).Redirect(writer, request)

			assert.Equal(t, tt.wantCode, writer.Code)
			assert.Equal(t, tt.original, writer.Header().Get("Location"))
		})
	}
}
