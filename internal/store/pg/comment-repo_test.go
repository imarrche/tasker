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
		mock        func()
		expComments []model.Comment
		expError    error
	}{
		{
			name: "OK, comments are retrieved",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "text", "created_at", "task_id"}).AddRow(
					1, "Comment.", time.Now(), 1,
				)
				mock.ExpectQuery("SELECT (.+) FROM comments WHERE task_id = (.+);").WillReturnRows(rows)
			},
			expComments: []model.Comment{
				model.Comment{ID: 1, Text: "Comment."},
			},
			expError: nil,
		},
	}

	for _, tc := range testcases {
		tc.mock()

		ps, err := r.GetByTaskID(1)

		assert.Equal(t, tc.expError, err)
		for i := range ps {
			assert.Equal(t, tc.expComments[i].ID, ps[i].ID)
		}
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
			name: "OK, comment is created",
			mock: func(c model.Comment) {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("INSERT INTO comments (.+) VALUES (.+)").WithArgs(
					c.Text, c.CreatedAt, c.TaskID,
				).WillReturnRows(rows)
			},
			comment:    model.Comment{Text: "Comment.", CreatedAt: time.Now(), TaskID: 1},
			expComment: model.Comment{ID: 1, Text: "Comment.", CreatedAt: time.Now(), TaskID: 1},
			expError:   nil,
		},
	}

	for _, tc := range testcases {
		tc.mock(tc.comment)

		p, err := r.Create(tc.comment)

		assert.Equal(t, tc.expError, err)
		assert.Equal(t, tc.expComment.ID, p.ID)
		assert.Equal(t, tc.expComment.Text, p.Text)
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
			name: "OK, comment is retrieved",
			mock: func(c model.Comment) {
				rows := sqlmock.NewRows([]string{"id", "text", "created_at", "task_id"}).AddRow(
					1, "Comment.", time.Now(), 1,
				)
				mock.ExpectQuery("SELECT FROM comments WHERE (.+);").WithArgs(
					c.ID,
				).WillReturnRows(rows)
			},
			comment:    model.Comment{ID: 1, Text: "Comment."},
			expComment: model.Comment{ID: 1, Text: "Comment."},
			expError:   nil,
		},
	}

	for _, tc := range testcases {
		tc.mock(tc.comment)

		p, err := r.GetByID(tc.comment.ID)

		assert.Equal(t, tc.expError, err)
		assert.Equal(t, tc.expComment.ID, p.ID)
		assert.Equal(t, tc.expComment.Text, p.Text)
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
			name: "OK, comment is updated",
			mock: func(c model.Comment) {
				mock.ExpectExec("UPDATE comments SET (.+) WHERE (.+)").WithArgs(
					c.Text, c.ID,
				).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			comment:    model.Comment{ID: 1, Text: "Comment."},
			expComment: model.Comment{ID: 1, Text: "Comment."},
			expError:   nil,
		},
	}

	for _, tc := range testcases {
		tc.mock(tc.comment)

		p, err := r.Update(tc.comment)

		assert.Equal(t, tc.expError, err)
		assert.Equal(t, tc.expComment.Text, p.Text)
	}
}

func TestCommentRepo_Delete(t *testing.T) {
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
			name: "OK, comment is deleted",
			mock: func(c model.Comment) {
				mock.ExpectExec("DELETE FROM comments WHERE (.+)").WithArgs(
					c.ID,
				).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			comment:  model.Comment{ID: 1, Text: "Comment."},
			expError: nil,
		},
	}

	for _, tc := range testcases {
		tc.mock(tc.comment)

		err := r.DeleteByID(tc.comment.ID)

		assert.Equal(t, tc.expError, err)
	}
}
