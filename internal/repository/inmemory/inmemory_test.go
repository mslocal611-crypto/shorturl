package inmemory

import (
	"testing"

	"github.com/mslocal611-crypto/shorturl/internal/errors"
	"github.com/mslocal611-crypto/shorturl/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestMemory(t *testing.T) {
	store := New()

	got, err := store.Load("undefined_hash")
	assert.ErrorIs(t, err, errors.ErrNotFound)
	assert.Empty(t, got)

	link := model.URL{
		Original: "some_url",
		Code:     "some_hash",
	}
	err = store.Save(link)
	assert.NoError(t, err)

	got, err = store.Load(link.Code)
	assert.NoError(t, err)
	assert.Equal(t, link.Original, got)
}
