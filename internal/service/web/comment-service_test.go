package web

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/imarrche/tasker/internal/model"
	mock_store "github.com/imarrche/tasker/internal/store/mocks"
)

func TestCommentService_GetAll(t *testing.T) {
	testcases := [...]struct {
		name             string
		mock             func(*mock_store.MockCommentRepository)
		expectedComments []model.Comment
		expectedError    error
	}{
		{
			name: "Comments are retrieved and sorted by creating date",
			mock: func(cr *mock_store.MockCommentRepository) {
				cr.EXPECT().GetAll().Return(
					[]model.Comment{
						model.Comment{CreatedAt: time.Date(2020, 12, 1, 0, 0, 0, 0, &time.Location{})},
						model.Comment{CreatedAt: time.Date(2020, 12, 2, 0, 0, 0, 0, &time.Location{})},
						model.Comment{CreatedAt: time.Date(2020, 12, 3, 0, 0, 0, 0, &time.Location{})},
					},
					nil,
				)
			},
			expectedComments: []model.Comment{
				model.Comment{CreatedAt: time.Date(2020, 12, 3, 0, 0, 0, 0, &time.Location{})},
				model.Comment{CreatedAt: time.Date(2020, 12, 2, 0, 0, 0, 0, &time.Location{})},
				model.Comment{CreatedAt: time.Date(2020, 12, 1, 0, 0, 0, 0, &time.Location{})},
			},
			expectedError: nil,
		},
		{
			name: "Error occured while retrieving columns",
			mock: func(cr *mock_store.MockCommentRepository) {
				cr.EXPECT().GetAll().Return(nil, errors.New("couldn't get comments"))
			},
			expectedComments: nil,
			expectedError:    errors.New("couldn't get comments"),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			cr := mock_store.NewMockCommentRepository(c)
			tc.mock(cr)
			s := NewCommentService(nil, cr)

			cs, err := s.GetAll()
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedComments, cs)
		})
	}
}

func TestCommentService_Create(t *testing.T) {
	testcases := [...]struct {
		name string
		mock func(
			*mock_store.MockTaskRepository,
			*mock_store.MockCommentRepository,
			model.Comment,
		)
		comment         model.Comment
		expectedComment model.Comment
		expectedError   error
	}{
		{
			name: "Comment is created.",
			mock: func(
				tr *mock_store.MockTaskRepository,
				cr *mock_store.MockCommentRepository,
				c model.Comment,
			) {
				tr.EXPECT().GetByID(c.Task.ID).Return(c.Task, nil)
				cr.EXPECT().Create(c).Return(c, nil)
			},
			comment:         model.Comment{Text: "C1", Task: model.Task{ID: 1}},
			expectedComment: model.Comment{Text: "C1", Task: model.Task{ID: 1}},
			expectedError:   nil,
		},
		{
			name: "Comment didn't pass validation.",
			mock: func(
				tr *mock_store.MockTaskRepository,
				cr *mock_store.MockCommentRepository,
				c model.Comment,
			) {
			},
			comment:         model.Comment{},
			expectedComment: model.Comment{},
			expectedError:   ErrTextIsRequired,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			tr := mock_store.NewMockTaskRepository(c)
			cr := mock_store.NewMockCommentRepository(c)
			tc.mock(tr, cr, tc.comment)
			s := NewCommentService(tr, cr)

			comment, err := s.Create(tc.comment)
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedComment, comment)
		})
	}
}

func TestCommentService_GetByID(t *testing.T) {
	testcases := [...]struct {
		name            string
		mock            func(*mock_store.MockCommentRepository, model.Comment)
		comment         model.Comment
		expectedComment model.Comment
		expectedError   error
	}{
		{
			name: "Comment is retrieved by ID.",
			mock: func(cr *mock_store.MockCommentRepository, c model.Comment) {
				cr.EXPECT().GetByID(c.ID).Return(c, nil)
			},
			comment:         model.Comment{ID: 1},
			expectedComment: model.Comment{ID: 1},
			expectedError:   nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			cr := mock_store.NewMockCommentRepository(c)
			tc.mock(cr, tc.comment)
			s := NewCommentService(nil, cr)

			comment, err := s.GetByID(tc.comment.ID)
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedComment, comment)
		})
	}
}

func TestCommentService_Update(t *testing.T) {
	testcases := [...]struct {
		name string
		mock func(
			*mock_store.MockTaskRepository,
			*mock_store.MockCommentRepository,
			model.Comment,
		)
		comment         model.Comment
		expectedComment model.Comment
		expectedError   error
	}{
		{
			name: "Comment is updated.",
			mock: func(
				tr *mock_store.MockTaskRepository,
				cr *mock_store.MockCommentRepository,
				c model.Comment,
			) {
				tr.EXPECT().GetByID(c.Task.ID).Return(c.Task, nil)
				cr.EXPECT().Update(c).Return(nil)
			},
			comment:       model.Comment{Text: "C1", Task: model.Task{ID: 1}},
			expectedError: nil,
		},
		{
			name: "Comment didn't pass validation.",
			mock: func(
				tr *mock_store.MockTaskRepository,
				cr *mock_store.MockCommentRepository,
				c model.Comment,
			) {
			},
			comment:       model.Comment{},
			expectedError: ErrTextIsRequired,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			tr := mock_store.NewMockTaskRepository(c)
			cr := mock_store.NewMockCommentRepository(c)
			tc.mock(tr, cr, tc.comment)
			s := NewCommentService(tr, cr)

			err := s.Update(tc.comment)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestCommentService_DeleteByID(t *testing.T) {
	testcases := [...]struct {
		name          string
		mock          func(*mock_store.MockCommentRepository, model.Comment)
		comment       model.Comment
		expectedError error
	}{
		{
			name: "Task is deleted.",
			mock: func(cr *mock_store.MockCommentRepository, c model.Comment) {
				cr.EXPECT().DeleteByID(c.ID).Return(nil)
			},
			comment:       model.Comment{ID: 1},
			expectedError: nil,
		},
		{
			name: "Error occured while deleting comment.",
			mock: func(cr *mock_store.MockCommentRepository, c model.Comment) {
				cr.EXPECT().DeleteByID(c.ID).Return(errors.New("couldn't delete comment"))
			},
			comment:       model.Comment{ID: 1},
			expectedError: errors.New("couldn't delete comment"),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			cr := mock_store.NewMockCommentRepository(c)
			tc.mock(cr, tc.comment)
			s := NewCommentService(nil, cr)

			err := s.DeleteByID(tc.comment.ID)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestCommentService_Validate(t *testing.T) {
	testcases := [...]struct {
		name string
		mock func(
			*mock_store.MockTaskRepository,
			*mock_store.MockCommentRepository,
			model.Comment,
		)
		comment       model.Comment
		expectedError error
	}{
		{
			name: "Comment passes validation.",
			mock: func(
				tr *mock_store.MockTaskRepository,
				cr *mock_store.MockCommentRepository,
				c model.Comment,
			) {
				tr.EXPECT().GetByID(c.Task.ID).Return(c.Task, nil)
			},
			comment:       model.Comment{Text: "C1", Task: model.Task{ID: 1}},
			expectedError: nil,
		},
		{
			name: "Comment's text is not provided.",
			mock: func(
				tr *mock_store.MockTaskRepository,
				cr *mock_store.MockCommentRepository,
				c model.Comment,
			) {
			},
			comment:       model.Comment{},
			expectedError: ErrTextIsRequired,
		},
		{
			name: "Comment's text is too long.",
			mock: func(
				tr *mock_store.MockTaskRepository,
				cr *mock_store.MockCommentRepository,
				c model.Comment,
			) {
			},
			comment:       model.Comment{Text: fixedLengthString(5001)},
			expectedError: ErrTextIsTooLong,
		},
		{
			name: "Comment's task is invalid.",
			mock: func(
				tr *mock_store.MockTaskRepository,
				cr *mock_store.MockCommentRepository,
				c model.Comment,
			) {
				tr.EXPECT().GetByID(c.Task.ID).Return(model.Task{}, ErrInvalidTask)
			},
			comment:       model.Comment{Text: "C1", Task: model.Task{ID: 1}},
			expectedError: ErrInvalidTask,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			tr := mock_store.NewMockTaskRepository(c)
			cr := mock_store.NewMockCommentRepository(c)
			tc.mock(tr, cr, tc.comment)
			s := NewCommentService(tr, cr)

			err := s.Validate(tc.comment)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
