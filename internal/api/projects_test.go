package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	"github.com/imarrche/tasker/internal/model"
	mock_service "github.com/imarrche/tasker/internal/service/mocks"
)

func TestServer_ProjectList(t *testing.T) {
	server := &Server{router: mux.NewRouter()}
	server.configureRouter()

	testcases := []struct {
		name     string
		mock     func(*gomock.Controller, *mock_service.MockService, []model.Project)
		expCode  int
		projects []model.Project
		expBody  []model.Project
	}{
		{
			name: "project list is retrieved",
			mock: func(c *gomock.Controller, s *mock_service.MockService, projects []model.Project) {
				ps := mock_service.NewMockProjectService(c)
				ps.EXPECT().GetAll().Return(projects, nil)
				s.EXPECT().Projects().Return(ps)
			},
			expCode: http.StatusOK,
			projects: []model.Project{
				{ID: 1, Name: "Project 1"}, {ID: 2, Name: "Project 2"},
			},
			expBody: []model.Project{
				{ID: 1, Name: "Project 1"}, {ID: 2, Name: "Project 2"},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			s := mock_service.NewMockService(c)
			tc.mock(c, s, tc.projects)
			server.service = s

			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodGet, "/api/v1/projects", nil)

			server.router.ServeHTTP(w, r)
			var ps []model.Project
			err := json.NewDecoder(w.Body).Decode(&ps)

			assert.NoError(t, err)
			assert.Equal(t, tc.expCode, w.Code)
			assert.Equal(t, tc.expBody, ps)
		})
	}
}

func TestServer_ProjectCreate(t *testing.T) {
	server := &Server{router: mux.NewRouter()}
	server.configureRouter()

	testcases := []struct {
		name    string
		mock    func(*gomock.Controller, *mock_service.MockService, model.Project)
		project model.Project
		expCode int
		expBody model.Project
	}{
		{
			name: "project is created",
			mock: func(c *gomock.Controller, s *mock_service.MockService, p model.Project) {
				createdProject := p
				createdProject.ID = 1
				ps := mock_service.NewMockProjectService(c)
				ps.EXPECT().Create(p).Return(createdProject, nil)
				s.EXPECT().Projects().Return(ps)
			},
			project: model.Project{Name: "Project 1", Description: "Project description."},
			expCode: http.StatusCreated,
			expBody: model.Project{ID: 1, Name: "Project 1", Description: "Project description."},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			s := mock_service.NewMockService(c)
			tc.mock(c, s, tc.project)
			server.service = s

			w := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.project)
			r, _ := http.NewRequest(http.MethodPost, "/api/v1/projects", b)

			server.router.ServeHTTP(w, r)
			var p model.Project
			err := json.NewDecoder(w.Body).Decode(&p)

			assert.NoError(t, err)
			assert.Equal(t, tc.expCode, w.Code)
			assert.Equal(t, tc.expBody, p)
		})
	}
}

func TestServer_ProjectDetail(t *testing.T) {
	server := &Server{router: mux.NewRouter()}
	server.configureRouter()

	testcases := []struct {
		name    string
		mock    func(*gomock.Controller, *mock_service.MockService, model.Project)
		project model.Project
		expCode int
		expBody model.Project
	}{
		{
			name: "project is retrieved",
			mock: func(c *gomock.Controller, s *mock_service.MockService, p model.Project) {
				ps := mock_service.NewMockProjectService(c)
				ps.EXPECT().GetByID(p.ID).Return(p, nil)
				s.EXPECT().Projects().Return(ps)
			},
			project: model.Project{ID: 1, Name: "Project 1"},
			expCode: http.StatusOK,
			expBody: model.Project{ID: 1, Name: "Project 1"},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			s := mock_service.NewMockService(c)
			tc.mock(c, s, tc.project)
			server.service = s

			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodGet, "/api/v1/projects/1", nil)

			server.router.ServeHTTP(w, r)
			var p model.Project
			err := json.NewDecoder(w.Body).Decode(&p)

			assert.NoError(t, err)
			assert.Equal(t, tc.expCode, w.Code)
			assert.Equal(t, tc.expBody, p)
		})
	}
}

func TestServer_ProjectUpdate(t *testing.T) {
	server := &Server{router: mux.NewRouter()}
	server.configureRouter()

	testcases := []struct {
		name    string
		mock    func(*gomock.Controller, *mock_service.MockService, model.Project)
		project model.Project
		expCode int
		expBody model.Project
	}{
		{
			name: "project is updated",
			mock: func(c *gomock.Controller, s *mock_service.MockService, p model.Project) {
				ps := mock_service.NewMockProjectService(c)
				ps.EXPECT().Update(p).Return(p, nil)
				s.EXPECT().Projects().Return(ps)
			},
			project: model.Project{ID: 1, Name: "Updated project", Description: "Updated description"},
			expCode: http.StatusOK,
			expBody: model.Project{ID: 1, Name: "Updated project", Description: "Updated description"},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			s := mock_service.NewMockService(c)
			tc.mock(c, s, tc.project)
			server.service = s

			w := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.project)
			r, _ := http.NewRequest(http.MethodPut, "/api/v1/projects/1", b)

			server.router.ServeHTTP(w, r)
			var p model.Project
			err := json.NewDecoder(w.Body).Decode(&p)

			assert.NoError(t, err)
			assert.Equal(t, tc.expCode, w.Code)
			assert.Equal(t, tc.expBody, p)
		})
	}
}

func TestServer_ProjectDelete(t *testing.T) {
	server := &Server{router: mux.NewRouter()}
	server.configureRouter()

	testcases := []struct {
		name    string
		mock    func(*gomock.Controller, *mock_service.MockService, model.Project)
		project model.Project
		expCode int
	}{
		{
			name: "project is deleted",
			mock: func(c *gomock.Controller, s *mock_service.MockService, p model.Project) {
				ps := mock_service.NewMockProjectService(c)
				ps.EXPECT().DeleteByID(p.ID).Return(nil)
				s.EXPECT().Projects().Return(ps)
			},
			project: model.Project{ID: 1, Name: "Project 1"},
			expCode: http.StatusNoContent,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			s := mock_service.NewMockService(c)
			tc.mock(c, s, tc.project)
			server.service = s

			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodDelete, "/api/v1/projects/1", nil)

			server.router.ServeHTTP(w, r)

			assert.Equal(t, tc.expCode, w.Code)
		})
	}
}
