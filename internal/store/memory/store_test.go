package memory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStore_Open(t *testing.T) {
	assert.NoError(t, NewStore().Open())
}

func TestStore_Close(t *testing.T) {
	s := NewStore()
	err := s.Open()

	assert.NoError(t, err)
	assert.NoError(t, s.Close())
}
