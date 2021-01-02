package web

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/imarrche/tasker/internal/model"
	"github.com/imarrche/tasker/internal/store"
	mock_store "github.com/imarrche/tasker/internal/store/mocks"
)

func TestCommentService_GetByTaskID(t *testing.T) {
	testcases := []struct {
		name        string
		mock        func(*mock_store.MockStore, *gomock.Controller, model.Task)
		task        model.Task
		expComments []model.Comment
		expError    error
	}{
		{
			name: "Comments are retrieved and sorted by creating date",
			mock: func(s *mock_store.MockStore, c *gomock.Controller, t model.Task) {
				cr := mock_store.NewMockCommentRepo(c)
				tr := mock_store.NewMockTaskRepo(c)

				tr.EXPECT().GetByID(t.ID).Return(t, nil)
				cr.EXPECT().GetByTaskID(t.ID).Return(
					[]model.Comment{
						model.Comment{
							ID:        1,
							Text:      "C1",
							CreatedAt: time.Date(2020, 12, 1, 0, 0, 0, 0, &time.Location{}),
							TaskID:    t.ID,
						},
						model.Comment{
							ID:        2,
							Text:      "C2",
							CreatedAt: time.Date(2020, 12, 2, 0, 0, 0, 0, &time.Location{}),
							TaskID:    t.ID,
						},
						model.Comment{
							ID:        3,
							Text:      "C3",
							CreatedAt: time.Date(2020, 12, 3, 0, 0, 0, 0, &time.Location{}),
							TaskID:    t.ID,
						},
					},
					nil,
				)
				s.EXPECT().Comments().Return(cr)
				s.EXPECT().Tasks().Return(tr)
			},
			task: model.Task{ID: 1, Name: "T", Index: 1, ColumnID: 1},
			expComments: []model.Comment{
				model.Comment{
					ID:        3,
					Text:      "C3",
					CreatedAt: time.Date(2020, 12, 3, 0, 0, 0, 0, &time.Location{}),
					TaskID:    1,
				},
				model.Comment{
					ID:        2,
					Text:      "C2",
					CreatedAt: time.Date(2020, 12, 2, 0, 0, 0, 0, &time.Location{}),
					TaskID:    1,
				},
				model.Comment{
					ID:        1,
					Text:      "C1",
					CreatedAt: time.Date(2020, 12, 1, 0, 0, 0, 0, &time.Location{}),
					TaskID:    1,
				},
			},
			expError: nil,
		},
		{
			name: "Error occures while retrieving columns",
			mock: func(s *mock_store.MockStore, c *gomock.Controller, t model.Task) {
				cr := mock_store.NewMockCommentRepo(c)
				tr := mock_store.NewMockTaskRepo(c)

				tr.EXPECT().GetByID(t.ID).Return(t, nil)
				cr.EXPECT().GetByTaskID(t.ID).Return(nil, store.ErrDbQuery)
				s.EXPECT().Comments().Return(cr)
				s.EXPECT().Tasks().Return(tr)
			},
			task:        model.Task{ID: 1, Name: "T", Index: 1, ColumnID: 1},
			expComments: nil,
			expError:    store.ErrDbQuery,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(store, c, tc.task)
			s := newCommentService(store)

			cs, err := s.GetByTaskID(tc.task.ID)
			assert.Equal(t, tc.expError, err)
			assert.Equal(t, tc.expComments, cs)
		})
	}
}

func TestCommentService_Create(t *testing.T) {
	testcases := []struct {
		name       string
		mock       func(*mock_store.MockStore, *gomock.Controller, model.Comment)
		comment    model.Comment
		expComment model.Comment
		expError   error
	}{
		{
			name: "Comment is created",
			mock: func(s *mock_store.MockStore, c *gomock.Controller, comment model.Comment) {
				tr := mock_store.NewMockTaskRepo(c)
				cr := mock_store.NewMockCommentRepo(c)
				tr.EXPECT().GetByID(comment.TaskID).Return(model.Task{ID: 1}, nil)
				cr.EXPECT().Create(gomock.Any()).Return(
					model.Comment{
						ID:     1,
						Text:   comment.Text,
						TaskID: comment.TaskID,
					},
					nil,
				)
				s.EXPECT().Tasks().Return(tr)
				s.EXPECT().Comments().Return(cr)
			},
			comment:    model.Comment{Text: "C", TaskID: 1},
			expComment: model.Comment{ID: 1, Text: "C", TaskID: 1},
			expError:   nil,
		},
		{
			name:       "Comment doesn't pass validation",
			mock:       func(s *mock_store.MockStore, c *gomock.Controller, comment model.Comment) {},
			comment:    model.Comment{},
			expComment: model.Comment{},
			expError:   ErrTextIsRequired,
		},
		{
			name: "Invalid task ID provided",
			mock: func(s *mock_store.MockStore, c *gomock.Controller, comment model.Comment) {
				tr := mock_store.NewMockTaskRepo(c)

				tr.EXPECT().GetByID(comment.TaskID).Return(model.Task{}, store.ErrNotFound)
				s.EXPECT().Tasks().Return(tr)
			},
			comment:    model.Comment{Text: "C", TaskID: 1},
			expComment: model.Comment{},
			expError:   store.ErrNotFound,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(store, c, tc.comment)
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
		mock       func(*mock_store.MockStore, *gomock.Controller, model.Comment)
		comment    model.Comment
		expComment model.Comment
		expError   error
	}{
		{
			name: "Comment is retrieved",
			mock: func(s *mock_store.MockStore, c *gomock.Controller, comment model.Comment) {
				cr := mock_store.NewMockCommentRepo(c)

				cr.EXPECT().GetByID(comment.ID).Return(comment, nil)
				s.EXPECT().Comments().Return(cr)
			},
			comment:    model.Comment{ID: 1, Text: "C", TaskID: 1},
			expComment: model.Comment{ID: 1, Text: "C", TaskID: 1},
			expError:   nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(store, c, tc.comment)
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
		mock       func(*mock_store.MockStore, *gomock.Controller, model.Comment)
		comment    model.Comment
		expComment model.Comment
		expError   error
	}{
		{
			name: "Comment is updated",
			mock: func(s *mock_store.MockStore, c *gomock.Controller, comment model.Comment) {
				cr := mock_store.NewMockCommentRepo(c)

				cr.EXPECT().GetByID(comment.ID).Return(comment, nil)
				cr.EXPECT().Update(comment).Return(comment, nil)
				s.EXPECT().Comments().Times(2).Return(cr)
			},
			comment:    model.Comment{ID: 1, Text: "C", TaskID: 1},
			expComment: model.Comment{ID: 1, Text: "C", TaskID: 1},
			expError:   nil,
		},
		{
			name:       "Comment doesn't pass validation.",
			mock:       func(s *mock_store.MockStore, c *gomock.Controller, comment model.Comment) {},
			comment:    model.Comment{},
			expComment: model.Comment{},
			expError:   ErrTextIsRequired,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(store, c, tc.comment)
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
		mock     func(*mock_store.MockStore, *gomock.Controller, model.Comment)
		comment  model.Comment
		expError error
	}{
		{
			name: "Task is deleted",
			mock: func(s *mock_store.MockStore, c *gomock.Controller, comment model.Comment) {
				cr := mock_store.NewMockCommentRepo(c)

				cr.EXPECT().GetByID(comment.ID).Return(comment, nil)
				cr.EXPECT().DeleteByID(comment.ID).Return(nil)
				s.EXPECT().Comments().Times(2).Return(cr)
			},
			comment:  model.Comment{ID: 1, Text: "C", TaskID: 1},
			expError: nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(store, c, tc.comment)
			s := newCommentService(store)

			err := s.DeleteByID(tc.comment.ID)
			assert.Equal(t, tc.expError, err)
		})
	}
}

func TestCommentService_Validate(t *testing.T) {
	testcases := []struct {
		name     string
		mock     func(s *mock_store.MockStore, c *gomock.Controller, cooment model.Comment)
		comment  model.Comment
		expError error
	}{
		{
			name:     "Comment passes validation",
			mock:     func(s *mock_store.MockStore, c *gomock.Controller, comment model.Comment) {},
			comment:  model.Comment{Text: "C"},
			expError: nil,
		},
		{
			name:     "Comment doesn't pass validation because of empty text",
			mock:     func(s *mock_store.MockStore, c *gomock.Controller, comment model.Comment) {},
			comment:  model.Comment{},
			expError: ErrTextIsRequired,
		},
		{
			name:     "Comment doesn't pass validation bacause of too long text",
			mock:     func(s *mock_store.MockStore, c *gomock.Controller, comment model.Comment) {},
			comment:  model.Comment{Text: fixedLengthString(5001)},
			expError: ErrTextIsTooLong,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(store, c, tc.comment)
			s := newCommentService(store)

			err := s.Validate(tc.comment)
			assert.Equal(t, tc.expError, err)
		})
	}
}
