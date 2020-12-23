package inmem

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/imarrche/tasker/internal/model"
	"github.com/imarrche/tasker/internal/store"
)

func TestProjectRepository_GetAll(t *testing.T) {
	s := TestStoreWithFixtures()

	ps, err := s.Projects().GetAll()

	assert.NoError(t, err)
	assert.Equal(t, len(s.db.projects), len(ps))
}

func TestProjectRepository_Create(t *testing.T) {
	s := TestStoreWithFixtures()

	p1, err1 := s.Projects().Create(model.Project{Name: "Project 3"})

	assert.NoError(t, err1)
	assert.Equal(t, len(s.db.projects), p1.ID)
	assert.Equal(t, "Project 3", p1.Name)
}

func TestProjectRepository_GetByID(t *testing.T) {
	s := TestStoreWithFixtures()

	p1, err1 := s.Projects().GetByID(1)
	_, err2 := s.Projects().GetByID(3)

	assert.NoError(t, err1)
	assert.Equal(t, 1, p1.ID)
	assert.Equal(t, store.ErrNotFound, err2)
}

func TestProjectRepository_Update(t *testing.T) {
	s := TestStoreWithFixtures()

	p1, err1 := s.Projects().Update(model.Project{ID: 1, Name: "Updated project 1"})
	_, err2 := s.Projects().Update(model.Project{ID: 3})

	assert.NoError(t, err1)
	assert.Equal(t, model.Project{ID: 1, Name: "Updated project 1"}, p1)
	assert.Equal(t, "Updated project 1", p1.Name)
	assert.Equal(t, store.ErrNotFound, err2)
}

func TestProjectRepository_DeleteByID(t *testing.T) {
	s := TestStoreWithFixtures()

	err1 := s.Projects().DeleteByID(1)
	err2 := s.Projects().DeleteByID(3)

	assert.NoError(t, err1)
	assert.Equal(t, 1, len(s.db.projects))
	assert.Equal(t, 1, len(s.db.columns))
	assert.Equal(t, 0, len(s.db.tasks))
	assert.Equal(t, 0, len(s.db.comments))
	assert.Equal(t, store.ErrNotFound, err2)
}
