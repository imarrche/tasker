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

func TestServer_CommentList(t *testing.T) {
	s := NewServer(inmem.TestStoreWithFixtures())

	testcases := []struct {
		name    string
		expCode int
		expBody []model.Comment
	}{
		{
			name:    "Ok, comment list is retrieved",
			expCode: http.StatusOK,
			expBody: []model.Comment{
				model.Comment{ID: 1}, model.Comment{ID: 2},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodGet, "/api/v1/tasks/task_id/comments", nil)
			r = mux.SetURLVars(r, map[string]string{
				"task_id": "1",
			})

			s.commentList().ServeHTTP(w, r)
			var cs []model.Comment
			err := json.NewDecoder(w.Body).Decode(&cs)

			assert.NoError(t, err)
			assert.Equal(t, tc.expCode, w.Code)
			assert.Equal(t, len(tc.expBody), len(cs))
		})
	}
}

func TestServer_CommentCreate(t *testing.T) {
	s := NewServer(inmem.TestStoreWithFixtures())

	testcases := []struct {
		name    string
		comment model.Comment
		expCode int
		expBody model.Comment
	}{
		{
			name:    "Ok, comment is created",
			comment: model.Comment{Text: "Comment"},
			expCode: http.StatusCreated,
			expBody: model.Comment{Text: "Comment"},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.comment)
			r, _ := http.NewRequest(http.MethodPost, "/api/v1/tasks/task_id/comments", b)
			r = mux.SetURLVars(r, map[string]string{
				"task_id": "1",
			})

			s.commentCreate().ServeHTTP(w, r)
			var c model.Comment
			err := json.NewDecoder(w.Body).Decode(&c)

			assert.NoError(t, err)
			assert.Equal(t, tc.expCode, w.Code)
			assert.Equal(t, tc.expBody.Text, c.Text)
		})
	}
}

func TestServer_CommentDetail(t *testing.T) {
	s := NewServer(inmem.TestStoreWithFixtures())

	testcases := []struct {
		name    string
		expCode int
		expBody model.Comment
	}{
		{
			name:    "Ok, comment is retrieved",
			expCode: http.StatusOK,
			expBody: model.Comment{Text: "Comment 1"},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodPost, "/api/v1/comments/comment_id", nil)
			r = mux.SetURLVars(r, map[string]string{
				"comment_id": "1",
			})

			s.commentDetail().ServeHTTP(w, r)
			var c model.Comment
			err := json.NewDecoder(w.Body).Decode(&c)

			assert.NoError(t, err)
			assert.Equal(t, tc.expCode, w.Code)
			assert.Equal(t, tc.expBody.Text, c.Text)
		})
	}
}

func TestServer_CommentUpdate(t *testing.T) {
	s := NewServer(inmem.TestStoreWithFixtures())

	testcases := []struct {
		name    string
		comment model.Comment
		expCode int
		expBody model.Comment
	}{
		{
			name:    "Ok, comment is updated",
			comment: model.Comment{Text: "Updated comment"},
			expCode: http.StatusOK,
			expBody: model.Comment{Text: "Updated comment"},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.comment)
			r, _ := http.NewRequest(http.MethodPut, "/api/v1/comments/comment_id", b)
			r = mux.SetURLVars(r, map[string]string{
				"comment_id": "1",
			})

			s.commentUpdate().ServeHTTP(w, r)
			var c model.Comment
			err := json.NewDecoder(w.Body).Decode(&c)

			assert.NoError(t, err)
			assert.Equal(t, tc.expCode, w.Code)
			assert.Equal(t, tc.expBody.Text, c.Text)
		})
	}
}

func TestServer_CommentDelete(t *testing.T) {
	s := NewServer(inmem.TestStoreWithFixtures())

	testcases := []struct {
		name    string
		expCode int
	}{
		{
			name:    "Ok, comment is deleted",
			expCode: http.StatusNoContent,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodDelete, "/api/v1/comments/comment_id", nil)
			r = mux.SetURLVars(r, map[string]string{
				"comment_id": "1",
			})

			s.commentDelete().ServeHTTP(w, r)

			assert.Equal(t, tc.expCode, w.Code)
		})
	}
}
