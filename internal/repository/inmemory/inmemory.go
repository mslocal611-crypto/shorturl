package inmemory

import (
	"github.com/mslocal611-crypto/shorturl/internal/errors"
	"github.com/mslocal611-crypto/shorturl/internal/model"
)

type Memory struct {
	urlMap map[string]string
}

func New() *Memory {
	return &Memory{
		urlMap: make(map[string]string),
	}
}

func (m *Memory) Save(link model.URL) error {
	m.urlMap[link.Code] = link.Original

	return nil
}

func (m *Memory) Load(code string) (string, error) {
	original, ok := m.urlMap[code]
	if !ok {
		return "", errors.ErrNotFound
	}

	return original, nil
}
