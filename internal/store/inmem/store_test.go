package inmem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStore_Open(t *testing.T) {
	assert.NoError(t, NewStore().Open())
}

func TestStore_Close(t *testing.T) {
	s := NewStore()
	s.Open()

	assert.NoError(t, s.Close())
}
