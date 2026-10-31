package service

import (
	"testing"

	"github.com/mslocal611-crypto/shorturl/internal/errors"
	"github.com/mslocal611-crypto/shorturl/internal/repository/inmemory"
	"github.com/stretchr/testify/assert"
)

func TestService_Put(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		s    *Service
		args args
		err  error
	}{
		{
			name: "test valid url",
			s:    New(inmemory.New()),
			args: args{url: "https://example.com"},
			err:  nil,
		},
		{
			name: "test invalid url",
			s:    New(inmemory.New()),
			args: args{url: "example.com"},
			err:  errors.ErrInvalidURL,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Put(tt.args.url)

			if tt.err != nil {
				assert.ErrorIs(t, err, tt.err)
				assert.Empty(t, got)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, encode(tt.args.url), got)
		})
	}
}

func TestService_Get(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name string
		s    *Service
		args args
		want string
		err  error
	}{
		{
			name: "test valid code",
			s:    New(inmemory.New()),
			args: args{code: encode("https://example.com")},
			want: "https://example.com",
			err:  nil,
		},
		{
			name: "test not found code",
			s:    New(inmemory.New()),
			args: args{code: "invalid"},
			want: "",
			err:  errors.ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.err != nil {
				got, err := tt.s.Get(tt.args.code)
				assert.ErrorIs(t, err, tt.err)
				assert.Empty(t, got)
				return
			}

			code, err := tt.s.Put(tt.want)
			assert.NoError(t, err)
			assert.Equal(t, tt.args.code, code)

			got, err := tt.s.Get(code)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_encode(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test str encoding",
			args: args{url: "https://example.com"},
			want: "100680ad",
		},
		{
			name: "test empty encoding",
			args: args{url: ""},
			want: "e3b0c442",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, encode(tt.args.url))
			assert.Equal(t, len(tt.want), len(encode(tt.args.url)))
			assert.Equal(t, encode(tt.args.url), encode(tt.args.url))
		})
	}

	t.Run("test different urls encoding", func(t *testing.T) {
		assert.NotEqual(t, encode("https://example.com"), encode("https://example.org"))
	})
}

func Test_validate(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want bool
	}{
		{
			name: "valid url",
			str:  "https://example.com",
			want: true,
		},
		{
			name: "invalid url",
			str:  "example.com",
			want: false,
		},
		{
			name: "empty url",
			str:  "",
			want: false,
		},
		{
			name: "invalid scheme",
			str:  "ftp://example.com",
			want: false,
		},
		{
			name: "invalid host",
			str:  "https://",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validate(tt.str)
			assert.Equal(t, tt.want, got)
		})
	}
}
