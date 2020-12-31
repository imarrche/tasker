package pg

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/imarrche/tasker/internal/model"
)

func TestColumnRepo_GetByProjectID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	r := newColumnRepo(db)

	testcases := []struct {
		name       string
		mock       func()
		expColumns []model.Column
		expError   error
	}{
		{
			name: "OK, columns are retrieved",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "index", "project_id"}).AddRow(
					1, "Column 1", 1, 1,
				)
				mock.ExpectQuery("SELECT (.+) FROM columns WHERE project_id = (.+);").WillReturnRows(rows)
			},
			expColumns: []model.Column{
				model.Column{ID: 1, Name: "Column 1", Index: 1, ProjectID: 1},
			},
			expError: nil,
		},
	}

	for _, tc := range testcases {
		tc.mock()

		cs, err := r.GetByProjectID(1)

		assert.Equal(t, tc.expError, err)
		for i := range cs {
			assert.Equal(t, tc.expColumns[i], cs[i])
		}
	}
}

func TestColumnRepo_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	r := newColumnRepo(db)

	testcases := []struct {
		name      string
		mock      func(model.Column)
		column    model.Column
		expColumn model.Column
		expError  error
	}{
		{
			name: "OK, column is created",
			mock: func(c model.Column) {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("INSERT INTO columns (.+) VALUES (.+)").WithArgs(
					c.Name, c.Index, c.ProjectID,
				).WillReturnRows(rows)
			},
			column:    model.Column{Name: "Column 1", Index: 1, ProjectID: 1},
			expColumn: model.Column{ID: 1, Name: "Column 1", Index: 1, ProjectID: 1},
			expError:  nil,
		},
	}

	for _, tc := range testcases {
		tc.mock(tc.column)

		c, err := r.Create(tc.column)

		assert.Equal(t, tc.expError, err)
		assert.Equal(t, tc.expColumn, c)
	}
}

func TestColumnRepo_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	r := newColumnRepo(db)

	testcases := []struct {
		name      string
		mock      func(model.Column)
		column    model.Column
		expColumn model.Column
		expError  error
	}{
		{
			name: "OK, column is retrieved",
			mock: func(c model.Column) {
				rows := sqlmock.NewRows([]string{"id", "name", "index", "project_id"}).AddRow(
					1, "Column 1", 1, 1,
				)
				mock.ExpectQuery("SELECT (.+) FROM columns WHERE (.+);").WithArgs(
					c.ID,
				).WillReturnRows(rows)
			},
			column:    model.Column{ID: 1, Name: "Column 1", Index: 1, ProjectID: 1},
			expColumn: model.Column{ID: 1, Name: "Column 1", Index: 1, ProjectID: 1},
			expError:  nil,
		},
	}

	for _, tc := range testcases {
		tc.mock(tc.column)

		c, err := r.GetByID(tc.column.ID)

		assert.Equal(t, tc.expError, err)
		assert.Equal(t, tc.expColumn, c)
	}
}

func TestColumnRepo_GetByIndexAndProjectID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	r := newColumnRepo(db)

	testcases := []struct {
		name      string
		mock      func(model.Column)
		column    model.Column
		expColumn model.Column
		expError  error
	}{
		{
			name: "OK, column is retrieved",
			mock: func(c model.Column) {
				rows := sqlmock.NewRows([]string{"id", "name", "index", "project_id"}).AddRow(
					1, "Column 1", 1, 1,
				)
				mock.ExpectQuery("SELECT (.+) FROM columns WHERE (.+);").WithArgs(
					c.Index, c.ProjectID,
				).WillReturnRows(rows)
			},
			column:    model.Column{ID: 1, Name: "Column 1", Index: 1, ProjectID: 1},
			expColumn: model.Column{ID: 1, Name: "Column 1", Index: 1, ProjectID: 1},
			expError:  nil,
		},
	}

	for _, tc := range testcases {
		tc.mock(tc.column)

		c, err := r.GetByIndexAndProjectID(tc.column.Index, tc.column.ProjectID)

		assert.Equal(t, tc.expError, err)
		assert.Equal(t, tc.expColumn, c)
	}
}

func TestColumnRepo_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	r := newColumnRepo(db)

	testcases := []struct {
		name      string
		mock      func(model.Column)
		column    model.Column
		expColumn model.Column
		expError  error
	}{
		{
			name: "OK, column is updated",
			mock: func(c model.Column) {
				mock.ExpectExec("UPDATE columns SET (.+) WHERE (.+)").WithArgs(
					c.Name, c.Index, c.ID,
				).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			column:    model.Column{ID: 1, Name: "Column 1", Index: 1, ProjectID: 1},
			expColumn: model.Column{ID: 1, Name: "Column 1", Index: 1, ProjectID: 1},
			expError:  nil,
		},
	}

	for _, tc := range testcases {
		tc.mock(tc.column)

		c, err := r.Update(tc.column)

		assert.Equal(t, tc.expError, err)
		assert.Equal(t, tc.expColumn, c)
	}
}

func TestColumnRepo_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	r := newColumnRepo(db)

	testcases := []struct {
		name     string
		mock     func(model.Column)
		column   model.Column
		expError error
	}{
		{
			name: "OK, column is deleted",
			mock: func(c model.Column) {
				mock.ExpectExec("DELETE FROM columns WHERE (.+)").WithArgs(
					c.ID,
				).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			column:   model.Column{ID: 1, Name: "Column 1", Index: 1, ProjectID: 1},
			expError: nil,
		},
	}

	for _, tc := range testcases {
		tc.mock(tc.column)

		err := r.DeleteByID(tc.column.ID)

		assert.Equal(t, tc.expError, err)
	}
}
