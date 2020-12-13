package web

import (
	"errors"
	"math/rand"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/imarrche/tasker/internal/model"
	mock_store "github.com/imarrche/tasker/internal/store/mocks"
)

func TestProjectService_GetAll(t *testing.T) {
	testcases := [...]struct {
		name             string
		mock             func(*mock_store.MockProjectRepository)
		expectedProjects []model.Project
		expectedError    error
	}{
		{
			name: "Projects are retrieved and sorted by name alphabetically.",
			mock: func(pr *mock_store.MockProjectRepository) {
				pr.EXPECT().GetAll().Return(
					[]model.Project{
						model.Project{Name: "C"},
						model.Project{Name: "B"},
						model.Project{Name: "A"},
					},
					nil,
				)
			},
			expectedProjects: []model.Project{
				model.Project{Name: "A"},
				model.Project{Name: "B"},
				model.Project{Name: "C"},
			},
			expectedError: nil,
		},
		{
			name: "Error occured while retrieving projects.",
			mock: func(s *mock_store.MockProjectRepository) {
				s.EXPECT().GetAll().Return(
					nil,
					errors.New("couldn't get projects"),
				)
			},
			expectedProjects: nil,
			expectedError:    errors.New("couldn't get projects"),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			pr := mock_store.NewMockProjectRepository(c)
			tc.mock(pr)
			s := NewProjectService(pr, nil)

			ps, err := s.GetAll()
			assert.Equal(t, tc.expectedProjects, ps)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestProjectService_Create(t *testing.T) {
	testcases := [...]struct {
		name string
		mock func(
			*mock_store.MockProjectRepository,
			*mock_store.MockColumnRepository,
			model.Project,
		)
		project         model.Project
		expectedProject model.Project
		expectedError   error
	}{
		{
			name: "Project is created with default column.",
			mock: func(
				pr *mock_store.MockProjectRepository,
				cr *mock_store.MockColumnRepository,
				p model.Project,
			) {
				defaultColumn := model.Column{Name: "default", Project: p}
				pr.EXPECT().Create(p).Return(p, nil)
				cr.EXPECT().Create(defaultColumn).Return(defaultColumn, nil)
			},
			project:         model.Project{Name: "P"},
			expectedProject: model.Project{Name: "P"},
			expectedError:   nil,
		},
		{
			name: "Project didn't pass validation.",
			mock: func(
				pr *mock_store.MockProjectRepository,
				cr *mock_store.MockColumnRepository,
				p model.Project,
			) {
			},
			project:         model.Project{Name: ""},
			expectedProject: model.Project{},
			expectedError:   ErrNameIsRequired,
		},
		{
			name: "Error occured while creating a project.",
			mock: func(
				pr *mock_store.MockProjectRepository,
				cr *mock_store.MockColumnRepository,
				p model.Project,
			) {
				pr.EXPECT().Create(p).Return(model.Project{}, errors.New("couldn't create project"))
			},
			project:         model.Project{Name: "P"},
			expectedProject: model.Project{},
			expectedError:   errors.New("couldn't create project"),
		},
		{
			name: "Project is created, but error occured while creating default column.",
			mock: func(
				pr *mock_store.MockProjectRepository,
				cr *mock_store.MockColumnRepository,
				p model.Project,
			) {
				defaultColumn := model.Column{Name: "default", Project: p}
				pr.EXPECT().Create(p).Return(p, nil)
				cr.EXPECT().Create(defaultColumn).Return(model.Column{}, errors.New("couldn't create column"))
			},
			project:         model.Project{Name: "P"},
			expectedProject: model.Project{},
			expectedError:   errors.New("couldn't create column"),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			pr := mock_store.NewMockProjectRepository(c)
			cr := mock_store.NewMockColumnRepository(c)
			tc.mock(pr, cr, tc.project)
			s := NewProjectService(pr, cr)

			p, err := s.Create(tc.project)
			assert.Equal(t, tc.expectedProject, p)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestProjectService_GetByID(t *testing.T) {
	testcases := [...]struct {
		name            string
		mock            func(*mock_store.MockProjectRepository, model.Project)
		project         model.Project
		expectedProject model.Project
		expectedError   error
	}{
		{
			name: "Project is retrieved by ID.",
			mock: func(pr *mock_store.MockProjectRepository, p model.Project) {
				pr.EXPECT().GetByID(p.ID).Return(p, nil)
			},
			project:         model.Project{ID: 1, Name: "P1"},
			expectedProject: model.Project{ID: 1, Name: "P1"},
			expectedError:   nil,
		},
		{
			name: "Error occured while retrieving project.",
			mock: func(pr *mock_store.MockProjectRepository, p model.Project) {
				pr.EXPECT().GetByID(p.ID).Return(model.Project{}, errors.New("couldn't get project"))
			},
			project:         model.Project{ID: 1, Name: "P1"},
			expectedProject: model.Project{},
			expectedError:   errors.New("couldn't get project"),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			pr := mock_store.NewMockProjectRepository(c)
			tc.mock(pr, tc.project)
			s := NewProjectService(pr, nil)

			p, err := s.GetByID(tc.project.ID)
			assert.Equal(t, tc.expectedProject, p)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestProjectService_Update(t *testing.T) {
	testcases := [...]struct {
		name          string
		mock          func(*mock_store.MockProjectRepository, model.Project)
		project       model.Project
		expectedError error
	}{
		{
			name: "Project is updated.",
			mock: func(pr *mock_store.MockProjectRepository, p model.Project) {
				pr.EXPECT().Update(p).Return(nil)
			},
			project:       model.Project{ID: 1, Name: "P1"},
			expectedError: nil,
		},
		{
			name: "Error occured while updating project.",
			mock: func(pr *mock_store.MockProjectRepository, p model.Project) {
				pr.EXPECT().Update(p).Return(errors.New("couldn't update project"))
			},
			project:       model.Project{ID: 1, Name: "P1"},
			expectedError: errors.New("couldn't update project"),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			pr := mock_store.NewMockProjectRepository(c)
			tc.mock(pr, tc.project)
			s := NewProjectService(pr, nil)

			err := s.Update(tc.project)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestProjectService_Delete(t *testing.T) {
	testcases := [...]struct {
		name          string
		mock          func(*mock_store.MockProjectRepository, model.Project)
		project       model.Project
		expectedError error
	}{
		{
			name: "Project is deleted.",
			mock: func(pr *mock_store.MockProjectRepository, p model.Project) {
				pr.EXPECT().Delete(p).Return(nil)
			},
			project:       model.Project{ID: 1, Name: "P1"},
			expectedError: nil,
		},
		{
			name: "Error occured while deleting project.",
			mock: func(pr *mock_store.MockProjectRepository, p model.Project) {
				pr.EXPECT().Delete(p).Return(errors.New("couldn't delete project"))
			},
			project:       model.Project{ID: 1, Name: "P1"},
			expectedError: errors.New("couldn't delete project"),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			pr := mock_store.NewMockProjectRepository(c)
			tc.mock(pr, tc.project)
			s := NewProjectService(pr, nil)

			err := s.Delete(tc.project)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestProjectService_Validate(t *testing.T) {
	fixedLengthString := func(length int) string {
		var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

		rs := make([]rune, length)
		for i := range rs {
			rs[i] = letters[rand.Intn(len(letters))]
		}

		return string(rs)
	}

	testcases := [...]struct {
		name          string
		project       model.Project
		expectedError error
	}{
		{
			name:          "Project passed validation.",
			project:       model.Project{Name: "P"},
			expectedError: nil,
		},
		{
			name:          "Project's name was not provided.",
			project:       model.Project{},
			expectedError: ErrNameIsRequired,
		},
		{
			name:          "Project's name was too long.",
			project:       model.Project{Name: fixedLengthString(501)},
			expectedError: ErrNameIsTooLong,
		},
		{
			name:          "Project's description was too long.",
			project:       model.Project{Name: "P", Description: fixedLengthString(1001)},
			expectedError: ErrDescriptionIsTooLong,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			s := NewProjectService(nil, nil)

			assert.Equal(t, tc.expectedError, s.Validate(tc.project))
		})
	}
}
