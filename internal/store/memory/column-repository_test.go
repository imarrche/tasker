package memory

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/imarrche/tasker/internal/model"
	"github.com/imarrche/tasker/internal/store"
)

func TestColumnRepository_GetAll(t *testing.T) {
	r := NewColumnRepository()
	c1 := model.Column{Name: "Column 1"}
	c2 := model.Column{Name: "Column 2"}
	r.store.columns[1] = c1
	r.store.columns[2] = c2

	cs, err := r.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, 2, len(cs))
}

func TestColumnRepository_Create(t *testing.T) {
	r := NewColumnRepository()
	c1 := model.Column{Name: "Column 1"}
	c2 := model.Column{Name: "Column 2"}

	c1FromRepo, err1 := r.Create(c1)
	c2FromRepo, err2 := r.Create(c2)

	assert.NoError(t, err1)
	assert.Equal(t, 1, c1FromRepo.ID)
	assert.Equal(t, c1.Name, c1FromRepo.Name)
	assert.NoError(t, err2)
	assert.Equal(t, 2, c2FromRepo.ID)
	assert.Equal(t, c2.Name, c2FromRepo.Name)
}

func TestColumnRepository_GetByID(t *testing.T) {
	r := NewColumnRepository()
	c1 := model.Column{ID: 1, Name: "Column 1"}
	r.store.columns[c1.ID] = c1

	c1FromRepo, err1 := r.GetByID(c1.ID)
	_, err2 := r.GetByID(2)

	assert.NoError(t, err1)
	assert.Equal(t, c1.ID, c1FromRepo.ID)
	assert.Equal(t, c1.Name, c1FromRepo.Name)
	assert.Equal(t, store.ErrNotFound, err2)
}

func TestColumnRepository_Update(t *testing.T) {
	r := NewColumnRepository()
	c1 := model.Column{ID: 1, Name: "Column 1"}
	c2 := model.Column{ID: 2, Name: "Column 2"}
	r.store.columns[c1.ID] = c1

	c1.Name = "Updated name"
	err1 := r.Update(c1)
	err2 := r.Update(c2)

	assert.NoError(t, err1)
	assert.Equal(t, c1.Name, r.store.columns[c1.ID].Name)
	assert.Equal(t, store.ErrNotFound, err2)
}

func TestColumnRepository_Delete(t *testing.T) {
	r := NewColumnRepository()
	c1 := model.Column{ID: 1, Name: "Column 1"}
	c2 := model.Column{ID: 2, Name: "Column 2"}
	r.store.columns[c1.ID] = c1

	err1 := r.Delete(c1)
	err2 := r.Delete(c2)

	assert.NoError(t, err1)
	assert.Equal(t, 0, len(r.store.columns))
	assert.Equal(t, store.ErrNotFound, err2)
}
