package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	"github.com/imarrche/tasker/internal/model"
)

func TestServer_ProjectList(t *testing.T) {
	s := NewTestServer()

	testcases := []struct {
		name    string
		expCode int
		expBody []model.Project
	}{
		{
			name:    "project list is retrieved",
			expCode: http.StatusOK,
			expBody: []model.Project{
				model.Project{ID: 1, Name: "Project 1"},
				model.Project{ID: 2, Name: "Project 2"},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodGet, "/api/v1/projects", nil)

			s.projectList().ServeHTTP(w, r)
			var ps []model.Project
			err := json.NewDecoder(w.Body).Decode(&ps)

			assert.NoError(t, err)
			assert.Equal(t, tc.expCode, w.Code)
			assert.Equal(t, tc.expBody, ps)
		})
	}
}

func TestServer_ProjectCreate(t *testing.T) {
	s := NewTestServer()

	testcases := []struct {
		name    string
		project model.Project
		expCode int
		expBody model.Project
	}{
		{
			name:    "project is created",
			project: model.Project{Name: "Project 3", Description: "Project description."},
			expCode: http.StatusCreated,
			expBody: model.Project{ID: 3, Name: "Project 3", Description: "Project description."},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.project)
			r, _ := http.NewRequest(http.MethodPost, "/api/v1/projects", b)

			s.projectCreate().ServeHTTP(w, r)
			var p model.Project
			err := json.NewDecoder(w.Body).Decode(&p)

			assert.NoError(t, err)
			assert.Equal(t, tc.expCode, w.Code)
			assert.Equal(t, tc.expBody, p)
		})
	}
}

func TestServer_ProjectDetail(t *testing.T) {
	s := NewTestServer()

	testcases := []struct {
		name    string
		id      string
		expCode int
		expBody model.Project
	}{
		{
			name:    "project is retrieved",
			id:      "1",
			expCode: http.StatusOK,
			expBody: model.Project{ID: 1, Name: "Project 1"},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodGet, "/api/v1/projects/project_id", nil)
			r = mux.SetURLVars(r, map[string]string{"project_id": tc.id})

			s.projectDetail().ServeHTTP(w, r)
			var p model.Project
			err := json.NewDecoder(w.Body).Decode(&p)

			assert.NoError(t, err)
			assert.Equal(t, tc.expCode, w.Code)
			assert.Equal(t, tc.expBody, p)
		})
	}
}

func TestServer_ProjectUpdate(t *testing.T) {
	s := NewTestServer()

	testcases := []struct {
		name    string
		id      string
		project model.Project
		expCode int
		expBody model.Project
	}{
		{
			name:    "project is updated",
			id:      "1",
			project: model.Project{Name: "Updated project", Description: "Updated description"},
			expCode: http.StatusOK,
			expBody: model.Project{ID: 1, Name: "Updated project", Description: "Updated description"},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.project)
			r, _ := http.NewRequest(http.MethodPut, "/api/v1/projects/project_id", b)
			r = mux.SetURLVars(r, map[string]string{"project_id": tc.id})

			s.projectUpdate().ServeHTTP(w, r)
			var p model.Project
			err := json.NewDecoder(w.Body).Decode(&p)

			assert.NoError(t, err)
			assert.Equal(t, tc.expCode, w.Code)
			assert.Equal(t, tc.expBody, p)
		})
	}
}

func TestServer_ProjectDelete(t *testing.T) {
	s := NewTestServer()

	testcases := []struct {
		name    string
		id      string
		expCode int
	}{
		{
			name:    "project is deleted",
			id:      "1",
			expCode: http.StatusNoContent,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodDelete, "/api/v1/projects/project_id", nil)
			r = mux.SetURLVars(r, map[string]string{"project_id": tc.id})

			s.projectDelete().ServeHTTP(w, r)

			assert.Equal(t, tc.expCode, w.Code)
		})
	}
}
