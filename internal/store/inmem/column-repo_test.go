package inmem

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/imarrche/tasker/internal/model"
)

func TestColumnRepo_GetByProjectID(t *testing.T) {
	s := TestStoreWithFixtures()

	cs, err := s.Columns().GetByProjectID(1)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(cs))
}

func TestColumnRepo_Create(t *testing.T) {
	s := TestStoreWithFixtures()

	c, err := s.Columns().Create(model.Column{Name: "Column 4", Index: 2, ProjectID: 2})

	assert.NoError(t, err)
	assert.Equal(t, model.Column{ID: 4, Name: "Column 4", Index: 2, ProjectID: 2}, c)
}

func TestColumnRepo_GetByID(t *testing.T) {
	s := TestStoreWithFixtures()

	c, err := s.Columns().GetByID(1)

	assert.NoError(t, err)
	assert.Equal(t, 1, c.ID)
}

func TestColumnRepo_GetByIndexAndProjectID(t *testing.T) {
	s := TestStoreWithFixtures()

	c, err := s.Columns().GetByIndexAndProjectID(1, 1)

	assert.NoError(t, err)
	assert.Equal(t, 1, c.ID)
}

func TestColumnRepo_Update(t *testing.T) {
	s := TestStoreWithFixtures()
	column := model.Column{ID: 1, Name: "Updated column 1", Index: 1, ProjectID: 1}

	c, err := s.Columns().Update(column)

	assert.NoError(t, err)
	assert.Equal(t, column, c)
}

func TestColumnRepo_DeleteByID(t *testing.T) {
	s := TestStoreWithFixtures()

	err := s.Columns().DeleteByID(1)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(s.db.columns))
	assert.Equal(t, 1, len(s.db.tasks))
	assert.Equal(t, 0, len(s.db.comments))
}
