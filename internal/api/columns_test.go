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

func TestServer_ColumnList(t *testing.T) {
	server := &Server{router: mux.NewRouter()}
	server.configureRouter()

	testcases := []struct {
		name      string
		mock      func(*gomock.Controller, *mock_service.MockService, int, []model.Column)
		projectID int
		columns   []model.Column
		expCode   int
		expBody   []model.Column
	}{
		{
			name: "column list is retrieved",
			mock: func(c *gomock.Controller, s *mock_service.MockService, pID int, columns []model.Column) {
				cs := mock_service.NewMockColumnService(c)
				cs.EXPECT().GetByProjectID(pID).Return(columns, nil)
				s.EXPECT().Columns().Return(cs)
			},
			projectID: 1,
			columns: []model.Column{
				{ID: 1, Name: "Column 1", Index: 1, ProjectID: 1},
				{ID: 2, Name: "Column 2", Index: 2, ProjectID: 1},
			},
			expCode: http.StatusOK,
			expBody: []model.Column{
				{ID: 1, Name: "Column 1", Index: 1, ProjectID: 1},
				{ID: 2, Name: "Column 2", Index: 2, ProjectID: 1},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			s := mock_service.NewMockService(c)
			tc.mock(c, s, tc.projectID, tc.columns)
			server.service = s

			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodGet, "/api/v1/projects/1/columns", nil)

			server.router.ServeHTTP(w, r)
			var cs []model.Column
			err := json.NewDecoder(w.Body).Decode(&cs)

			assert.NoError(t, err)
			assert.Equal(t, tc.expCode, w.Code)
			assert.Equal(t, tc.expBody, cs)
		})
	}
}

func TestServer_ColumnCreate(t *testing.T) {
	server := &Server{router: mux.NewRouter()}
	server.configureRouter()

	testcases := []struct {
		name      string
		mock      func(*gomock.Controller, *mock_service.MockService, int, model.Column)
		projectID int
		column    model.Column
		expCode   int
		expBody   model.Column
	}{
		{
			name: "column is created",
			mock: func(c *gomock.Controller, s *mock_service.MockService, pID int, column model.Column) {
				createdColumn := model.Column{
					ID: 1, Name: column.Name, Index: 1, ProjectID: column.ProjectID,
				}
				cs := mock_service.NewMockColumnService(c)
				cs.EXPECT().Create(column).Return(createdColumn, nil)
				s.EXPECT().Columns().Return(cs)
			},
			projectID: 1,
			column:    model.Column{Name: "Column 1", ProjectID: 1},
			expCode:   http.StatusCreated,
			expBody:   model.Column{ID: 1, Name: "Column 1", Index: 1, ProjectID: 1},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			s := mock_service.NewMockService(c)
			tc.mock(c, s, tc.projectID, tc.column)
			server.service = s

			w := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.column)
			r, _ := http.NewRequest(http.MethodPost, "/api/v1/projects/1/columns", b)

			server.router.ServeHTTP(w, r)
			var column model.Column
			err := json.NewDecoder(w.Body).Decode(&column)

			assert.NoError(t, err)
			assert.Equal(t, tc.expCode, w.Code)
			assert.Equal(t, tc.expBody, column)
		})
	}
}

func TestServer_ColumnDetail(t *testing.T) {
	server := &Server{router: mux.NewRouter()}
	server.configureRouter()

	testcases := []struct {
		name    string
		mock    func(c *gomock.Controller, s *mock_service.MockService, column model.Column)
		column  model.Column
		expCode int
		expBody model.Column
	}{
		{
			name: "column is retrieved",
			mock: func(c *gomock.Controller, s *mock_service.MockService, column model.Column) {
				cs := mock_service.NewMockColumnService(c)
				cs.EXPECT().GetByID(column.ID).Return(column, nil)
				s.EXPECT().Columns().Return(cs)
			},
			column:  model.Column{ID: 1, Name: "Column 1", Index: 1, ProjectID: 1},
			expCode: http.StatusOK,
			expBody: model.Column{ID: 1, Name: "Column 1", Index: 1, ProjectID: 1},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			s := mock_service.NewMockService(c)
			tc.mock(c, s, tc.column)
			server.service = s

			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodGet, "/api/v1/columns/1", nil)

			server.router.ServeHTTP(w, r)
			var column model.Column
			err := json.NewDecoder(w.Body).Decode(&column)

			assert.NoError(t, err)
			assert.Equal(t, tc.expCode, w.Code)
			assert.Equal(t, tc.expBody, column)
		})
	}
}

func TestServer_ColumnMove(t *testing.T) {
	server := &Server{router: mux.NewRouter()}
	server.configureRouter()

	testcases := []struct {
		name    string
		mock    func(c *gomock.Controller, s *mock_service.MockService, left bool, column model.Column)
		column  model.Column
		left    bool
		expCode int
	}{
		{
			name: "column is moved right",
			mock: func(c *gomock.Controller, s *mock_service.MockService, left bool, column model.Column) {
				cs := mock_service.NewMockColumnService(c)
				cs.EXPECT().MoveByID(column.ID, left).Return(nil)
				s.EXPECT().Columns().Return(cs)
			},
			column:  model.Column{ID: 1, Name: "Column 1", Index: 1, ProjectID: 1},
			left:    false,
			expCode: http.StatusOK,
		},
		{
			name: "column is moved left",
			mock: func(c *gomock.Controller, s *mock_service.MockService, left bool, column model.Column) {
				cs := mock_service.NewMockColumnService(c)
				cs.EXPECT().MoveByID(column.ID, left).Return(nil)
				s.EXPECT().Columns().Return(cs)
			},
			column:  model.Column{ID: 1, Name: "Column 1", Index: 2, ProjectID: 1},
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
			tc.mock(c, s, tc.left, tc.column)
			server.service = s

			w := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(&request{Left: tc.left})
			r, _ := http.NewRequest(http.MethodPost, "/api/v1/columns/1/move", b)

			server.router.ServeHTTP(w, r)

			assert.Equal(t, tc.expCode, w.Code)
		})
	}
}

func TestServer_ColumnUpdate(t *testing.T) {
	server := &Server{router: mux.NewRouter()}
	server.configureRouter()

	testcases := []struct {
		name    string
		mock    func(c *gomock.Controller, s *mock_service.MockService, column model.Column)
		column  model.Column
		expCode int
		expBody model.Column
	}{
		{
			name: "column is updated",
			mock: func(c *gomock.Controller, s *mock_service.MockService, column model.Column) {
				updatedColumn := model.Column{
					ID: column.ID, Name: column.Name, Index: 1, ProjectID: 1,
				}
				cs := mock_service.NewMockColumnService(c)
				cs.EXPECT().Update(column).Return(updatedColumn, nil)
				s.EXPECT().Columns().Return(cs)
			},
			column:  model.Column{ID: 1, Name: "Updated column"},
			expCode: http.StatusOK,
			expBody: model.Column{ID: 1, Name: "Updated column", Index: 1, ProjectID: 1},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			s := mock_service.NewMockService(c)
			tc.mock(c, s, tc.column)
			server.service = s

			w := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.column)
			r, _ := http.NewRequest(http.MethodPut, "/api/v1/columns/1", b)

			server.router.ServeHTTP(w, r)
			var column model.Column
			err := json.NewDecoder(w.Body).Decode(&column)

			assert.NoError(t, err)
			assert.Equal(t, tc.expCode, w.Code)
			assert.Equal(t, tc.expBody, column)
		})
	}
}

func TestServer_ColumnDelete(t *testing.T) {
	server := &Server{router: mux.NewRouter()}
	server.configureRouter()

	testcases := []struct {
		name    string
		mock    func(c *gomock.Controller, s *mock_service.MockService, column model.Column)
		column  model.Column
		expCode int
	}{
		{
			name: "column is deleted",
			mock: func(c *gomock.Controller, s *mock_service.MockService, column model.Column) {
				cs := mock_service.NewMockColumnService(c)
				cs.EXPECT().DeleteByID(column.ID).Return(nil)
				s.EXPECT().Columns().Return(cs)
			},
			column:  model.Column{ID: 1, Name: "Column 1"},
			expCode: http.StatusNoContent,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			s := mock_service.NewMockService(c)
			tc.mock(c, s, tc.column)
			server.service = s

			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodDelete, "/api/v1/columns/1", nil)

			server.router.ServeHTTP(w, r)

			assert.Equal(t, tc.expCode, w.Code)
		})
	}
}
