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

func TestServer_V1ProjectList(t *testing.T) {
	s := NewTestServer()

	testcases := []struct {
		name    string
		expCode int
		expBody []model.Project
	}{
		{
			name:    "Ok, project list is retrieved",
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

func TestServer_V1ProjectCreate(t *testing.T) {
	s := NewTestServer()

	testcases := []struct {
		name    string
		project model.Project
		expCode int
		expBody model.Project
	}{
		{
			name:    "Ok, project is created",
			project: model.Project{Name: "Project 3"},
			expCode: http.StatusCreated,
			expBody: model.Project{ID: 3, Name: "Project 3"},
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

func TestServer_V1ProjectDetail(t *testing.T) {
	s := NewTestServer()

	testcases := []struct {
		name    string
		id      int
		expCode int
		expBody model.Project
	}{
		{
			name:    "Ok, project is retrieved",
			id:      1,
			expCode: http.StatusOK,
			expBody: model.Project{ID: 1, Name: "Project 1"},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodGet, "/api/v1/projects/project_id", nil)
			r = mux.SetURLVars(r, map[string]string{
				"project_id": "1",
			})

			s.projectDetail().ServeHTTP(w, r)
			var p model.Project
			err := json.NewDecoder(w.Body).Decode(&p)

			assert.NoError(t, err)
			assert.Equal(t, tc.expCode, w.Code)
			assert.Equal(t, tc.expBody, p)
		})
	}
}

func TestServer_V1ProjectUpdate(t *testing.T) {
	s := NewTestServer()

	testcases := []struct {
		name    string
		project model.Project
		expCode int
		expBody model.Project
	}{
		{
			name:    "Ok, project is updated",
			project: model.Project{Name: "Updated project"},
			expCode: http.StatusOK,
			expBody: model.Project{ID: 1, Name: "Updated project"},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.project)
			r, _ := http.NewRequest(http.MethodPut, "/api/v1/projects/project_id", b)
			r = mux.SetURLVars(r, map[string]string{
				"project_id": "1",
			})

			s.projectUpdate().ServeHTTP(w, r)
			var p model.Project
			err := json.NewDecoder(w.Body).Decode(&p)

			assert.NoError(t, err)
			assert.Equal(t, tc.expCode, w.Code)
			assert.Equal(t, tc.expBody, p)
		})
	}
}

func TestServer_V1ProjectDelete(t *testing.T) {
	s := NewTestServer()

	testcases := []struct {
		name    string
		id      int
		expCode int
	}{
		{
			name:    "Ok, project is deleted",
			id:      1,
			expCode: http.StatusNoContent,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodDelete, "/api/v1/projects/project_id", nil)
			r = mux.SetURLVars(r, map[string]string{
				"project_id": "1",
			})

			s.projectDelete().ServeHTTP(w, r)

			assert.Equal(t, tc.expCode, w.Code)
		})
	}
}
