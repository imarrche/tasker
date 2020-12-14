package web

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/imarrche/tasker/internal/model"
	mock_store "github.com/imarrche/tasker/internal/store/mocks"
)

func TestTaskService_GetAll(t *testing.T) {
	testcases := [...]struct {
		name          string
		mock          func(tr *mock_store.MockTaskRepository)
		expectedTasks []model.Task
		expectedError error
	}{
		{
			name: "Task are retrieved and sorted by index.",
			mock: func(tr *mock_store.MockTaskRepository) {
				tr.EXPECT().GetAll().Return(
					[]model.Task{model.Task{Index: 3}, model.Task{Index: 2}, model.Task{Index: 1}},
					nil,
				)
			},
			expectedTasks: []model.Task{
				model.Task{Index: 1},
				model.Task{Index: 2},
				model.Task{Index: 3},
			},
			expectedError: nil,
		},
		{
			name: "Error occured while retrieving tasks.",
			mock: func(tr *mock_store.MockTaskRepository) {
				tr.EXPECT().GetAll().Return(nil, errors.New("couldn't get tasks"))
			},
			expectedTasks: nil,
			expectedError: errors.New("couldn't get tasks"),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			tr := mock_store.NewMockTaskRepository(c)
			tc.mock(tr)
			s := NewTaskService(nil, tr)

			ts, err := s.GetAll()
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedTasks, ts)
		})
	}
}

func TestTaskService_Create(t *testing.T) {
	testcases := [...]struct {
		name string
		mock func(
			*mock_store.MockColumnRepository,
			*mock_store.MockTaskRepository,
			model.Task,
		)
		task          model.Task
		expectedTask  model.Task
		expectedError error
	}{
		{
			name: "Task is created.",
			mock: func(
				cr *mock_store.MockColumnRepository,
				tr *mock_store.MockTaskRepository,
				t model.Task,
			) {
				tr.EXPECT().Create(t).Return(t, nil)
				cr.EXPECT().GetByID(t.Column.ID).Return(model.Column{ID: t.Column.ID}, nil)
			},
			task:          model.Task{Name: "T1", Column: model.Column{ID: 1}},
			expectedTask:  model.Task{Name: "T1", Column: model.Column{ID: 1}},
			expectedError: nil,
		},
		{
			name: "Task didn't pass validation.",
			mock: func(
				cr *mock_store.MockColumnRepository,
				tr *mock_store.MockTaskRepository,
				t model.Task,
			) {
			},
			task:          model.Task{},
			expectedTask:  model.Task{},
			expectedError: ErrNameIsRequired,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			cr := mock_store.NewMockColumnRepository(c)
			tr := mock_store.NewMockTaskRepository(c)
			tc.mock(cr, tr, tc.task)
			s := NewTaskService(cr, tr)

			task, err := s.Create(tc.task)
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedTask, task)
		})
	}
}

func TestTaskService_GetByID(t *testing.T) {
	testcases := [...]struct {
		name          string
		mock          func(*mock_store.MockTaskRepository, model.Task)
		task          model.Task
		expectedTask  model.Task
		expectedError error
	}{
		{
			name: "Task is retrieved by ID.",
			mock: func(tr *mock_store.MockTaskRepository, t model.Task) {
				tr.EXPECT().GetByID(t.ID).Return(t, nil)
			},
			task:          model.Task{ID: 1},
			expectedTask:  model.Task{ID: 1},
			expectedError: nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			tr := mock_store.NewMockTaskRepository(c)
			tc.mock(tr, tc.task)
			s := NewTaskService(nil, tr)

			task, err := s.GetByID(tc.task.ID)
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedTask, task)
		})
	}
}

func TestTaskService_Update(t *testing.T) {
	testcases := [...]struct {
		name string
		mock func(
			*mock_store.MockColumnRepository,
			*mock_store.MockTaskRepository,
			model.Task,
		)
		task          model.Task
		expectedTask  model.Task
		expectedError error
	}{
		{
			name: "Task is updated.",
			mock: func(
				cr *mock_store.MockColumnRepository,
				tr *mock_store.MockTaskRepository,
				t model.Task,
			) {
				tr.EXPECT().Update(t).Return(nil)
				cr.EXPECT().GetByID(t.Column.ID).Return(model.Column{ID: t.Column.ID}, nil)
			},
			task:          model.Task{Name: "T1", Column: model.Column{ID: 1}},
			expectedError: nil,
		},
		{
			name: "Column didn't pass validation.",
			mock: func(
				cr *mock_store.MockColumnRepository,
				tr *mock_store.MockTaskRepository,
				c model.Task,
			) {
			},
			task:          model.Task{},
			expectedError: ErrNameIsRequired,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			cr := mock_store.NewMockColumnRepository(c)
			tr := mock_store.NewMockTaskRepository(c)
			tc.mock(cr, tr, tc.task)
			s := NewTaskService(cr, tr)

			err := s.Update(tc.task)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestTaskService_DeleteByID(t *testing.T) {
	testcases := [...]struct {
		name          string
		mock          func(*mock_store.MockTaskRepository, model.Task)
		task          model.Task
		expectedError error
	}{
		{
			name: "Task is deleted.",
			mock: func(tr *mock_store.MockTaskRepository, t model.Task) {
				tr.EXPECT().DeleteByID(t.ID).Return(nil)
			},
			task:          model.Task{ID: 1},
			expectedError: nil,
		},
		{
			name: "Error occured while deleting task.",
			mock: func(tr *mock_store.MockTaskRepository, t model.Task) {
				tr.EXPECT().DeleteByID(t.ID).Return(errors.New("couldn't delete task"))
			},
			task:          model.Task{ID: 1},
			expectedError: errors.New("couldn't delete task"),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			tr := mock_store.NewMockTaskRepository(c)
			tc.mock(tr, tc.task)
			s := NewTaskService(nil, tr)

			err := s.DeleteByID(tc.task.ID)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestTaskService_Validate(t *testing.T) {
	testcases := [...]struct {
		name          string
		mock          func(*mock_store.MockColumnRepository, model.Task)
		task          model.Task
		expectedError error
	}{
		{
			name: "Task passes validation.",
			mock: func(cr *mock_store.MockColumnRepository, t model.Task) {
				cr.EXPECT().GetByID(t.Column.ID).Return(t.Column, nil)
			},
			task:          model.Task{Name: "T1", Column: model.Column{ID: 1}},
			expectedError: nil,
		},
		{
			name:          "Task's name is not provided.",
			mock:          func(cr *mock_store.MockColumnRepository, t model.Task) {},
			task:          model.Task{},
			expectedError: ErrNameIsRequired,
		},
		{
			name:          "Task's name is too long.",
			mock:          func(cr *mock_store.MockColumnRepository, t model.Task) {},
			task:          model.Task{Name: fixedLengthString(501)},
			expectedError: ErrNameIsTooLong,
		},
		{
			name:          "Task's description is too long.",
			mock:          func(cr *mock_store.MockColumnRepository, t model.Task) {},
			task:          model.Task{Name: "T1", Description: fixedLengthString(5001)},
			expectedError: ErrDescriptionIsTooLong,
		},
		{
			name: "Task's column is invalid.",
			mock: func(cr *mock_store.MockColumnRepository, t model.Task) {
				cr.EXPECT().GetByID(t.Column.ID).Return(model.Column{}, ErrInvalidColumn)
			},
			task:          model.Task{Name: "T1", Column: model.Column{ID: 1}},
			expectedError: ErrInvalidColumn,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			cr := mock_store.NewMockColumnRepository(c)
			tc.mock(cr, tc.task)
			s := NewTaskService(cr, nil)

			err := s.Validate(tc.task)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
