package pg

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/imarrche/tasker/internal/model"
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
		mock     func([]model.Task)
		columnID int
		expTasks []model.Task
		expError error
	}{
		{
			name: "tasks are retrieved",
			mock: func(ts []model.Task) {
				rows := sqlmock.NewRows([]string{"id", "name", "index", "project_id"}).AddRow(
					1, "Column 1", 1, 1,
				)
				mock.ExpectQuery("SELECT (.+) FROM columns WHERE id = (.+);").WillReturnRows(rows)

				rows = sqlmock.NewRows([]string{"id", "name", "description", "index", "column_id"})
				for _, task := range ts {
					rows = rows.AddRow(task.ID, task.Name, task.Description, task.Index, task.ColumnID)
				}
				mock.ExpectQuery("SELECT (.+) FROM tasks WHERE column_id = (.+);").WillReturnRows(rows)
			},
			columnID: 1,
			expTasks: []model.Task{
				{ID: 1, Name: "Task 1", Description: "", Index: 1, ColumnID: 1},
				{ID: 2, Name: "Task 2", Description: "", Index: 2, ColumnID: 1},
			},
			expError: nil,
		},
	}

	for _, tc := range testcases {
		tc.mock(tc.expTasks)

		ts, err := r.GetByColumnID(tc.columnID)

		assert.Equal(t, tc.expError, err)
		assert.Equal(t, tc.expTasks, ts)
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
			name: "task is created",
			mock: func(task model.Task) {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("INSERT INTO tasks (.+) VALUES (.+);").WithArgs(
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

		task, err := r.Create(tc.task)

		assert.Equal(t, tc.expError, err)
		assert.Equal(t, tc.expTask, task)
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
			name: "task is retrieved",
			mock: func(task model.Task) {
				rows := sqlmock.NewRows([]string{"id", "name", "description", "index", "column_id"}).AddRow(
					task.ID, task.Name, task.Description, task.Index, task.ColumnID,
				)
				mock.ExpectQuery("SELECT (.+) FROM tasks WHERE id = (.+);").WithArgs(
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

		task, err := r.GetByID(tc.task.ID)

		assert.Equal(t, tc.expError, err)
		assert.Equal(t, tc.expTask, task)
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
			name: "task is retrieved by index and column ID",
			mock: func(task model.Task) {
				rows := sqlmock.NewRows([]string{"id", "name", "description", "index", "column_id"}).AddRow(
					task.ID, task.Name, task.Description, task.Index, task.ColumnID,
				)
				mock.ExpectQuery("SELECT (.+) FROM tasks WHERE index = (.+) AND column_id = (.+);").WithArgs(
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

		task, err := r.GetByIndexAndColumnID(tc.task.Index, tc.task.ColumnID)

		assert.Equal(t, tc.expError, err)
		assert.Equal(t, tc.expTask, task)
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
			name: "task is updated",
			mock: func(task model.Task) {
				mock.ExpectExec("UPDATE tasks SET (.+) WHERE id = (.+);").WithArgs(
					task.Name, task.Description, task.Index, task.ColumnID, task.ID,
				).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			task:     model.Task{ID: 1, Name: "Task 1", Index: 1, ColumnID: 1},
			expTask:  model.Task{ID: 1, Name: "Task 1", Index: 1, ColumnID: 1},
			expError: nil,
		},
	}

	for _, tc := range testcases {
		tc.mock(tc.task)

		task, err := r.Update(tc.task)

		assert.Equal(t, tc.expError, err)
		assert.Equal(t, tc.expTask, task)
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
			name: "task is deleted",
			mock: func(task model.Task) {
				mock.ExpectExec("DELETE FROM tasks WHERE id = (.+);").WithArgs(
					task.ID,
				).WillReturnResult(sqlmock.NewResult(1, 1))
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
