package memory

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/imarrche/tasker/internal/model"
	"github.com/imarrche/tasker/internal/store"
)

func TestColumnRepository_GetAll(t *testing.T) {
	s := NewStore()
	c1 := model.Column{Name: "Column 1"}
	c2 := model.Column{Name: "Column 2"}
	s.db.columns[1] = c1
	s.db.columns[2] = c2

	cs, err := s.Columns().GetAll()

	assert.NoError(t, err)
	assert.Equal(t, 2, len(cs))
}

func TestColumnRepository_GetAllByProjectID(t *testing.T) {
	s := NewStore()
	p1 := model.Project{ID: 1, Name: "Project1 "}
	c1 := model.Column{Name: "Column 1", Project: p1}
	c2 := model.Column{Name: "Column 2"}
	s.db.columns[1] = c1
	s.db.columns[2] = c2

	cs, err := s.Columns().GetAllByProjectID(p1.ID)

	assert.NoError(t, err)
	assert.Equal(t, 1, len(cs))
	assert.Equal(t, c1, cs[0])
}

func TestColumnRepository_Create(t *testing.T) {
	s := NewStore()
	c1 := model.Column{Name: "Column 1"}
	c2 := model.Column{Name: "Column 2"}

	c1FromRepo, err1 := s.Columns().Create(c1)
	c2FromRepo, err2 := s.Columns().Create(c2)

	assert.NoError(t, err1)
	assert.Equal(t, 1, c1FromRepo.ID)
	assert.Equal(t, c1.Name, c1FromRepo.Name)
	assert.NoError(t, err2)
	assert.Equal(t, 2, c2FromRepo.ID)
	assert.Equal(t, c2.Name, c2FromRepo.Name)
}

func TestColumnRepository_GetByID(t *testing.T) {
	s := NewStore()
	c1 := model.Column{ID: 1, Name: "Column 1"}
	s.db.columns[c1.ID] = c1

	c1FromRepo, err1 := s.Columns().GetByID(c1.ID)
	_, err2 := s.Columns().GetByID(2)

	assert.NoError(t, err1)
	assert.Equal(t, c1.ID, c1FromRepo.ID)
	assert.Equal(t, c1.Name, c1FromRepo.Name)
	assert.Equal(t, store.ErrNotFound, err2)
}

func TestColumnRepository_Update(t *testing.T) {
	s := NewStore()
	c1 := model.Column{ID: 1, Name: "Column 1"}
	c2 := model.Column{ID: 2, Name: "Column 2"}
	s.db.columns[c1.ID] = c1

	c1.Name = "Updated name"
	err1 := s.Columns().Update(c1)
	err2 := s.Columns().Update(c2)

	assert.NoError(t, err1)
	assert.Equal(t, c1.Name, s.db.columns[c1.ID].Name)
	assert.Equal(t, store.ErrNotFound, err2)
}

func TestColumnRepository_Delete(t *testing.T) {
	s := NewStore()
	c1 := model.Column{ID: 1, Name: "Column 1"}
	c2 := model.Column{ID: 2, Name: "Column 2"}
	s.db.columns[c1.ID] = c1

	err1 := s.Columns().DeleteByID(c1.ID)
	err2 := s.Columns().DeleteByID(c2.ID)

	assert.NoError(t, err1)
	assert.Equal(t, 0, len(s.db.columns))
	assert.Equal(t, store.ErrNotFound, err2)
}
