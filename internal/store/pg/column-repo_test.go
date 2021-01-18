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
		mock       func([]model.Column)
		projectID  int
		expColumns []model.Column
		expError   error
	}{
		{
			name: "columns are retrieved",
			mock: func(cs []model.Column) {
				rows := sqlmock.NewRows([]string{"id", "name", "description"}).AddRow(
					1, "Project 1", "",
				)
				mock.ExpectQuery("SELECT (.+) FROM projects WHERE id = (.+);").WillReturnRows(rows)

				rows = sqlmock.NewRows([]string{"id", "name", "index", "project_id"})
				for _, c := range cs {
					rows = rows.AddRow(c.ID, c.Name, c.Index, c.ProjectID)
				}
				mock.ExpectQuery("SELECT (.+) FROM columns WHERE project_id = (.+);").WillReturnRows(rows)
			},
			projectID: 1,
			expColumns: []model.Column{
				{ID: 1, Name: "Column 1", Index: 1, ProjectID: 1},
				{ID: 2, Name: "Column 2", Index: 2, ProjectID: 1},
			},
			expError: nil,
		},
	}

	for _, tc := range testcases {
		tc.mock(tc.expColumns)

		cs, err := r.GetByProjectID(tc.projectID)

		assert.Equal(t, tc.expError, err)
		assert.Equal(t, tc.expColumns, cs)
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
			name: "column is created",
			mock: func(c model.Column) {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("INSERT INTO columns (.+) VALUES (.+);").WithArgs(
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
			name: "column is retrieved",
			mock: func(c model.Column) {
				rows := sqlmock.NewRows([]string{"id", "name", "index", "project_id"}).AddRow(
					c.ID, c.Name, c.Index, c.ProjectID,
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
			name: "column is retrieved by index and project ID",
			mock: func(c model.Column) {
				rows := sqlmock.NewRows([]string{"id", "name", "index", "project_id"}).AddRow(
					c.ID, c.Name, c.Index, c.ProjectID,
				)
				mock.ExpectQuery("SELECT (.+) FROM columns WHERE index = (.+) AND project_id = (.+);").WithArgs(
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
			name: "column is updated",
			mock: func(c model.Column) {
				mock.ExpectExec("UPDATE columns SET (.+) WHERE id = (.+);").WithArgs(
					c.Name, c.Index, c.ProjectID, c.ID,
				).WillReturnResult(sqlmock.NewResult(1, 1))
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
			name: "column is deleted",
			mock: func(c model.Column) {
				mock.ExpectExec("DELETE FROM columns WHERE id = (.+);").WithArgs(
					c.ID,
				).WillReturnResult(sqlmock.NewResult(1, 1))
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
