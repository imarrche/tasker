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
	"github.com/imarrche/tasker/internal/store/inmem"
)

func TestServer_ColumnList(t *testing.T) {
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

			s.columnList().ServeHTTP(w, r)
			var cs []model.Column
			err := json.NewDecoder(w.Body).Decode(&cs)

			assert.NoError(t, err)
			assert.Equal(t, tc.expCode, w.Code)
			assert.Equal(t, tc.expBody, cs)
		})
	}
}

func TestServer_ColumnCreate(t *testing.T) {
	s := NewServer(inmem.TestStoreWithFixtures())

	testcases := []struct {
		name    string
		column  model.Column
		expCode int
		expBody model.Column
	}{
		{
			name:    "Ok, column is created",
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

			s.columnCreate().ServeHTTP(w, r)
			var c model.Column
			err := json.NewDecoder(w.Body).Decode(&c)

			assert.NoError(t, err)
			assert.Equal(t, tc.expCode, w.Code)
			assert.Equal(t, tc.expBody, c)
		})
	}
}

func TestServer_ColumnDetail(t *testing.T) {
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
			url := "columns/column_id"
			r, _ := http.NewRequest(http.MethodGet, url, nil)
			r = mux.SetURLVars(r, map[string]string{
				"column_id": "1",
			})

			s.columnDetail().ServeHTTP(w, r)
			var c model.Column
			err := json.NewDecoder(w.Body).Decode(&c)

			assert.NoError(t, err)
			assert.Equal(t, tc.expCode, w.Code)
			assert.Equal(t, tc.expBody, c)
		})
	}
}

func TestServer_ColumnMove(t *testing.T) {
	testcases := []struct {
		name    string
		left    bool
		id      string
		expCode int
	}{
		{
			name:    "Ok, column is moved left",
			id:      "2",
			left:    true,
			expCode: http.StatusOK,
		},
		{
			name:    "Ok, column is moved right",
			id:      "1",
			left:    false,
			expCode: http.StatusOK,
		},
	}

	type request struct {
		Left bool `json:"left"`
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			s := NewServer(inmem.TestStoreWithFixtures())
			w := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(&request{Left: tc.left})
			url := "columns/column_id/move"
			r, _ := http.NewRequest(http.MethodPost, url, b)
			r = mux.SetURLVars(r, map[string]string{
				"column_id": tc.id,
			})

			s.columnMove().ServeHTTP(w, r)

			assert.Equal(t, tc.expCode, w.Code)
		})
	}
}

func TestServer_ColumnUpdate(t *testing.T) {
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
			url := "columns/column_id"
			r, _ := http.NewRequest(http.MethodPut, url, b)
			r = mux.SetURLVars(r, map[string]string{
				"column_id": "1",
			})

			s.columnUpdate().ServeHTTP(w, r)
			var c model.Column
			err := json.NewDecoder(w.Body).Decode(&c)

			assert.NoError(t, err)
			assert.Equal(t, tc.expCode, w.Code)
			assert.Equal(t, tc.expBody, c)
		})
	}
}

func TestServer_ColumnDelete(t *testing.T) {
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
			url := "columns/column_id"
			r, _ := http.NewRequest(http.MethodDelete, url, nil)
			r = mux.SetURLVars(r, map[string]string{
				"column_id": "1",
			})

			s.columnDelete().ServeHTTP(w, r)

			assert.Equal(t, tc.expCode, w.Code)
		})
	}
}