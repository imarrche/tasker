package pg

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/imarrche/tasker/internal/model"
)

func TestCommentRepo_GetByTaskID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	r := newCommentRepo(db)

	testcases := []struct {
		name        string
		mock        func([]model.Comment)
		taskID      int
		expComments []model.Comment
		expError    error
	}{
		{
			name: "comments are retrieved",
			mock: func(cs []model.Comment) {
				rows := sqlmock.NewRows([]string{"id", "name", "description", "index", "column_id"}).AddRow(
					1, "Task 1", "", 1, 1,
				)
				mock.ExpectQuery("SELECT (.+) FROM tasks WHERE id = (.+);").WillReturnRows(rows)

				rows = sqlmock.NewRows([]string{"id", "text", "created_at", "task_id"})
				for _, c := range cs {
					rows = rows.AddRow(c.ID, c.Text, c.CreatedAt, c.TaskID)
				}
				mock.ExpectQuery("SELECT (.+) FROM comments WHERE task_id = (.+);").WillReturnRows(rows)
			},
			taskID: 1,
			expComments: []model.Comment{
				{ID: 1, Text: "Comment.", CreatedAt: time.Time{}, TaskID: 1},
				{ID: 2, Text: "Comment.", CreatedAt: time.Time{}, TaskID: 1},
			},
			expError: nil,
		},
	}

	for _, tc := range testcases {
		tc.mock(tc.expComments)

		cs, err := r.GetByTaskID(tc.taskID)

		assert.Equal(t, tc.expError, err)
		assert.Equal(t, tc.expComments, cs)
	}
}

func TestCommentRepo_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	r := newCommentRepo(db)

	testcases := []struct {
		name       string
		mock       func(model.Comment)
		comment    model.Comment
		expComment model.Comment
		expError   error
	}{
		{
			name: "comment is created",
			mock: func(c model.Comment) {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("INSERT INTO comments (.+) VALUES (.+);").WithArgs(
					c.Text, c.CreatedAt, c.TaskID,
				).WillReturnRows(rows)
			},
			comment:    model.Comment{Text: "Comment.", CreatedAt: time.Time{}, TaskID: 1},
			expComment: model.Comment{ID: 1, Text: "Comment.", CreatedAt: time.Time{}, TaskID: 1},
			expError:   nil,
		},
	}

	for _, tc := range testcases {
		tc.mock(tc.comment)

		c, err := r.Create(tc.comment)

		assert.Equal(t, tc.expError, err)
		assert.Equal(t, tc.expComment, c)
	}
}

func TestCommentRepo_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	r := newCommentRepo(db)

	testcases := []struct {
		name       string
		mock       func(model.Comment)
		comment    model.Comment
		expComment model.Comment
		expError   error
	}{
		{
			name: "comment is retrieved",
			mock: func(c model.Comment) {
				rows := sqlmock.NewRows([]string{"id", "text", "created_at", "task_id"}).AddRow(
					c.ID, c.Text, c.CreatedAt, c.TaskID,
				)
				mock.ExpectQuery("SELECT (.+) FROM comments WHERE id = (.+);").WithArgs(
					c.ID,
				).WillReturnRows(rows)
			},
			comment:    model.Comment{ID: 1, Text: "Comment.", CreatedAt: time.Time{}, TaskID: 1},
			expComment: model.Comment{ID: 1, Text: "Comment.", CreatedAt: time.Time{}, TaskID: 1},
			expError:   nil,
		},
	}

	for _, tc := range testcases {
		tc.mock(tc.comment)

		c, err := r.GetByID(tc.comment.ID)

		assert.Equal(t, tc.expError, err)
		assert.Equal(t, tc.expComment, c)
	}
}

func TestCommentRepo_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	r := newCommentRepo(db)

	testcases := []struct {
		name       string
		mock       func(model.Comment)
		comment    model.Comment
		expComment model.Comment
		expError   error
	}{
		{
			name: "comment is updated",
			mock: func(c model.Comment) {
				mock.ExpectExec("UPDATE comments SET (.+) WHERE id = (.+);").WithArgs(
					c.Text, c.CreatedAt, c.TaskID, c.ID,
				).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			comment:    model.Comment{ID: 1, Text: "Comment.", CreatedAt: time.Time{}, TaskID: 1},
			expComment: model.Comment{ID: 1, Text: "Comment.", CreatedAt: time.Time{}, TaskID: 1},
			expError:   nil,
		},
	}

	for _, tc := range testcases {
		tc.mock(tc.comment)

		c, err := r.Update(tc.comment)

		assert.Equal(t, tc.expError, err)
		assert.Equal(t, tc.expComment, c)
	}
}

func TestCommentRepo_DeleteByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	r := newCommentRepo(db)

	testcases := []struct {
		name     string
		mock     func(model.Comment)
		comment  model.Comment
		expError error
	}{
		{
			name: "comment is deleted",
			mock: func(c model.Comment) {
				mock.ExpectExec("DELETE FROM comments WHERE id = (.+);").WithArgs(
					c.ID,
				).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			comment:  model.Comment{ID: 1, Text: "Comment.", CreatedAt: time.Time{}, TaskID: 1},
			expError: nil,
		},
	}

	for _, tc := range testcases {
		tc.mock(tc.comment)

		err := r.DeleteByID(tc.comment.ID)

		assert.Equal(t, tc.expError, err)
	}
}
