package web

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/imarrche/tasker/internal/model"
	"github.com/imarrche/tasker/internal/store"
	mock_store "github.com/imarrche/tasker/internal/store/mocks"
)

func TestProjectService_GetAll(t *testing.T) {
	testcases := []struct {
		name        string
		mock        func(*mock_store.MockStore, *gomock.Controller)
		expProjects []model.Project
		expError    error
	}{
		{
			name: "Projects are retrieved and sorted by name alphabetically",
			mock: func(s *mock_store.MockStore, c *gomock.Controller) {
				pr := mock_store.NewMockProjectRepo(c)
				pr.EXPECT().GetAll().Return(
					[]model.Project{
						model.Project{ID: 1, Name: "C"},
						model.Project{ID: 2, Name: "B"},
						model.Project{ID: 3, Name: "A"},
					},
					nil,
				)
				s.EXPECT().Projects().Return(pr)
			},
			expProjects: []model.Project{
				model.Project{ID: 3, Name: "A"},
				model.Project{ID: 2, Name: "B"},
				model.Project{ID: 1, Name: "C"},
			},
			expError: nil,
		},
		{
			name: "Error occures while retrieving projects",
			mock: func(s *mock_store.MockStore, c *gomock.Controller) {
				pr := mock_store.NewMockProjectRepo(c)

				pr.EXPECT().GetAll().Return(nil, store.ErrDbQuery)
				s.EXPECT().Projects().Return(pr)
			},
			expProjects: nil,
			expError:    store.ErrDbQuery,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(store, c)
			s := newProjectService(store)

			ps, err := s.GetAll()
			assert.Equal(t, tc.expProjects, ps)
			assert.Equal(t, tc.expError, err)
		})
	}
}

func TestProjectService_Create(t *testing.T) {
	testcases := []struct {
		name       string
		mock       func(*mock_store.MockStore, *gomock.Controller, model.Project)
		project    model.Project
		expProject model.Project
		expError   error
	}{
		{
			name: "Project is created with default column",
			mock: func(s *mock_store.MockStore, c *gomock.Controller, p model.Project) {
				pr := mock_store.NewMockProjectRepo(c)
				cr := mock_store.NewMockColumnRepo(c)

				pr.EXPECT().Create(p).Return(p, nil)
				defaultColumn := model.Column{Name: "default", Index: 1, ProjectID: p.ID}
				cr.EXPECT().Create(defaultColumn).Return(
					model.Column{
						ID:        1,
						Name:      defaultColumn.Name,
						ProjectID: defaultColumn.ProjectID,
					},
					nil,
				)
				s.EXPECT().Projects().Return(pr)
				s.EXPECT().Columns().Return(cr)
			},
			project:    model.Project{Name: "P"},
			expProject: model.Project{Name: "P"},
			expError:   nil,
		},
		{
			name:       "Project doesn't pass validation",
			mock:       func(s *mock_store.MockStore, c *gomock.Controller, p model.Project) {},
			project:    model.Project{Name: ""},
			expProject: model.Project{},
			expError:   ErrNameIsRequired,
		},
		{
			name: "Error occures while creating a project",
			mock: func(s *mock_store.MockStore, c *gomock.Controller, p model.Project) {
				pr := mock_store.NewMockProjectRepo(c)

				pr.EXPECT().Create(p).Return(model.Project{}, store.ErrDbQuery)
				s.EXPECT().Projects().Return(pr)
			},
			project:    model.Project{Name: "P"},
			expProject: model.Project{},
			expError:   store.ErrDbQuery,
		},
		{
			name: "Error occures while creating default column",
			mock: func(s *mock_store.MockStore, c *gomock.Controller, p model.Project) {
				pr := mock_store.NewMockProjectRepo(c)
				cr := mock_store.NewMockColumnRepo(c)

				pr.EXPECT().Create(p).Return(p, nil)
				defaultColumn := model.Column{Name: "default", Index: 1, ProjectID: p.ID}
				cr.EXPECT().Create(defaultColumn).Return(model.Column{}, store.ErrDbQuery)
				s.EXPECT().Projects().Return(pr)
				s.EXPECT().Columns().Return(cr)
			},
			project:    model.Project{Name: "P"},
			expProject: model.Project{},
			expError:   store.ErrDbQuery,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(store, c, tc.project)
			s := newProjectService(store)

			p, err := s.Create(tc.project)
			assert.Equal(t, tc.expProject, p)
			assert.Equal(t, tc.expError, err)
		})
	}
}

func TestProjectService_GetByID(t *testing.T) {
	testcases := []struct {
		name       string
		mock       func(*mock_store.MockStore, *gomock.Controller, model.Project)
		project    model.Project
		expProject model.Project
		expError   error
	}{
		{
			name: "Project is retrieved",
			mock: func(s *mock_store.MockStore, c *gomock.Controller, p model.Project) {
				pr := mock_store.NewMockProjectRepo(c)

				pr.EXPECT().GetByID(p.ID).Return(p, nil)
				s.EXPECT().Projects().Return(pr)
			},
			project:    model.Project{ID: 1, Name: "P"},
			expProject: model.Project{ID: 1, Name: "P"},
			expError:   nil,
		},
		{
			name: "Error occures while retrieving project",
			mock: func(s *mock_store.MockStore, c *gomock.Controller, p model.Project) {
				pr := mock_store.NewMockProjectRepo(c)

				pr.EXPECT().GetByID(p.ID).Return(model.Project{}, store.ErrDbQuery)
				s.EXPECT().Projects().Return(pr)
			},
			project:    model.Project{ID: 1, Name: "P"},
			expProject: model.Project{},
			expError:   store.ErrDbQuery,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(store, c, tc.project)
			s := newProjectService(store)

			p, err := s.GetByID(tc.project.ID)
			assert.Equal(t, tc.expProject, p)
			assert.Equal(t, tc.expError, err)
		})
	}
}

func TestProjectService_Update(t *testing.T) {
	testcases := []struct {
		name       string
		mock       func(*mock_store.MockStore, *gomock.Controller, model.Project)
		project    model.Project
		expProject model.Project
		expError   error
	}{
		{
			name: "Project is updated",
			mock: func(s *mock_store.MockStore, c *gomock.Controller, p model.Project) {
				pr := mock_store.NewMockProjectRepo(c)

				pr.EXPECT().Update(p).Return(p, nil)
				s.EXPECT().Projects().Return(pr)
			},
			project:    model.Project{ID: 1, Name: "P"},
			expProject: model.Project{ID: 1, Name: "P"},
			expError:   nil,
		},
		{
			name:       "Project doesn't pass validation",
			mock:       func(s *mock_store.MockStore, c *gomock.Controller, p model.Project) {},
			project:    model.Project{Name: ""},
			expProject: model.Project{},
			expError:   ErrNameIsRequired,
		},
		{
			name: "Error occured while updating project.",
			mock: func(s *mock_store.MockStore, c *gomock.Controller, p model.Project) {
				pr := mock_store.NewMockProjectRepo(c)

				pr.EXPECT().Update(p).Return(model.Project{}, store.ErrDbQuery)
				s.EXPECT().Projects().Return(pr)
			},
			project:    model.Project{ID: 1, Name: "P"},
			expProject: model.Project{},
			expError:   store.ErrDbQuery,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(store, c, tc.project)
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
		mock     func(*mock_store.MockStore, *gomock.Controller, model.Project)
		project  model.Project
		expError error
	}{
		{
			name: "Project is deleted",
			mock: func(s *mock_store.MockStore, c *gomock.Controller, p model.Project) {
				pr := mock_store.NewMockProjectRepo(c)

				pr.EXPECT().DeleteByID(p.ID).Return(nil)
				s.EXPECT().Projects().Return(pr)
			},
			project:  model.Project{ID: 1, Name: "P"},
			expError: nil,
		},
		{
			name: "Error occures while deleting project",
			mock: func(s *mock_store.MockStore, c *gomock.Controller, p model.Project) {
				pr := mock_store.NewMockProjectRepo(c)

				pr.EXPECT().DeleteByID(p.ID).Return(store.ErrDbQuery)
				s.EXPECT().Projects().Return(pr)
			},
			project:  model.Project{ID: 1, Name: "P"},
			expError: store.ErrDbQuery,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(store, c, tc.project)
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
			name:     "Project passes validation",
			project:  model.Project{Name: "P"},
			expError: nil,
		},
		{
			name:     "Project doesn't pass validation because of empty name",
			project:  model.Project{},
			expError: ErrNameIsRequired,
		},
		{
			name:     "Project doesn't pass validation because of too long name",
			project:  model.Project{Name: fixedLengthString(501)},
			expError: ErrNameIsTooLong,
		},
		{
			name:     "Project doesn't pass validation because of too long description",
			project:  model.Project{Name: "P", Description: fixedLengthString(1001)},
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
