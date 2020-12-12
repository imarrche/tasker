package memory

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/imarrche/tasker/internal/model"
	"github.com/imarrche/tasker/internal/store"
)

func TestProjectRepository_GetAll(t *testing.T) {
	s := NewStore()
	p1 := model.Project{Name: "Project 1"}
	p2 := model.Project{Name: "Project 2"}
	s.db.projects[1] = p1
	s.db.projects[2] = p2

	ps, err := s.Projects().GetAll()

	assert.NoError(t, err)
	assert.Equal(t, 2, len(ps))
}

func TestProjectRepository_Create(t *testing.T) {
	s := NewStore()
	p1 := model.Project{Name: "Project 1"}
	p2 := model.Project{Name: "Project 2"}

	p1FromRepo, err1 := s.Projects().Create(p1)
	p2FromRepo, err2 := s.Projects().Create(p2)

	assert.NoError(t, err1)
	assert.Equal(t, 1, p1FromRepo.ID)
	assert.Equal(t, p1.Name, p1FromRepo.Name)
	assert.NoError(t, err2)
	assert.Equal(t, 2, p2FromRepo.ID)
	assert.Equal(t, p2.Name, p2FromRepo.Name)
}

func TestProjectRepository_GetByID(t *testing.T) {
	s := NewStore()
	p1 := model.Project{ID: 1, Name: "Project 1"}
	s.db.projects[p1.ID] = p1

	p1FromRepo, err1 := s.Projects().GetByID(p1.ID)
	_, err2 := s.Projects().GetByID(2)

	assert.NoError(t, err1)
	assert.Equal(t, p1.ID, p1FromRepo.ID)
	assert.Equal(t, p1.Name, p1FromRepo.Name)
	assert.Equal(t, store.ErrNotFound, err2)
}

func TestProjectRepository_Update(t *testing.T) {
	s := NewStore()
	p1 := model.Project{ID: 1, Name: "Project 1"}
	p2 := model.Project{ID: 2, Name: "Project 2"}
	s.db.projects[p1.ID] = p1

	p1.Name = "Updated name"
	err1 := s.Projects().Update(p1)
	err2 := s.Projects().Update(p2)

	assert.NoError(t, err1)
	assert.Equal(t, p1.Name, s.db.projects[p1.ID].Name)
	assert.Equal(t, store.ErrNotFound, err2)
}

func TestProjectRepository_Delete(t *testing.T) {
	s := NewStore()
	p1 := model.Project{ID: 1, Name: "Project 1"}
	p2 := model.Project{ID: 2, Name: "Project 2"}
	s.db.projects[p1.ID] = p1

	err1 := s.Projects().Delete(p1)
	err2 := s.Projects().Delete(p2)

	assert.NoError(t, err1)
	assert.Equal(t, 0, len(s.db.projects))
	assert.Equal(t, store.ErrNotFound, err2)
}
