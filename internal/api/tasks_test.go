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

func TestServer_TaskList(t *testing.T) {
	s := NewServer(inmem.TestStoreWithFixtures())

	testcases := []struct {
		name    string
		expCode int
		expBody []model.Task
	}{
		{
			name:    "Ok, task list is retrieved",
			expCode: http.StatusOK,
			expBody: []model.Task{
				model.Task{ID: 1, Name: "Task 1", Index: 1, ColumnID: 1},
				model.Task{ID: 2, Name: "Task 2", Index: 2, ColumnID: 1},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodGet, "/api/v1/columns/column_id/tasks", nil)
			r = mux.SetURLVars(r, map[string]string{
				"column_id": "1",
			})

			s.taskList().ServeHTTP(w, r)
			var ts []model.Task
			err := json.NewDecoder(w.Body).Decode(&ts)

			assert.NoError(t, err)
			assert.Equal(t, tc.expCode, w.Code)
			assert.Equal(t, tc.expBody, ts)
		})
	}
}

func TestServer_TaskСreate(t *testing.T) {
	s := NewServer(inmem.TestStoreWithFixtures())

	testcases := []struct {
		name    string
		task    model.Task
		expCode int
		expBody model.Task
	}{
		{
			name:    "Ok, task is created",
			task:    model.Task{Name: "Task 3", Index: 3, ColumnID: 1},
			expCode: http.StatusCreated,
			expBody: model.Task{ID: 4, Name: "Task 3", Index: 3, ColumnID: 1},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.task)
			r, _ := http.NewRequest(http.MethodPost, "/api/v1/columns/column_id/tasks", b)
			r = mux.SetURLVars(r, map[string]string{
				"column_id": "1",
			})

			s.taskCreate().ServeHTTP(w, r)
			var task model.Task
			err := json.NewDecoder(w.Body).Decode(&task)

			assert.NoError(t, err)
			assert.Equal(t, tc.expCode, w.Code)
			assert.Equal(t, tc.expBody, task)
		})
	}
}

func TestServer_TaskDetail(t *testing.T) {
	s := NewServer(inmem.TestStoreWithFixtures())

	testcases := []struct {
		name    string
		expCode int
		expBody model.Task
	}{
		{
			name:    "Ok, task is retrieved",
			expCode: http.StatusOK,
			expBody: model.Task{ID: 1, Name: "Task 1", Index: 1, ColumnID: 1},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodGet, "/api/v1/tasks/task_id", nil)
			r = mux.SetURLVars(r, map[string]string{
				"task_id": "1",
			})

			s.taskDetail().ServeHTTP(w, r)
			var task model.Task
			err := json.NewDecoder(w.Body).Decode(&task)

			assert.NoError(t, err)
			assert.Equal(t, tc.expCode, w.Code)
			assert.Equal(t, tc.expBody, task)
		})
	}
}

func TestServer_TaskMoveX(t *testing.T) {
	s := NewServer(inmem.TestStoreWithFixtures())

	testcases := []struct {
		name    string
		left    bool
		expCode int
	}{
		{
			name:    "Ok, task is moved right",
			left:    false,
			expCode: http.StatusOK,
		},
		{
			name:    "Task is not moved left, because it's the first one",
			left:    true,
			expCode: http.StatusOK,
		},
	}

	type request struct {
		Left bool `json:"left"`
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(request{Left: tc.left})
			r, _ := http.NewRequest(http.MethodPost, "/api/v1/tasks/task_id", b)
			r = mux.SetURLVars(r, map[string]string{
				"task_id": "1",
			})

			s.taskMoveX().ServeHTTP(w, r)

			assert.Equal(t, tc.expCode, w.Code)
		})
	}
}

func TestServer_TaskMoveY(t *testing.T) {
	s := NewServer(inmem.TestStoreWithFixtures())

	testcases := []struct {
		name    string
		up      bool
		expCode int
	}{
		{
			name:    "Ok, task is moved down",
			up:      false,
			expCode: http.StatusOK,
		},
		{
			name:    "Task is not moved up, because it's the first one",
			up:      true,
			expCode: http.StatusOK,
		},
	}

	type request struct {
		Up bool `json:"up"`
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(request{Up: tc.up})
			r, _ := http.NewRequest(http.MethodPost, "/api/v1/tasks/task_id", b)
			r = mux.SetURLVars(r, map[string]string{
				"task_id": "1",
			})

			s.taskMoveY().ServeHTTP(w, r)

			assert.Equal(t, tc.expCode, w.Code)
		})
	}
}

func TestServer_TaskUpdate(t *testing.T) {
	s := NewServer(inmem.TestStoreWithFixtures())

	testcases := []struct {
		name    string
		task    model.Task
		expCode int
		expBody model.Task
	}{
		{
			name:    "Ok, task is update",
			task:    model.Task{Name: "Updated task", Index: 1, ColumnID: 1},
			expCode: http.StatusOK,
			expBody: model.Task{ID: 1, Name: "Updated task", Index: 1, ColumnID: 1},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.task)
			r, _ := http.NewRequest(http.MethodPut, "/api/v1/tasks/task_id", b)
			r = mux.SetURLVars(r, map[string]string{
				"task_id": "1",
			})

			s.taskUpdate().ServeHTTP(w, r)
			var task model.Task
			err := json.NewDecoder(w.Body).Decode(&task)

			assert.NoError(t, err)
			assert.Equal(t, tc.expCode, w.Code)
			assert.Equal(t, tc.expBody, task)
		})
	}
}

func TestServer_TaskDelete(t *testing.T) {
	s := NewServer(inmem.TestStoreWithFixtures())

	testcases := []struct {
		name    string
		expCode int
	}{
		{
			name:    "Ok, task is deleted",
			expCode: http.StatusNoContent,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodDelete, "/api/v1/tasks/task_id", nil)
			r = mux.SetURLVars(r, map[string]string{
				"task_id": "1",
			})

			s.taskDelete().ServeHTTP(w, r)

			assert.Equal(t, tc.expCode, w.Code)
		})
	}
}
