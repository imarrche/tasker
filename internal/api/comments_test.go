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
			r, _ := http.NewRequest(http.MethodGet, "tasks/task_id/comments", nil)
			r = mux.SetURLVars(r, map[string]string{
				"task_id": "1",
			})

			s.commentList().ServeHTTP(w, r)
			var cs []model.Comment
			err := json.NewDecoder(w.Body).Decode(&cs)

			assert.NoError(t, err)
			assert.Equal(t, tc.expCode, w.Code)
			assert.Equal(t, tc.expBody[0].ID, cs[0].ID)
			assert.Equal(t, tc.expBody[1].ID, cs[1].ID)
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
			name:    "Ok, comment list is created",
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
			r, _ := http.NewRequest(http.MethodPost, "tasks/task_id/comments", b)
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
