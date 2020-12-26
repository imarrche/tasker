package inmem

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/imarrche/tasker/internal/model"
	"github.com/imarrche/tasker/internal/store"
)

func TestColumnRepo_GetByProjectID(t *testing.T) {
	s := TestStoreWithFixtures()

	cs, err := s.Columns().GetByProjectID(1)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(cs))
}

func TestColumnRepo_Create(t *testing.T) {
	s := TestStoreWithFixtures()

	c, err := s.Columns().Create(model.Column{Name: "Column 4", ProjectID: 2})

	assert.NoError(t, err)
	assert.Equal(t, len(s.db.columns), c.ID)
	assert.Equal(t, "Column 4", c.Name)
}

func TestColumnRepo_GetByID(t *testing.T) {
	s := TestStoreWithFixtures()

	c, err1 := s.Columns().GetByID(1)
	_, err2 := s.Columns().GetByID(4)

	assert.NoError(t, err1)
	assert.Equal(t, 1, c.ID)
	assert.Equal(t, store.ErrNotFound, err2)
}

func TestColumnRepo_GetByIndexAndProjectID(t *testing.T) {
	s := TestStoreWithFixtures()

	c, err1 := s.Columns().GetByIndexAndProjectID(1, 1)
	_, err2 := s.Columns().GetByIndexAndProjectID(2, 2)

	assert.NoError(t, err1)
	assert.Equal(t, 1, c.ID)
	assert.Equal(t, store.ErrNotFound, err2)
}

func TestColumnRepo_Update(t *testing.T) {
	s := TestStoreWithFixtures()

	c1, err1 := s.Columns().Update(model.Column{ID: 1, Name: "Updated column 1"})
	_, err2 := s.Columns().Update(model.Column{ID: 4})

	assert.NoError(t, err1)
	assert.Equal(t, "Updated column 1", c1.Name)
	assert.Equal(t, store.ErrNotFound, err2)
}

func TestColumnRepo_DeleteByID(t *testing.T) {
	s := TestStoreWithFixtures()

	err1 := s.Columns().DeleteByID(1)
	err2 := s.Columns().DeleteByID(4)

	assert.NoError(t, err1)
	assert.Equal(t, 2, len(s.db.columns))
	assert.Equal(t, 1, len(s.db.tasks))
	assert.Equal(t, 0, len(s.db.comments))
	assert.Equal(t, store.ErrNotFound, err2)
}
