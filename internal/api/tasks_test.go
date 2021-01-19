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

func TestServer_TaskList(t *testing.T) {
	server := &Server{router: mux.NewRouter()}
	server.configureRouter()

	testcases := []struct {
		name     string
		mock     func(*gomock.Controller, *mock_service.MockService, int, []model.Task)
		columnID int
		tasks    []model.Task
		expCode  int
		expBody  []model.Task
	}{
		{
			name: "task list is retrieved",
			mock: func(c *gomock.Controller, s *mock_service.MockService, cID int, tasks []model.Task) {
				ts := mock_service.NewMockTaskService(c)
				ts.EXPECT().GetByColumnID(cID).Return(tasks, nil)
				s.EXPECT().Tasks().Return(ts)
			},
			columnID: 1,
			tasks: []model.Task{
				{ID: 1, Name: "Task 1", Index: 1, ColumnID: 1},
				{ID: 2, Name: "Task 2", Index: 2, ColumnID: 1},
			},
			expCode: http.StatusOK,
			expBody: []model.Task{
				{ID: 1, Name: "Task 1", Index: 1, ColumnID: 1},
				{ID: 2, Name: "Task 2", Index: 2, ColumnID: 1},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			s := mock_service.NewMockService(c)
			tc.mock(c, s, tc.columnID, tc.tasks)
			server.service = s

			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodGet, "/api/v1/columns/1/tasks", nil)

			server.router.ServeHTTP(w, r)
			var ts []model.Task
			err := json.NewDecoder(w.Body).Decode(&ts)

			assert.NoError(t, err)
			assert.Equal(t, tc.expCode, w.Code)
			assert.Equal(t, tc.expBody, ts)
		})
	}
}

func TestServer_TaskСreate(t *testing.T) {
	server := &Server{router: mux.NewRouter()}
	server.configureRouter()

	testcases := []struct {
		name     string
		mock     func(*gomock.Controller, *mock_service.MockService, int, model.Task)
		columnID int
		task     model.Task
		expCode  int
		expBody  model.Task
	}{
		{
			name: "task is created",
			mock: func(c *gomock.Controller, s *mock_service.MockService, cID int, task model.Task) {
				createdTask := model.Task{
					ID: 1, Name: task.Name, Index: 1, ColumnID: task.ColumnID,
				}
				ts := mock_service.NewMockTaskService(c)
				ts.EXPECT().Create(task).Return(createdTask, nil)
				s.EXPECT().Tasks().Return(ts)
			},
			columnID: 1,
			task:     model.Task{Name: "Task 1", ColumnID: 1},
			expCode:  http.StatusCreated,
			expBody:  model.Task{ID: 1, Name: "Task 1", Index: 1, ColumnID: 1},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			s := mock_service.NewMockService(c)
			tc.mock(c, s, tc.columnID, tc.task)
			server.service = s

			w := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.task)
			r, _ := http.NewRequest(http.MethodPost, "/api/v1/columns/1/tasks", b)

			server.router.ServeHTTP(w, r)
			var task model.Task
			err := json.NewDecoder(w.Body).Decode(&task)

			assert.NoError(t, err)
			assert.Equal(t, tc.expCode, w.Code)
			assert.Equal(t, tc.expBody, task)
		})
	}
}

func TestServer_TaskDetail(t *testing.T) {
	server := &Server{router: mux.NewRouter()}
	server.configureRouter()

	testcases := []struct {
		name    string
		mock    func(*gomock.Controller, *mock_service.MockService, model.Task)
		task    model.Task
		expCode int
		expBody model.Task
	}{
		{
			name: "task is retrieved",
			mock: func(c *gomock.Controller, s *mock_service.MockService, task model.Task) {
				ts := mock_service.NewMockTaskService(c)
				ts.EXPECT().GetByID(task.ID).Return(task, nil)
				s.EXPECT().Tasks().Return(ts)
			},
			task:    model.Task{ID: 1, Name: "Task 1", Index: 1, ColumnID: 1},
			expCode: http.StatusOK,
			expBody: model.Task{ID: 1, Name: "Task 1", Index: 1, ColumnID: 1},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			s := mock_service.NewMockService(c)
			tc.mock(c, s, tc.task)
			server.service = s

			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodGet, "/api/v1/tasks/1", nil)

			server.router.ServeHTTP(w, r)
			var task model.Task
			err := json.NewDecoder(w.Body).Decode(&task)

			assert.NoError(t, err)
			assert.Equal(t, tc.expCode, w.Code)
			assert.Equal(t, tc.expBody, task)
		})
	}
}

func TestServer_TaskMoveX(t *testing.T) {
	server := &Server{router: mux.NewRouter()}
	server.configureRouter()

	testcases := []struct {
		name    string
		mock    func(*gomock.Controller, *mock_service.MockService, bool, model.Task)
		task    model.Task
		left    bool
		expCode int
	}{
		{
			name: "task is moved right",
			mock: func(c *gomock.Controller, s *mock_service.MockService, left bool, task model.Task) {
				ts := mock_service.NewMockTaskService(c)
				ts.EXPECT().MoveToColumnByID(task.ID, left).Return(nil)
				s.EXPECT().Tasks().Return(ts)
			},
			task:    model.Task{ID: 1, Name: "Task 1", Index: 1, ColumnID: 1},
			left:    false,
			expCode: http.StatusOK,
		},
		{
			name: "task is moved left",
			mock: func(c *gomock.Controller, s *mock_service.MockService, left bool, task model.Task) {
				ts := mock_service.NewMockTaskService(c)
				ts.EXPECT().MoveToColumnByID(task.ID, left).Return(nil)
				s.EXPECT().Tasks().Return(ts)
			},
			task:    model.Task{ID: 1, Name: "Task 1", Index: 1, ColumnID: 2},
			left:    true,
			expCode: http.StatusOK,
		},
	}

	type request struct {
		Left bool `json:"left"`
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			s := mock_service.NewMockService(c)
			tc.mock(c, s, tc.left, tc.task)
			server.service = s

			w := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(request{Left: tc.left})
			r, _ := http.NewRequest(http.MethodPost, "/api/v1/tasks/1/movex", b)

			server.router.ServeHTTP(w, r)

			assert.Equal(t, tc.expCode, w.Code)
		})
	}
}

func TestServer_TaskMoveY(t *testing.T) {
	server := &Server{router: mux.NewRouter()}
	server.configureRouter()

	testcases := []struct {
		name    string
		mock    func(*gomock.Controller, *mock_service.MockService, bool, model.Task)
		task    model.Task
		up      bool
		expCode int
	}{
		{
			name: "task is moved down",
			mock: func(c *gomock.Controller, s *mock_service.MockService, up bool, task model.Task) {
				ts := mock_service.NewMockTaskService(c)
				ts.EXPECT().MoveByID(task.ID, up).Return(nil)
				s.EXPECT().Tasks().Return(ts)
			},
			task:    model.Task{ID: 1, Name: "Task 1", Index: 1, ColumnID: 1},
			up:      false,
			expCode: http.StatusOK,
		},
		{
			name: "task is moved up",
			mock: func(c *gomock.Controller, s *mock_service.MockService, up bool, task model.Task) {
				ts := mock_service.NewMockTaskService(c)
				ts.EXPECT().MoveByID(task.ID, up).Return(nil)
				s.EXPECT().Tasks().Return(ts)
			},
			task:    model.Task{ID: 1, Name: "Task 1", Index: 2, ColumnID: 1},
			up:      true,
			expCode: http.StatusOK,
		},
	}

	type request struct {
		Up bool `json:"up"`
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			s := mock_service.NewMockService(c)
			tc.mock(c, s, tc.up, tc.task)
			server.service = s

			w := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(request{Up: tc.up})
			r, _ := http.NewRequest(http.MethodPost, "/api/v1/tasks/1/movey", b)

			server.router.ServeHTTP(w, r)

			assert.Equal(t, tc.expCode, w.Code)
		})
	}
}

func TestServer_TaskUpdate(t *testing.T) {
	server := &Server{router: mux.NewRouter()}
	server.configureRouter()

	testcases := []struct {
		name    string
		mock    func(*gomock.Controller, *mock_service.MockService, model.Task)
		task    model.Task
		expCode int
		expBody model.Task
	}{
		{
			name: "task is updated",
			mock: func(c *gomock.Controller, s *mock_service.MockService, task model.Task) {
				updatedTask := model.Task{
					ID: 1, Name: task.Name, Description: task.Description, Index: 1, ColumnID: 1,
				}
				ts := mock_service.NewMockTaskService(c)
				ts.EXPECT().Update(task).Return(updatedTask, nil)
				s.EXPECT().Tasks().Return(ts)
			},
			task:    model.Task{ID: 1, Name: "Updated task", Description: "Task description."},
			expCode: http.StatusOK,
			expBody: model.Task{ID: 1, Name: "Updated task", Description: "Task description.", Index: 1, ColumnID: 1},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			s := mock_service.NewMockService(c)
			tc.mock(c, s, tc.task)
			server.service = s

			w := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.task)
			r, _ := http.NewRequest(http.MethodPut, "/api/v1/tasks/1", b)

			server.router.ServeHTTP(w, r)
			var task model.Task
			err := json.NewDecoder(w.Body).Decode(&task)

			assert.NoError(t, err)
			assert.Equal(t, tc.expCode, w.Code)
			assert.Equal(t, tc.expBody, task)
		})
	}
}

func TestServer_TaskDelete(t *testing.T) {
	server := &Server{router: mux.NewRouter()}
	server.configureRouter()

	testcases := []struct {
		name    string
		mock    func(*gomock.Controller, *mock_service.MockService, model.Task)
		task    model.Task
		expCode int
	}{
		{
			name: "task is deleted",
			mock: func(c *gomock.Controller, s *mock_service.MockService, task model.Task) {
				ts := mock_service.NewMockTaskService(c)
				ts.EXPECT().DeleteByID(task.ID).Return(nil)
				s.EXPECT().Tasks().Return(ts)
			},
			task:    model.Task{ID: 1, Name: "Task 1"},
			expCode: http.StatusNoContent,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			s := mock_service.NewMockService(c)
			tc.mock(c, s, tc.task)
			server.service = s

			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodDelete, "/api/v1/tasks/1", nil)

			server.router.ServeHTTP(w, r)

			assert.Equal(t, tc.expCode, w.Code)
		})
	}
}
