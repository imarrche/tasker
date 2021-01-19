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

func TestServer_CommentList(t *testing.T) {
	server := &Server{router: mux.NewRouter()}
	server.configureRouter()

	testcases := []struct {
		name     string
		mock     func(*gomock.Controller, *mock_service.MockService, int, []model.Comment)
		taskID   int
		comments []model.Comment
		expCode  int
		expBody  []model.Comment
	}{
		{
			name: "comment list is retrieved",
			mock: func(c *gomock.Controller, s *mock_service.MockService, tID int, comments []model.Comment) {
				ts := mock_service.NewMockCommentService(c)
				ts.EXPECT().GetByTaskID(tID).Return(comments, nil)
				s.EXPECT().Comments().Return(ts)
			},
			taskID: 1,
			comments: []model.Comment{
				{ID: 1, Text: "Comment 1"}, {ID: 2, Text: "Comment 2"},
			},
			expCode: http.StatusOK,
			expBody: []model.Comment{
				{ID: 1, Text: "Comment 1"}, {ID: 2, Text: "Comment 2"},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			s := mock_service.NewMockService(c)
			tc.mock(c, s, tc.taskID, tc.comments)
			server.service = s

			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodGet, "/api/v1/tasks/1/comments", nil)

			server.router.ServeHTTP(w, r)
			var cs []model.Comment
			err := json.NewDecoder(w.Body).Decode(&cs)

			assert.NoError(t, err)
			assert.Equal(t, tc.expCode, w.Code)
			assert.Equal(t, len(tc.expBody), len(cs))
		})
	}
}

func TestServer_CommentCreate(t *testing.T) {
	server := &Server{router: mux.NewRouter()}
	server.configureRouter()

	testcases := []struct {
		name    string
		mock    func(*gomock.Controller, *mock_service.MockService, int, model.Comment)
		taskID  int
		comment model.Comment
		expCode int
		expBody model.Comment
	}{
		{
			name: "comment is created",
			mock: func(c *gomock.Controller, s *mock_service.MockService, tID int, comment model.Comment) {
				createdComment := model.Comment{ID: 1, Text: comment.Text, TaskID: comment.TaskID}
				ts := mock_service.NewMockCommentService(c)
				ts.EXPECT().Create(comment).Return(createdComment, nil)
				s.EXPECT().Comments().Return(ts)
			},
			taskID:  1,
			comment: model.Comment{Text: "Comment", TaskID: 1},
			expCode: http.StatusCreated,
			expBody: model.Comment{ID: 1, Text: "Comment", TaskID: 1},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			s := mock_service.NewMockService(c)
			tc.mock(c, s, tc.taskID, tc.comment)
			server.service = s

			w := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.comment)
			r, _ := http.NewRequest(http.MethodPost, "/api/v1/tasks/1/comments", b)

			server.router.ServeHTTP(w, r)
			var comment model.Comment
			err := json.NewDecoder(w.Body).Decode(&comment)

			assert.NoError(t, err)
			assert.Equal(t, tc.expCode, w.Code)
			assert.Equal(t, tc.expBody, comment)
		})
	}
}

func TestServer_CommentDetail(t *testing.T) {
	server := &Server{router: mux.NewRouter()}
	server.configureRouter()

	testcases := []struct {
		name    string
		mock    func(*gomock.Controller, *mock_service.MockService, model.Comment)
		comment model.Comment
		expCode int
		expBody model.Comment
	}{
		{
			name: "comment is retrieved",
			mock: func(c *gomock.Controller, s *mock_service.MockService, comment model.Comment) {
				ts := mock_service.NewMockCommentService(c)
				ts.EXPECT().GetByID(comment.ID).Return(comment, nil)
				s.EXPECT().Comments().Return(ts)
			},
			comment: model.Comment{ID: 1, Text: "Comment 1"},
			expCode: http.StatusOK,
			expBody: model.Comment{ID: 1, Text: "Comment 1"},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			s := mock_service.NewMockService(c)
			tc.mock(c, s, tc.comment)
			server.service = s

			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodGet, "/api/v1/comments/1", nil)

			server.router.ServeHTTP(w, r)
			var comment model.Comment
			err := json.NewDecoder(w.Body).Decode(&comment)

			assert.NoError(t, err)
			assert.Equal(t, tc.expCode, w.Code)
			assert.Equal(t, tc.expBody, comment)
		})
	}
}

func TestServer_CommentUpdate(t *testing.T) {
	server := &Server{router: mux.NewRouter()}
	server.configureRouter()

	testcases := []struct {
		name    string
		mock    func(*gomock.Controller, *mock_service.MockService, model.Comment)
		comment model.Comment
		expCode int
		expBody model.Comment
	}{
		{
			name: "comment is updated",
			mock: func(c *gomock.Controller, s *mock_service.MockService, comment model.Comment) {
				ts := mock_service.NewMockCommentService(c)
				ts.EXPECT().Update(comment).Return(comment, nil)
				s.EXPECT().Comments().Return(ts)
			},
			comment: model.Comment{ID: 1, Text: "Updated comment"},
			expCode: http.StatusOK,
			expBody: model.Comment{ID: 1, Text: "Updated comment"},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			s := mock_service.NewMockService(c)
			tc.mock(c, s, tc.comment)
			server.service = s

			w := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.comment)
			r, _ := http.NewRequest(http.MethodPut, "/api/v1/comments/1", b)

			server.router.ServeHTTP(w, r)
			var comment model.Comment
			err := json.NewDecoder(w.Body).Decode(&comment)

			assert.NoError(t, err)
			assert.Equal(t, tc.expCode, w.Code)
			assert.Equal(t, tc.expBody, comment)
		})
	}
}

func TestServer_CommentDelete(t *testing.T) {
	server := &Server{router: mux.NewRouter()}
	server.configureRouter()

	testcases := []struct {
		name    string
		mock    func(*gomock.Controller, *mock_service.MockService, model.Comment)
		comment model.Comment
		expCode int
	}{
		{
			name: "comment is deleted",
			mock: func(c *gomock.Controller, s *mock_service.MockService, comment model.Comment) {
				ts := mock_service.NewMockCommentService(c)
				ts.EXPECT().DeleteByID(comment.ID).Return(nil)
				s.EXPECT().Comments().Return(ts)
			},
			comment: model.Comment{ID: 1, Text: "Comment 1"},
			expCode: http.StatusNoContent,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			s := mock_service.NewMockService(c)
			tc.mock(c, s, tc.comment)
			server.service = s

			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodDelete, "/api/v1/comments/1", nil)

			server.router.ServeHTTP(w, r)

			assert.Equal(t, tc.expCode, w.Code)
		})
	}
}
