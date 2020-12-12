package memory

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/imarrche/tasker/internal/model"
	"github.com/imarrche/tasker/internal/store"
)

func TestProjectRepository_GetAll(t *testing.T) {
	r := NewProjectRepository()
	p1 := model.Project{Name: "Project 1"}
	p2 := model.Project{Name: "Project 2"}
	r.store.projects[1] = p1
	r.store.projects[2] = p2

	ps, err := r.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, 2, len(ps))
}

func TestProjectRepository_Create(t *testing.T) {
	r := NewProjectRepository()
	p1 := model.Project{Name: "Project 1"}
	p2 := model.Project{Name: "Project 2"}

	p1FromRepo, err1 := r.Create(p1)
	p2FromRepo, err2 := r.Create(p2)

	assert.NoError(t, err1)
	assert.Equal(t, p1.Name, p1FromRepo.Name)
	assert.NoError(t, err2)
	assert.Equal(t, p2.Name, p2FromRepo.Name)
}

func TestProjectRepository_GetByID(t *testing.T) {
	r := NewProjectRepository()
	p := model.Project{ID: 1, Name: "Project 1"}
	r.store.projects[p.ID] = p

	p1FromRepo, err1 := r.GetByID(p.ID)
	_, err2 := r.GetByID(2)

	assert.NoError(t, err1)
	assert.Equal(t, p.ID, p1FromRepo.ID)
	assert.Equal(t, p.Name, p1FromRepo.Name)
	assert.Equal(t, store.ErrNotFound, err2)
}

func TestProjectRepository_Update(t *testing.T) {
	r := NewProjectRepository()
	p1 := model.Project{ID: 1, Name: "Project 1"}
	p2 := model.Project{ID: 2, Name: "Project 2"}
	r.store.projects[p1.ID] = p1

	p1.Name = "Updated name"
	err1 := r.Update(p1)
	err2 := r.Update(p2)

	assert.NoError(t, err1)
	assert.Equal(t, p1.Name, r.store.projects[p1.ID].Name)
	assert.Equal(t, store.ErrNotFound, err2)
}

func TestProjectRepository_Delete(t *testing.T) {
	r := NewProjectRepository()
	p1 := model.Project{ID: 1, Name: "Project 1"}
	p2 := model.Project{ID: 2, Name: "Project 2"}
	r.store.projects[p1.ID] = p1

	err1 := r.Delete(p1)
	err2 := r.Delete(p2)

	assert.NoError(t, err1)
	assert.Equal(t, 0, len(r.store.projects))
	assert.Equal(t, store.ErrNotFound, err2)
}
