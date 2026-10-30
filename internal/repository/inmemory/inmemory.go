package inmemory

import (
	"errors"

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

func (m *Memory) Save(link model.URL) {
	m.urlMap[link.Code] = link.Original
}

func (m *Memory) Load(code string) (string, error) {
	original, ok := m.urlMap[code]
	if !ok {
		return "", errors.New("ссылка не найдена")
	}

	return original, nil
}
