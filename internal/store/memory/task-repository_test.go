package memory

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/imarrche/tasker/internal/model"
	"github.com/imarrche/tasker/internal/store"
)

func TestTaskRepository_GetAll(t *testing.T) {
	s := NewStore()
	t1 := model.Task{Name: "Task 1"}
	t2 := model.Task{Name: "Task 2"}
	s.db.tasks[1] = t1
	s.db.tasks[2] = t2

	ts, err := s.Tasks().GetAll()

	assert.NoError(t, err)
	assert.Equal(t, 2, len(ts))
}

func TestTaskRepository_GetAllByProjectID(t *testing.T) {
	s := NewStore()
	c1 := model.Column{ID: 1, Name: "Column 1"}
	t1 := model.Task{Name: "Task 1", Column: c1}
	t2 := model.Task{Name: "Task 2"}
	s.db.tasks[1] = t1
	s.db.tasks[2] = t2

	ts, err := s.Tasks().GetAllByColumnID(c1.ID)

	assert.NoError(t, err)
	assert.Equal(t, 1, len(ts))
	assert.Equal(t, t1, ts[0])
}

func TestTaskRepository_Create(t *testing.T) {
	s := NewStore()
	t1 := model.Task{Name: "Task 1"}
	t2 := model.Task{Name: "Task 2"}

	t1FromRepo, err1 := s.Tasks().Create(t1)
	t2FromRepo, err2 := s.Tasks().Create(t2)

	assert.NoError(t, err1)
	assert.Equal(t, 1, t1FromRepo.ID)
	assert.Equal(t, t1.Name, t1FromRepo.Name)
	assert.NoError(t, err2)
	assert.Equal(t, 2, t2FromRepo.ID)
	assert.Equal(t, t2.Name, t2FromRepo.Name)
}

func TestTaskRepository_GetByID(t *testing.T) {
	s := NewStore()
	t1 := model.Task{ID: 1, Name: "Task 1"}
	s.db.tasks[t1.ID] = t1

	t1FromRepo, err1 := s.Tasks().GetByID(t1.ID)
	_, err2 := s.Tasks().GetByID(2)

	assert.NoError(t, err1)
	assert.Equal(t, t1.ID, t1FromRepo.ID)
	assert.Equal(t, t1.Name, t1FromRepo.Name)
	assert.Equal(t, store.ErrNotFound, err2)
}

func TestTaskRepository_Update(t *testing.T) {
	s := NewStore()
	t1 := model.Task{ID: 1, Name: "Task 1"}
	t2 := model.Task{ID: 2, Name: "Task 2"}
	s.db.tasks[t1.ID] = t1

	t1.Name = "Updated name"
	err1 := s.Tasks().Update(t1)
	err2 := s.Tasks().Update(t2)

	assert.NoError(t, err1)
	assert.Equal(t, t1.Name, s.db.tasks[t1.ID].Name)
	assert.Equal(t, store.ErrNotFound, err2)
}

func TestTaskRepository_Delete(t *testing.T) {
	s := NewStore()
	t1 := model.Task{ID: 1, Name: "Task 1"}
	t2 := model.Task{ID: 2, Name: "Task 2"}
	s.db.tasks[t1.ID] = t1

	err1 := s.Tasks().DeleteByID(t1.ID)
	err2 := s.Tasks().DeleteByID(t2.ID)

	assert.NoError(t, err1)
	assert.Equal(t, 0, len(s.db.tasks))
	assert.Equal(t, store.ErrNotFound, err2)
}
