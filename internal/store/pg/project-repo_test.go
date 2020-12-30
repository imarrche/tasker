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
		mock        func()
		expProjects []model.Project
		expError    error
	}{
		{
			name: "OK, projects are retrieved",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "description"}).AddRow(
					1, "Project 1", "",
				)
				mock.ExpectQuery("SELECT (.+) FROM projects;").WillReturnRows(rows)
			},
			expProjects: []model.Project{
				model.Project{ID: 1, Name: "Project 1"},
			},
			expError: nil,
		},
	}

	for _, tc := range testcases {
		tc.mock()

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
		mock       func(p model.Project)
		project    model.Project
		expProject model.Project
		expError   error
	}{
		{
			name: "OK, project is created",
			mock: func(p model.Project) {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("INSERT INTO projects (.+) VALUES (.+)").WithArgs(
					p.Name, p.Description,
				).WillReturnRows(rows)
			},
			project:    model.Project{Name: "Project 1", Description: "Project 1 description."},
			expProject: model.Project{ID: 1, Name: "Project 1", Description: "Project 1 description."},
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
		mock       func(p model.Project)
		project    model.Project
		expProject model.Project
		expError   error
	}{
		{
			name: "OK, project is retrieved",
			mock: func(p model.Project) {
				rows := sqlmock.NewRows([]string{"id", "name", "description"}).AddRow(
					1, "Project 1", "",
				)
				mock.ExpectQuery("SELECT (.+) FROM projects WHERE (.+)").WithArgs(
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
		mock       func(p model.Project)
		project    model.Project
		expProject model.Project
		expError   error
	}{
		{
			name: "OK, project is updated",
			mock: func(p model.Project) {
				mock.ExpectExec("UPDATE projects SET (.+) WHERE (.+)").WithArgs(
					p.Name, p.Description, p.ID,
				).WillReturnResult(sqlmock.NewResult(0, 1))
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

func TestProjectRepo_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	r := newProjectRepo(db)

	testcases := []struct {
		name     string
		mock     func(p model.Project)
		project  model.Project
		expError error
	}{
		{
			name: "OK, project is deleted",
			mock: func(p model.Project) {
				mock.ExpectExec("DELETE FROM projects WHERE (.+)").WithArgs(
					p.ID,
				).WillReturnResult(sqlmock.NewResult(0, 1))
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
