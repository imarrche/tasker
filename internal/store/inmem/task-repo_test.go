package inmem

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/imarrche/tasker/internal/model"
	"github.com/imarrche/tasker/internal/store"
)

func TestTaskRepo_GetByColumnID(t *testing.T) {
	s := TestStoreWithFixtures()

	ts, err := s.Tasks().GetByColumnID(1)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(ts))
}

func TestTaskRepo_Create(t *testing.T) {
	s := TestStoreWithFixtures()

	task, err := s.Tasks().Create(model.Task{Name: "Task 4", ColumnID: 2})

	assert.NoError(t, err)
	assert.Equal(t, len(s.db.tasks), task.ID)
	assert.Equal(t, "Task 4", task.Name)
}

func TestTaskRepo_GetByID(t *testing.T) {
	s := TestStoreWithFixtures()

	task, err1 := s.Tasks().GetByID(1)
	_, err2 := s.Tasks().GetByID(4)

	assert.NoError(t, err1)
	assert.Equal(t, 1, task.ID)
	assert.Equal(t, store.ErrNotFound, err2)
}

func TestTaskRepo_GetByIndexAndColumnID(t *testing.T) {
	s := TestStoreWithFixtures()

	task, err1 := s.Tasks().GetByIndexAndColumnID(1, 1)
	_, err2 := s.Tasks().GetByIndexAndColumnID(2, 2)

	assert.NoError(t, err1)
	assert.Equal(t, 1, task.ID)
	assert.Equal(t, store.ErrNotFound, err2)
}

func TestTaskRepo_Update(t *testing.T) {
	s := TestStoreWithFixtures()

	err1 := s.Tasks().Update(model.Task{ID: 1, Name: "Updated task 1"})
	err2 := s.Tasks().Update(model.Task{ID: 4})

	assert.NoError(t, err1)
	assert.Equal(t, "Updated task 1", s.db.tasks[1].Name)
	assert.Equal(t, store.ErrNotFound, err2)
}

func TestTaskRepository_DeleteByID(t *testing.T) {
	s := TestStoreWithFixtures()

	err1 := s.Tasks().DeleteByID(1)
	err2 := s.Tasks().DeleteByID(4)

	assert.NoError(t, err1)
	assert.Equal(t, 2, len(s.db.tasks))
	assert.Equal(t, 1, len(s.db.comments))
	assert.Equal(t, store.ErrNotFound, err2)
}
