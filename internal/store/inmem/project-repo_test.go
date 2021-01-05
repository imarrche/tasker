package inmem

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/imarrche/tasker/internal/model"
)

func TestProjectRepo_GetAll(t *testing.T) {
	s := TestStoreWithFixtures()

	ps, err := s.Projects().GetAll()

	assert.NoError(t, err)
	assert.Equal(t, 2, len(ps))
}

func TestProjectRepo_Create(t *testing.T) {
	s := TestStoreWithFixtures()

	p, err := s.Projects().Create(model.Project{Name: "Project 3"})

	assert.NoError(t, err)
	assert.Equal(t, model.Project{ID: 3, Name: "Project 3"}, p)
}

func TestProjectRepo_GetByID(t *testing.T) {
	s := TestStoreWithFixtures()

	p, err := s.Projects().GetByID(1)

	assert.NoError(t, err)
	assert.Equal(t, 1, p.ID)
}

func TestProjectRepo_Update(t *testing.T) {
	s := TestStoreWithFixtures()
	project := model.Project{ID: 1, Name: "Updated project 1"}

	p, err := s.Projects().Update(project)

	assert.NoError(t, err)
	assert.Equal(t, project, p)
}

func TestProjectRepo_DeleteByID(t *testing.T) {
	s := TestStoreWithFixtures()

	err := s.Projects().DeleteByID(1)

	assert.NoError(t, err)
	assert.Equal(t, 1, len(s.db.projects))
	assert.Equal(t, 1, len(s.db.columns))
	assert.Equal(t, 0, len(s.db.tasks))
	assert.Equal(t, 0, len(s.db.comments))
}
