package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"

	"github.com/imarrche/tasker/internal/model"
	"github.com/imarrche/tasker/internal/store/inmem"
	"github.com/stretchr/testify/assert"
)

func TestServer_HandleColumnList(t *testing.T) {
	s := NewServer(inmem.TestStoreWithFixtures())

	testcases := []struct {
		name    string
		expCode int
		expBody []model.Column
	}{
		{
			name:    "Ok, column list is retrieved",
			expCode: http.StatusOK,
			expBody: []model.Column{
				model.Column{ID: 1, Name: "Column 1", Index: 1, ProjectID: 1},
				model.Column{ID: 2, Name: "Column 2", Index: 2, ProjectID: 1},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodGet, "projects/project_id/columns", nil)
			r = mux.SetURLVars(r, map[string]string{
				"project_id": "1",
			})

			s.handleColumnList().ServeHTTP(w, r)
			var cs []model.Column
			err := json.NewDecoder(w.Body).Decode(&cs)

			assert.NoError(t, err)
			assert.Equal(t, tc.expCode, w.Code)
			assert.Equal(t, tc.expBody, cs)
		})
	}
}

func TestServer_HandleColumnCreate(t *testing.T) {
	s := NewServer(inmem.TestStoreWithFixtures())

	testcases := []struct {
		name    string
		column  model.Column
		expCode int
		expBody model.Column
	}{
		{
			name:    "Ok, column is craeted",
			column:  model.Column{Name: "Column 4"},
			expCode: http.StatusCreated,
			expBody: model.Column{ID: 4, Name: "Column 4", Index: 3, ProjectID: 1},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.column)
			r, _ := http.NewRequest(http.MethodPost, "projects/project_id/columns", b)
			r = mux.SetURLVars(r, map[string]string{
				"project_id": "1",
			})

			s.handleColumnCreate().ServeHTTP(w, r)
			var c model.Column
			err := json.NewDecoder(w.Body).Decode(&c)

			assert.NoError(t, err)
			assert.Equal(t, tc.expCode, w.Code)
			assert.Equal(t, tc.expBody, c)
		})
	}
}

func TestServer_HandleColumnDetail(t *testing.T) {
	s := NewServer(inmem.TestStoreWithFixtures())

	testcases := []struct {
		name    string
		expCode int
		expBody model.Column
	}{
		{
			name:    "Ok, column is retrieved",
			expCode: http.StatusOK,
			expBody: model.Column{ID: 1, Name: "Column 1", Index: 1, ProjectID: 1},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			url := "projects/project_id/columns/column_id"
			r, _ := http.NewRequest(http.MethodGet, url, nil)
			r = mux.SetURLVars(r, map[string]string{
				"project_id": "1",
				"column_id":  "1",
			})

			s.handleColumnDetail().ServeHTTP(w, r)
			var c model.Column
			err := json.NewDecoder(w.Body).Decode(&c)

			assert.NoError(t, err)
			assert.Equal(t, tc.expCode, w.Code)
			assert.Equal(t, tc.expBody, c)
		})
	}
}

func TestServer_HandleColumnUpdate(t *testing.T) {
	s := NewServer(inmem.TestStoreWithFixtures())

	testcases := []struct {
		name    string
		column  model.Column
		expCode int
		expBody model.Column
	}{
		{
			name:    "Ok, column is updated",
			column:  model.Column{Name: "Updated column"},
			expCode: http.StatusOK,
			expBody: model.Column{ID: 1, Name: "Updated column", Index: 1, ProjectID: 1},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.column)
			url := "projects/project_id/columns/column_id"
			r, _ := http.NewRequest(http.MethodPut, url, b)
			r = mux.SetURLVars(r, map[string]string{
				"project_id": "1",
				"column_id":  "1",
			})

			s.handleColumnUpdate().ServeHTTP(w, r)
			var c model.Column
			err := json.NewDecoder(w.Body).Decode(&c)

			assert.NoError(t, err)
			assert.Equal(t, tc.expCode, w.Code)
			assert.Equal(t, tc.expBody, c)
		})
	}
}

func TestServer_HandleColumnDelete(t *testing.T) {
	s := NewServer(inmem.TestStoreWithFixtures())

	testcases := []struct {
		name    string
		expCode int
	}{
		{
			name:    "Ok, column is deleted",
			expCode: http.StatusNoContent,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			url := "projects/project_id/columns/column_id"
			r, _ := http.NewRequest(http.MethodDelete, url, nil)
			r = mux.SetURLVars(r, map[string]string{
				"project_id": "1",
				"column_id":  "1",
			})

			s.handleColumnDelete().ServeHTTP(w, r)

			assert.Equal(t, tc.expCode, w.Code)
		})
	}
}
