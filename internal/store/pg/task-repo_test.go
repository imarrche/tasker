package pg

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/imarrche/tasker/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestTaskRepo_GetByColumnID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	r := newTaskRepo(db)

	testcases := []struct {
		name     string
		mock     func()
		expTasks []model.Task
		expError error
	}{
		{
			name: "OK, tasks are retrieved",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "description", "index", "column_id"}).AddRow(
					1, "Task 1", "", 1, 1,
				)
				mock.ExpectQuery("SELECT (.+) FROM tasks WHERE column_id = (.+);").WillReturnRows(rows)
			},
			expTasks: []model.Task{
				model.Task{ID: 1, Name: "Task 1", Description: "", Index: 1, ColumnID: 1},
			},
			expError: nil,
		},
	}

	for _, tc := range testcases {
		tc.mock()

		ts, err := r.GetByColumnID(1)

		assert.Equal(t, tc.expError, err)
		for i := range ts {
			assert.Equal(t, tc.expTasks[i], ts[i])
		}
	}
}

func TestTaskRepo_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	r := newTaskRepo(db)

	testcases := []struct {
		name     string
		mock     func(model.Task)
		task     model.Task
		expTask  model.Task
		expError error
	}{
		{
			name: "OK, task is created",
			mock: func(task model.Task) {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("INSERT INTO tasks (.+) VALUES (.+)").WithArgs(
					task.Name, task.Description, task.Index, task.ColumnID,
				).WillReturnRows(rows)
			},
			task:     model.Task{Name: "Task 1", Index: 1, ColumnID: 1},
			expTask:  model.Task{ID: 1, Name: "Task 1", Index: 1, ColumnID: 1},
			expError: nil,
		},
	}

	for _, tc := range testcases {
		tc.mock(tc.task)

		c, err := r.Create(tc.task)

		assert.Equal(t, tc.expError, err)
		assert.Equal(t, tc.expTask, c)
	}
}

func TestTaskRepo_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	r := newTaskRepo(db)

	testcases := []struct {
		name     string
		mock     func(model.Task)
		task     model.Task
		expTask  model.Task
		expError error
	}{
		{
			name: "OK, task is retrieved",
			mock: func(task model.Task) {
				rows := sqlmock.NewRows([]string{"id", "name", "description", "index", "column_id"}).AddRow(
					1, "Task 1", "", 1, 1,
				)
				mock.ExpectQuery("SELECT FROM tasks WHERE (.+);").WithArgs(
					task.ID,
				).WillReturnRows(rows)
			},
			task:     model.Task{ID: 1, Name: "Task 1", Index: 1, ColumnID: 1},
			expTask:  model.Task{ID: 1, Name: "Task 1", Index: 1, ColumnID: 1},
			expError: nil,
		},
	}

	for _, tc := range testcases {
		tc.mock(tc.task)

		c, err := r.GetByID(tc.task.ID)

		assert.Equal(t, tc.expError, err)
		assert.Equal(t, tc.expTask, c)
	}
}

func TestTaskRepo_GetByIndexAndColumnID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	r := newTaskRepo(db)

	testcases := []struct {
		name     string
		mock     func(model.Task)
		task     model.Task
		expTask  model.Task
		expError error
	}{
		{
			name: "OK, task is retrieved",
			mock: func(task model.Task) {
				rows := sqlmock.NewRows([]string{"id", "name", "description", "index", "column_id"}).AddRow(
					1, "Task 1", "", 1, 1,
				)
				mock.ExpectQuery("SELECT FROM tasks WHERE (.+);").WithArgs(
					task.Index, task.ColumnID,
				).WillReturnRows(rows)
			},
			task:     model.Task{ID: 1, Name: "Task 1", Index: 1, ColumnID: 1},
			expTask:  model.Task{ID: 1, Name: "Task 1", Index: 1, ColumnID: 1},
			expError: nil,
		},
	}

	for _, tc := range testcases {
		tc.mock(tc.task)

		c, err := r.GetByIndexAndColumnID(tc.task.Index, tc.task.ColumnID)

		assert.Equal(t, tc.expError, err)
		assert.Equal(t, tc.expTask, c)
	}
}

func TestTaskRepo_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	r := newTaskRepo(db)

	testcases := []struct {
		name     string
		mock     func(model.Task)
		task     model.Task
		expTask  model.Task
		expError error
	}{
		{
			name: "OK, task is updated",
			mock: func(task model.Task) {
				mock.ExpectExec("UPDATE tasks SET (.+) WHERE (.+)").WithArgs(
					task.Name, task.Description, task.Index, task.ID,
				).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			task:     model.Task{ID: 1, Name: "Task 1", Index: 1, ColumnID: 1},
			expTask:  model.Task{ID: 1, Name: "Task 1", Index: 1, ColumnID: 1},
			expError: nil,
		},
	}

	for _, tc := range testcases {
		tc.mock(tc.task)

		c, err := r.Update(tc.task)

		assert.Equal(t, tc.expError, err)
		assert.Equal(t, tc.expTask, c)
	}
}

func TestTaskRepo_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	r := newTaskRepo(db)

	testcases := []struct {
		name     string
		mock     func(model.Task)
		task     model.Task
		expError error
	}{
		{
			name: "OK, task is deleted",
			mock: func(task model.Task) {
				mock.ExpectExec("DELETE FROM tasks WHERE (.+)").WithArgs(
					task.ID,
				).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			task:     model.Task{ID: 1, Name: "Task 1", Index: 1, ColumnID: 1},
			expError: nil,
		},
	}

	for _, tc := range testcases {
		tc.mock(tc.task)

		err := r.DeleteByID(tc.task.ID)

		assert.Equal(t, tc.expError, err)
	}
}
