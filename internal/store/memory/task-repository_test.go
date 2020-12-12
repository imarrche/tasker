package memory

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/imarrche/tasker/internal/model"
	"github.com/imarrche/tasker/internal/store"
)

func TestTaskRepository_GetAll(t *testing.T) {
	r := NewTaskRepository()
	t1 := model.Task{Name: "Task 1"}
	t2 := model.Task{Name: "Task 2"}
	r.store.tasks[1] = t1
	r.store.tasks[2] = t2

	ts, err := r.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, 2, len(ts))
}

func TestTaskRepository_Create(t *testing.T) {
	r := NewTaskRepository()
	t1 := model.Task{Name: "Task 1"}
	t2 := model.Task{Name: "Task 2"}

	t1FromRepo, err1 := r.Create(t1)
	t2FromRepo, err2 := r.Create(t2)

	assert.NoError(t, err1)
	assert.Equal(t, t1.Name, t1FromRepo.Name)
	assert.NoError(t, err2)
	assert.Equal(t, t2.Name, t2FromRepo.Name)
}

func TestTaskRepository_GetByID(t *testing.T) {
	r := NewTaskRepository()
	t1 := model.Task{ID: 1, Name: "Task 1"}
	r.store.tasks[t1.ID] = t1

	t1FromRepo, err1 := r.GetByID(t1.ID)
	_, err2 := r.GetByID(2)

	assert.NoError(t, err1)
	assert.Equal(t, t1.ID, t1FromRepo.ID)
	assert.Equal(t, t1.Name, t1FromRepo.Name)
	assert.Equal(t, store.ErrNotFound, err2)
}

func TestTaskRepository_Update(t *testing.T) {
	r := NewTaskRepository()
	t1 := model.Task{ID: 1, Name: "Task 1"}
	t2 := model.Task{ID: 2, Name: "Task 2"}
	r.store.tasks[t1.ID] = t1

	t1.Name = "Updated name"
	err1 := r.Update(t1)
	err2 := r.Update(t2)

	assert.NoError(t, err1)
	assert.Equal(t, t1.Name, r.store.tasks[t1.ID].Name)
	assert.Equal(t, store.ErrNotFound, err2)
}

func TestTaskRepository_Delete(t *testing.T) {
	r := NewTaskRepository()
	t1 := model.Task{ID: 1, Name: "Task 1"}
	t2 := model.Task{ID: 2, Name: "Task 2"}
	r.store.tasks[t1.ID] = t1

	err1 := r.Delete(t1)
	err2 := r.Delete(t2)

	assert.NoError(t, err1)
	assert.Equal(t, 0, len(r.store.tasks))
	assert.Equal(t, store.ErrNotFound, err2)
}
