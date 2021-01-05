package inmem

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/imarrche/tasker/internal/model"
)

func TestTaskRepo_GetByColumnID(t *testing.T) {
	s := TestStoreWithFixtures()

	ts, err := s.Tasks().GetByColumnID(1)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(ts))
}

func TestTaskRepo_Create(t *testing.T) {
	s := TestStoreWithFixtures()

	task, err := s.Tasks().Create(model.Task{Name: "Task 4", Index: 2, ColumnID: 2})

	assert.NoError(t, err)
	assert.Equal(t, model.Task{ID: 4, Name: "Task 4", Index: 2, ColumnID: 2}, task)
}

func TestTaskRepo_GetByID(t *testing.T) {
	s := TestStoreWithFixtures()

	task, err := s.Tasks().GetByID(1)

	assert.NoError(t, err)
	assert.Equal(t, 1, task.ID)
}

func TestTaskRepo_GetByIndexAndColumnID(t *testing.T) {
	s := TestStoreWithFixtures()

	task, err := s.Tasks().GetByIndexAndColumnID(1, 1)

	assert.NoError(t, err)
	assert.Equal(t, 1, task.ID)
}

func TestTaskRepo_Update(t *testing.T) {
	s := TestStoreWithFixtures()
	task := model.Task{ID: 1, Name: "Updated task 1", Index: 1, ColumnID: 1}

	task1, err := s.Tasks().Update(task)

	assert.NoError(t, err)
	assert.Equal(t, task, task1)
}

func TestTaskRepository_DeleteByID(t *testing.T) {
	s := TestStoreWithFixtures()

	err := s.Tasks().DeleteByID(1)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(s.db.tasks))
	assert.Equal(t, 1, len(s.db.comments))
}
