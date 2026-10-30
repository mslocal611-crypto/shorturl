package handler

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/mslocal611-crypto/shorturl/internal/errors"
	"github.com/mslocal611-crypto/shorturl/internal/repository/inmemory"
	"github.com/mslocal611-crypto/shorturl/internal/service"
	"github.com/stretchr/testify/assert"
)

func TestShorten(t *testing.T) {
	type in struct {
		body string
	}
	type out struct {
		statusCode  int
		contentType string
		body        string
	}

	invalid := out{
		statusCode:  http.StatusBadRequest,
		contentType: "text/plain; charset=utf-8",
		body:        errors.ErrInvalidURL.Error() + "\n",
	}

	tests := []struct {
		name string
		in   in
		out  out
	}{
		{
			name: "valid form",
			in: in{
				body: url.Values{"url": {"http://example.com"}}.Encode(),
			},
			out: out{
				statusCode:  http.StatusCreated,
				contentType: "text/plain",
				body:        "http://localhost:8080/f0e6a6a9",
			},
		},
		{
			name: "invalid form",
			in: in{
				body: url.Values{"textUrl": {"http://example.com"}}.Encode(),
			},
			out: invalid,
		},
		{
			name: "empty url",
			in: in{
				body: url.Values{"url": {""}}.Encode(),
			},
			out: invalid,
		},
		{
			name: "invalid url",
			in: in{
				body: url.Values{"url": {"ftp://example.com/path/to/file"}}.Encode(),
			},
			out: invalid,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(test.in.body))
			request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			writer := httptest.NewRecorder()

			handler := New(service.New(inmemory.New()))
			handler.Shorten(writer, request)

			assert.Equal(t, test.out.statusCode, writer.Result().StatusCode)
			assert.Equal(t, test.out.contentType, writer.Result().Header.Get("Content-Type"))
			assert.Equal(t, test.out.body, writer.Body.String())
		})
	}
}

func TestHandler_Redirect(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		service *service.Service
		// Named input parameters for target function.
		w http.ResponseWriter
		r *http.Request
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(tt.service)
			h.Redirect(tt.w, tt.r)
		})
	}
}
