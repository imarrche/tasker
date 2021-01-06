package web

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/imarrche/tasker/internal/model"
	mock_store "github.com/imarrche/tasker/internal/store/mocks"
)

func TestCommentService_GetByTaskID(t *testing.T) {
	testcases := []struct {
		name        string
		mock        func(*gomock.Controller, *mock_store.MockStore, int, []model.Comment)
		taskID      int
		comments    []model.Comment
		expComments []model.Comment
		expError    error
	}{
		{
			name: "comments are retrieved and sorted by creation time from newest to oldest",
			mock: func(c *gomock.Controller, s *mock_store.MockStore, id int, cs []model.Comment) {
				cr := mock_store.NewMockCommentRepo(c)

				cr.EXPECT().GetByTaskID(id).Return(cs, nil)
				s.EXPECT().Comments().Return(cr)
			},
			taskID: 1,
			comments: []model.Comment{
				{
					ID:        1,
					Text:      "C1",
					CreatedAt: time.Date(2020, 12, 1, 0, 0, 0, 0, &time.Location{}),
					TaskID:    1,
				},
				{
					ID:        2,
					Text:      "C2",
					CreatedAt: time.Date(2020, 12, 2, 0, 0, 0, 0, &time.Location{}),
					TaskID:    1,
				},
				{
					ID:        3,
					Text:      "C3",
					CreatedAt: time.Date(2020, 12, 3, 0, 0, 0, 0, &time.Location{}),
					TaskID:    1,
				},
			},
			expComments: []model.Comment{
				{
					ID:        3,
					Text:      "C3",
					CreatedAt: time.Date(2020, 12, 3, 0, 0, 0, 0, &time.Location{}),
					TaskID:    1,
				},
				{
					ID:        2,
					Text:      "C2",
					CreatedAt: time.Date(2020, 12, 2, 0, 0, 0, 0, &time.Location{}),
					TaskID:    1,
				},
				{
					ID:        1,
					Text:      "C1",
					CreatedAt: time.Date(2020, 12, 1, 0, 0, 0, 0, &time.Location{}),
					TaskID:    1,
				},
			},
			expError: nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(c, store, tc.taskID, tc.comments)
			s := newCommentService(store)
			cs, err := s.GetByTaskID(tc.taskID)

			assert.Equal(t, tc.expError, err)
			assert.Equal(t, tc.expComments, cs)
		})
	}
}

func TestCommentService_Create(t *testing.T) {
	testcases := []struct {
		name       string
		mock       func(*gomock.Controller, *mock_store.MockStore, model.Comment)
		comment    model.Comment
		expComment model.Comment
		expError   error
	}{
		{
			name: "comment is created",
			mock: func(c *gomock.Controller, s *mock_store.MockStore, comment model.Comment) {
				cr := mock_store.NewMockCommentRepo(c)

				cr.EXPECT().Create(gomock.Any()).Return(
					model.Comment{ID: 1, Text: comment.Text, TaskID: comment.TaskID},
					nil,
				)
				s.EXPECT().Comments().Return(cr)
			},
			comment:    model.Comment{Text: "Comment 1", TaskID: 1},
			expComment: model.Comment{ID: 1, Text: "Comment 1", TaskID: 1},
			expError:   nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(c, store, tc.comment)
			s := newCommentService(store)
			comment, err := s.Create(tc.comment)

			assert.Equal(t, tc.expError, err)
			assert.Equal(t, tc.expComment, comment)
		})
	}
}

func TestCommentService_GetByID(t *testing.T) {
	testcases := []struct {
		name       string
		mock       func(*gomock.Controller, *mock_store.MockStore, model.Comment)
		comment    model.Comment
		expComment model.Comment
		expError   error
	}{
		{
			name: "comment is retrieved",
			mock: func(c *gomock.Controller, s *mock_store.MockStore, comment model.Comment) {
				cr := mock_store.NewMockCommentRepo(c)

				cr.EXPECT().GetByID(comment.ID).Return(comment, nil)
				s.EXPECT().Comments().Return(cr)
			},
			comment:    model.Comment{ID: 1, Text: "Comment 1", CreatedAt: time.Time{}, TaskID: 1},
			expComment: model.Comment{ID: 1, Text: "Comment 1", CreatedAt: time.Time{}, TaskID: 1},
			expError:   nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(c, store, tc.comment)
			s := newCommentService(store)
			comment, err := s.GetByID(tc.comment.ID)

			assert.Equal(t, tc.expError, err)
			assert.Equal(t, tc.expComment, comment)
		})
	}
}

func TestCommentService_Update(t *testing.T) {
	testcases := []struct {
		name       string
		mock       func(*gomock.Controller, *mock_store.MockStore, model.Comment)
		comment    model.Comment
		expComment model.Comment
		expError   error
	}{
		{
			name: "comment is updated",
			mock: func(c *gomock.Controller, s *mock_store.MockStore, comment model.Comment) {
				cr := mock_store.NewMockCommentRepo(c)

				cr.EXPECT().GetByID(comment.ID).Return(comment, nil)
				cr.EXPECT().Update(comment).Return(comment, nil)
				s.EXPECT().Comments().Times(2).Return(cr)
			},
			comment:    model.Comment{ID: 1, Text: "Comment 1", CreatedAt: time.Time{}, TaskID: 1},
			expComment: model.Comment{ID: 1, Text: "Comment 1", CreatedAt: time.Time{}, TaskID: 1},
			expError:   nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(c, store, tc.comment)
			s := newCommentService(store)
			comment, err := s.Update(tc.comment)

			assert.Equal(t, tc.expComment, comment)
			assert.Equal(t, tc.expError, err)
		})
	}
}

func TestCommentService_DeleteByID(t *testing.T) {
	testcases := []struct {
		name     string
		mock     func(*gomock.Controller, *mock_store.MockStore, model.Comment)
		comment  model.Comment
		expError error
	}{
		{
			name: "comment is deleted",
			mock: func(c *gomock.Controller, s *mock_store.MockStore, comment model.Comment) {
				cr := mock_store.NewMockCommentRepo(c)

				cr.EXPECT().DeleteByID(comment.ID).Return(nil)
				s.EXPECT().Comments().Return(cr)
			},
			comment:  model.Comment{ID: 1, Text: "Comment 1", CreatedAt: time.Time{}, TaskID: 1},
			expError: nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(c, store, tc.comment)
			s := newCommentService(store)
			err := s.DeleteByID(tc.comment.ID)

			assert.Equal(t, tc.expError, err)
		})
	}
}

func TestCommentService_Validate(t *testing.T) {
	testcases := []struct {
		name     string
		mock     func(c *gomock.Controller, s *mock_store.MockStore, cooment model.Comment)
		comment  model.Comment
		expError error
	}{
		{
			name:     "comment passes validation",
			mock:     func(c *gomock.Controller, s *mock_store.MockStore, comment model.Comment) {},
			comment:  model.Comment{Text: "Comment 1"},
			expError: nil,
		},
		{
			name:     "comment doesn't pass validation because of empty text",
			mock:     func(c *gomock.Controller, s *mock_store.MockStore, comment model.Comment) {},
			comment:  model.Comment{},
			expError: ErrTextIsRequired,
		},
		{
			name:     "comment doesn't pass validation bacause of too long text",
			mock:     func(c *gomock.Controller, s *mock_store.MockStore, comment model.Comment) {},
			comment:  model.Comment{Text: fixedLengthString(5001)},
			expError: ErrTextIsTooLong,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(c, store, tc.comment)
			s := newCommentService(store)
			err := s.Validate(tc.comment)

			assert.Equal(t, tc.expError, err)
		})
	}
}
