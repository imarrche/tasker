package pg

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/imarrche/tasker/internal/model"
)

func TestProjectRepo_GetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	r := newProjectRepo(db)

	testcases := []struct {
		name        string
		mock        func([]model.Project)
		expProjects []model.Project
		expError    error
	}{
		{
			name: "projects are retrieved",
			mock: func(ps []model.Project) {
				rows := sqlmock.NewRows([]string{"id", "name", "description"})
				for _, p := range ps {
					rows = rows.AddRow(p.ID, p.Name, p.Description)
				}
				mock.ExpectQuery("SELECT (.+) FROM projects;").WillReturnRows(rows)
			},
			expProjects: []model.Project{
				model.Project{ID: 1, Name: "Project 1"}, model.Project{ID: 2, Name: "Project 2"},
			},
			expError: nil,
		},
	}

	for _, tc := range testcases {
		tc.mock(tc.expProjects)

		ps, err := r.GetAll()

		assert.Equal(t, tc.expError, err)
		assert.Equal(t, tc.expProjects, ps)
	}
}

func TestProjectRepo_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	r := newProjectRepo(db)

	testcases := []struct {
		name       string
		mock       func(model.Project)
		project    model.Project
		expProject model.Project
		expError   error
	}{
		{
			name: "project is created",
			mock: func(p model.Project) {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("INSERT INTO projects (.+) VALUES (.+);").WithArgs(
					p.Name, p.Description,
				).WillReturnRows(rows)
			},
			project:    model.Project{Name: "Project 1", Description: "Description."},
			expProject: model.Project{ID: 1, Name: "Project 1", Description: "Description."},
			expError:   nil,
		},
	}

	for _, tc := range testcases {
		tc.mock(tc.project)

		p, err := r.Create(tc.project)

		assert.Equal(t, tc.expError, err)
		assert.Equal(t, tc.expProject, p)
	}
}

func TestProjectRepo_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	r := newProjectRepo(db)

	testcases := []struct {
		name       string
		mock       func(model.Project)
		project    model.Project
		expProject model.Project
		expError   error
	}{
		{
			name: "project is retrieved",
			mock: func(p model.Project) {
				rows := sqlmock.NewRows([]string{"id", "name", "description"}).AddRow(
					p.ID, p.Name, p.Description,
				)
				mock.ExpectQuery("SELECT (.+) FROM projects WHERE id = (.+);").WithArgs(
					p.ID,
				).WillReturnRows(rows)
			},
			project:    model.Project{ID: 1, Name: "Project 1"},
			expProject: model.Project{ID: 1, Name: "Project 1"},
			expError:   nil,
		},
	}

	for _, tc := range testcases {
		tc.mock(tc.project)

		p, err := r.GetByID(tc.project.ID)

		assert.Equal(t, tc.expError, err)
		assert.Equal(t, tc.expProject, p)
	}
}

func TestProjectRepo_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	r := newProjectRepo(db)

	testcases := []struct {
		name       string
		mock       func(model.Project)
		project    model.Project
		expProject model.Project
		expError   error
	}{
		{
			name: "project is updated",
			mock: func(p model.Project) {
				mock.ExpectExec("UPDATE projects SET (.+) WHERE id = (.+);").WithArgs(
					p.Name, p.Description, p.ID,
				).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			project:    model.Project{ID: 1, Name: "Project 1"},
			expProject: model.Project{ID: 1, Name: "Project 1"},
			expError:   nil,
		},
	}

	for _, tc := range testcases {
		tc.mock(tc.project)

		p, err := r.Update(tc.project)

		assert.Equal(t, tc.expError, err)
		assert.Equal(t, tc.expProject, p)
	}
}

func TestProjectRepo_DeleteByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	r := newProjectRepo(db)

	testcases := []struct {
		name     string
		mock     func(model.Project)
		project  model.Project
		expError error
	}{
		{
			name: "project is deleted",
			mock: func(p model.Project) {
				mock.ExpectExec("DELETE FROM projects WHERE id = (.+);").WithArgs(
					p.ID,
				).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			project:  model.Project{ID: 1, Name: "Project 1"},
			expError: nil,
		},
	}

	for _, tc := range testcases {
		tc.mock(tc.project)

		err := r.DeleteByID(tc.project.ID)

		assert.Equal(t, tc.expError, err)
	}
}
