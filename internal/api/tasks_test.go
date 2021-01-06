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

func TestServer_TaskList(t *testing.T) {
	s := NewTestServer()

	testcases := []struct {
		name     string
		columnID string
		expCode  int
		expBody  []model.Task
	}{
		{
			name:     "task list is retrieved",
			columnID: "1",
			expCode:  http.StatusOK,
			expBody: []model.Task{
				{ID: 1, Name: "Task 1", Index: 1, ColumnID: 1},
				{ID: 2, Name: "Task 2", Index: 2, ColumnID: 1},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodGet, "/api/v1/columns/column_id/tasks", nil)
			r = mux.SetURLVars(r, map[string]string{"column_id": tc.columnID})

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
	s := NewTestServer()

	testcases := []struct {
		name     string
		columnID string
		task     model.Task
		expCode  int
		expBody  model.Task
	}{
		{
			name:     "task is created",
			columnID: "1",
			task:     model.Task{Name: "Task 3", Index: 3, ColumnID: 1},
			expCode:  http.StatusCreated,
			expBody:  model.Task{ID: 4, Name: "Task 3", Index: 3, ColumnID: 1},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.task)
			r, _ := http.NewRequest(http.MethodPost, "/api/v1/columns/column_id/tasks", b)
			r = mux.SetURLVars(r, map[string]string{"column_id": tc.columnID})

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
	s := NewTestServer()

	testcases := []struct {
		name    string
		id      string
		expCode int
		expBody model.Task
	}{
		{
			name:    "task is retrieved",
			id:      "1",
			expCode: http.StatusOK,
			expBody: model.Task{ID: 1, Name: "Task 1", Index: 1, ColumnID: 1},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodGet, "/api/v1/tasks/task_id", nil)
			r = mux.SetURLVars(r, map[string]string{"task_id": tc.id})

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
	s := NewTestServer()

	testcases := []struct {
		name    string
		id      string
		left    bool
		expCode int
	}{
		{
			name:    "task is moved right",
			id:      "1",
			left:    false,
			expCode: http.StatusOK,
		},
		{
			name:    "task is moved left",
			id:      "1",
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
			r, _ := http.NewRequest(http.MethodPost, "/api/v1/tasks/task_id/movex", b)
			r = mux.SetURLVars(r, map[string]string{"task_id": tc.id})

			s.taskMoveX().ServeHTTP(w, r)

			assert.Equal(t, tc.expCode, w.Code)
		})
	}
}

func TestServer_TaskMoveY(t *testing.T) {
	s := NewTestServer()

	testcases := []struct {
		name    string
		id      string
		up      bool
		expCode int
	}{
		{
			name:    "task is moved down",
			id:      "1",
			up:      false,
			expCode: http.StatusOK,
		},
		{
			name:    "task is moved up",
			id:      "1",
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
			r, _ := http.NewRequest(http.MethodPost, "/api/v1/tasks/task_id/movey", b)
			r = mux.SetURLVars(r, map[string]string{"task_id": tc.id})

			s.taskMoveY().ServeHTTP(w, r)

			assert.Equal(t, tc.expCode, w.Code)
		})
	}
}

func TestServer_TaskUpdate(t *testing.T) {
	s := NewTestServer()

	testcases := []struct {
		name    string
		id      string
		task    model.Task
		expCode int
		expBody model.Task
	}{
		{
			name:    "task is updated",
			id:      "1",
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
			r = mux.SetURLVars(r, map[string]string{"task_id": tc.id})

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
	s := NewTestServer()

	testcases := []struct {
		name    string
		id      string
		expCode int
	}{
		{
			name:    "task is deleted",
			id:      "1",
			expCode: http.StatusNoContent,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodDelete, "/api/v1/tasks/task_id", nil)
			r = mux.SetURLVars(r, map[string]string{"task_id": tc.id})

			s.taskDelete().ServeHTTP(w, r)

			assert.Equal(t, tc.expCode, w.Code)
		})
	}
}
