package service

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"net/url"

	"github.com/mslocal611-crypto/shorturl/internal/model"
)

type Repository interface {
	Save(url model.URL)
	Load(code string) (string, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Put(url string) (string, error) {
	if !validate(url) {
		return "", errors.New("параметр url должен быть валидным http адресом")
	}

	code := encode(url)
	s.repo.Save(model.URL{Original: url, Code: code})

	return code, nil
}

func (s *Service) Get(code string) (string, error) {
	return s.repo.Load(code)
}

func encode(url string) string {
	hash := sha256.Sum256([]byte(url))
	return hex.EncodeToString(hash[:])[:8]
}

func validate(str string) bool {
	parsed, err := url.ParseRequestURI(str)
	if err != nil {
		return false
	}

	return (parsed.Scheme == "http" || parsed.Scheme == "https") && parsed.Host != ""
}
