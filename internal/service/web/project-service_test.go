package web

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/imarrche/tasker/internal/model"
	mock_store "github.com/imarrche/tasker/internal/store/mocks"
)

func TestProjectService_GetAll(t *testing.T) {
	testcases := []struct {
		name        string
		mock        func(*gomock.Controller, *mock_store.MockStore, []model.Project)
		projects    []model.Project
		expProjects []model.Project
		expError    error
	}{
		{
			name: "projects are retrieved and sorted by name alphabetically",
			mock: func(c *gomock.Controller, s *mock_store.MockStore, ps []model.Project) {
				pr := mock_store.NewMockProjectRepo(c)

				pr.EXPECT().GetAll().Return(ps, nil)
				s.EXPECT().Projects().Return(pr)
			},
			projects: []model.Project{
				{ID: 1, Name: "C"}, {ID: 2, Name: "B"}, {ID: 3, Name: "A"},
			},
			expProjects: []model.Project{
				{ID: 3, Name: "A"}, {ID: 2, Name: "B"}, {ID: 1, Name: "C"},
			},
			expError: nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(c, store, tc.projects)
			s := newProjectService(store)
			ps, err := s.GetAll()

			assert.Equal(t, tc.expError, err)
			assert.Equal(t, tc.expProjects, ps)
		})
	}
}

func TestProjectService_Create(t *testing.T) {
	testcases := []struct {
		name       string
		mock       func(*gomock.Controller, *mock_store.MockStore, model.Project)
		project    model.Project
		expProject model.Project
		expError   error
	}{
		{
			name: "project is created with default column",
			mock: func(c *gomock.Controller, s *mock_store.MockStore, p model.Project) {
				pr := mock_store.NewMockProjectRepo(c)
				cr := mock_store.NewMockColumnRepo(c)

				pr.EXPECT().Create(p).Return(p, nil)
				column := model.Column{Name: "default", Index: 1, ProjectID: p.ID}
				cr.EXPECT().Create(column).Return(
					model.Column{ID: 1, Name: column.Name, ProjectID: column.ProjectID},
					nil,
				)
				s.EXPECT().Projects().Return(pr)
				s.EXPECT().Columns().Return(cr)
			},
			project:    model.Project{Name: "Project 1"},
			expProject: model.Project{Name: "Project 1"},
			expError:   nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(c, store, tc.project)
			s := newProjectService(store)
			p, err := s.Create(tc.project)

			assert.Equal(t, tc.expError, err)
			assert.Equal(t, tc.expProject, p)
		})
	}
}

func TestProjectService_GetByID(t *testing.T) {
	testcases := []struct {
		name       string
		mock       func(*gomock.Controller, *mock_store.MockStore, model.Project)
		project    model.Project
		expProject model.Project
		expError   error
	}{
		{
			name: "project is retrieved",
			mock: func(c *gomock.Controller, s *mock_store.MockStore, p model.Project) {
				pr := mock_store.NewMockProjectRepo(c)

				pr.EXPECT().GetByID(p.ID).Return(p, nil)
				s.EXPECT().Projects().Return(pr)
			},
			project:    model.Project{ID: 1, Name: "Project 1"},
			expProject: model.Project{ID: 1, Name: "Project 1"},
			expError:   nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(c, store, tc.project)
			s := newProjectService(store)
			p, err := s.GetByID(tc.project.ID)

			assert.Equal(t, tc.expError, err)
			assert.Equal(t, tc.expProject, p)
		})
	}
}

func TestProjectService_Update(t *testing.T) {
	testcases := []struct {
		name       string
		mock       func(*gomock.Controller, *mock_store.MockStore, model.Project)
		project    model.Project
		expProject model.Project
		expError   error
	}{
		{
			name: "project is updated",
			mock: func(c *gomock.Controller, s *mock_store.MockStore, p model.Project) {
				pr := mock_store.NewMockProjectRepo(c)

				pr.EXPECT().GetByID(p.ID).Return(p, nil)
				pr.EXPECT().Update(p).Return(p, nil)
				s.EXPECT().Projects().Times(2).Return(pr)
			},
			project:    model.Project{ID: 1, Name: "Project 1"},
			expProject: model.Project{ID: 1, Name: "Project 1"},
			expError:   nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(c, store, tc.project)
			s := newProjectService(store)
			p, err := s.Update(tc.project)

			assert.Equal(t, tc.expError, err)
			assert.Equal(t, tc.expProject, p)
		})
	}
}

func TestProjectService_DeleteByID(t *testing.T) {
	testcases := []struct {
		name     string
		mock     func(*gomock.Controller, *mock_store.MockStore, model.Project)
		project  model.Project
		expError error
	}{
		{
			name: "project is deleted",
			mock: func(c *gomock.Controller, s *mock_store.MockStore, p model.Project) {
				pr := mock_store.NewMockProjectRepo(c)

				pr.EXPECT().DeleteByID(p.ID).Return(nil)
				s.EXPECT().Projects().Return(pr)
			},
			project:  model.Project{ID: 1, Name: "Project 1"},
			expError: nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(c, store, tc.project)
			s := newProjectService(store)
			err := s.DeleteByID(tc.project.ID)

			assert.Equal(t, tc.expError, err)
		})
	}
}

func TestProjectService_Validate(t *testing.T) {
	testcases := []struct {
		name     string
		project  model.Project
		expError error
	}{
		{
			name:     "project passes validation",
			project:  model.Project{Name: "Project"},
			expError: nil,
		},
		{
			name:     "project doesn't pass validation because of empty name",
			project:  model.Project{},
			expError: ErrNameIsRequired,
		},
		{
			name:     "project doesn't pass validation because of too long name",
			project:  model.Project{Name: fixedLengthString(501)},
			expError: ErrNameIsTooLong,
		},
		{
			name:     "project doesn't pass validation because of too long description",
			project:  model.Project{Name: "Project", Description: fixedLengthString(1001)},
			expError: ErrDescriptionIsTooLong,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			s := newProjectService(nil)

			assert.Equal(t, tc.expError, s.Validate(tc.project))
		})
	}
}
