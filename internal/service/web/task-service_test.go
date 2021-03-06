package web

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/imarrche/tasker/internal/model"
	mock_store "github.com/imarrche/tasker/internal/store/mocks"
)

func TestTaskService_GetByColumnID(t *testing.T) {
	testcases := []struct {
		name     string
		mock     func(*gomock.Controller, *mock_store.MockStore, int, []model.Task)
		columnID int
		tasks    []model.Task
		expTasks []model.Task
		expError error
	}{
		{
			name: "tasks are retrieved and sorted by index",
			mock: func(c *gomock.Controller, s *mock_store.MockStore, id int, ts []model.Task) {
				tr := mock_store.NewMockTaskRepo(c)

				tr.EXPECT().GetByColumnID(id).Return(ts, nil)
				s.EXPECT().Tasks().Return(tr)
			},
			tasks: []model.Task{
				{ID: 1, Index: 3, Name: "T1", ColumnID: 1},
				{ID: 2, Index: 2, Name: "T2", ColumnID: 1},
				{ID: 3, Index: 1, Name: "T3", ColumnID: 1},
			},
			expTasks: []model.Task{
				{ID: 3, Index: 1, Name: "T3", ColumnID: 1},
				{ID: 2, Index: 2, Name: "T2", ColumnID: 1},
				{ID: 1, Index: 3, Name: "T1", ColumnID: 1},
			},
			expError: nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(c, store, tc.columnID, tc.tasks)
			s := newTaskService(store)
			ts, err := s.GetByColumnID(tc.columnID)

			assert.Equal(t, tc.expError, err)
			assert.Equal(t, tc.expTasks, ts)
		})
	}
}

func TestTaskService_Create(t *testing.T) {
	testcases := []struct {
		name     string
		mock     func(*gomock.Controller, *mock_store.MockStore, model.Task)
		task     model.Task
		expTask  model.Task
		expError error
	}{
		{
			name: "task is created",
			mock: func(c *gomock.Controller, s *mock_store.MockStore, t model.Task) {
				tr := mock_store.NewMockTaskRepo(c)

				tr.EXPECT().GetByColumnID(t.ColumnID).Return([]model.Task{}, nil)
				tr.EXPECT().Create(t).Return(
					model.Task{ID: 1, Name: t.Name, Index: t.Index, ColumnID: t.ColumnID},
					nil,
				)
				s.EXPECT().Tasks().Times(2).Return(tr)
			},
			task:     model.Task{Name: "Task 1", Index: 1, ColumnID: 1},
			expTask:  model.Task{ID: 1, Name: "Task 1", Index: 1, ColumnID: 1},
			expError: nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(c, store, tc.task)
			s := newTaskService(store)
			task, err := s.Create(tc.task)

			assert.Equal(t, tc.expError, err)
			assert.Equal(t, tc.expTask, task)
		})
	}
}

func TestTaskService_GetByID(t *testing.T) {
	testcases := []struct {
		name     string
		mock     func(*gomock.Controller, *mock_store.MockStore, model.Task)
		task     model.Task
		expTask  model.Task
		expError error
	}{
		{
			name: "task is retrieved",
			mock: func(c *gomock.Controller, s *mock_store.MockStore, t model.Task) {
				tr := mock_store.NewMockTaskRepo(c)

				tr.EXPECT().GetByID(t.ID).Return(t, nil)
				s.EXPECT().Tasks().Return(tr)
			},
			task:     model.Task{ID: 1, Name: "Task 1", Index: 1, ColumnID: 1},
			expTask:  model.Task{ID: 1, Name: "Task 1", Index: 1, ColumnID: 1},
			expError: nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(c, store, tc.task)
			s := newTaskService(store)
			task, err := s.GetByID(tc.task.ID)

			assert.Equal(t, tc.expError, err)
			assert.Equal(t, tc.expTask, task)
		})
	}
}

func TestTaskService_Update(t *testing.T) {
	testcases := []struct {
		name     string
		mock     func(*gomock.Controller, *mock_store.MockStore, model.Task)
		task     model.Task
		expTask  model.Task
		expError error
	}{
		{
			name: "task is updated",
			mock: func(c *gomock.Controller, s *mock_store.MockStore, t model.Task) {
				tr := mock_store.NewMockTaskRepo(c)

				tr.EXPECT().GetByID(t.ID).Return(t, nil)
				tr.EXPECT().Update(t).Return(t, nil)
				s.EXPECT().Tasks().Times(2).Return(tr)
			},
			task:     model.Task{ID: 1, Name: "Task 1", Index: 1, ColumnID: 1},
			expTask:  model.Task{ID: 1, Name: "Task 1", Index: 1, ColumnID: 1},
			expError: nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(c, store, tc.task)
			s := newTaskService(store)
			task, err := s.Update(tc.task)

			assert.Equal(t, tc.expTask, task)
			assert.Equal(t, tc.expError, err)
		})
	}
}

func TestTaskService_MoveToColumnByID(t *testing.T) {
	testcases := []struct {
		name     string
		mock     func(*gomock.Controller, *mock_store.MockStore, model.Task)
		task     model.Task
		left     bool
		expError error
	}{
		{
			name: "task is moved left",
			mock: func(c *gomock.Controller, s *mock_store.MockStore, task model.Task) {
				cr := mock_store.NewMockColumnRepo(c)
				tr := mock_store.NewMockTaskRepo(c)

				tr.EXPECT().GetByID(task.ID).Return(task, nil)
				cr.EXPECT().GetByID(task.ColumnID).Return(
					model.Column{ID: task.ColumnID, Name: "Column 1", Index: 2, ProjectID: 1},
					nil,
				)
				cr.EXPECT().GetByIndexAndProjectID(1, 1).Return(
					model.Column{ID: 1, Name: "Column 1", Index: 1, ProjectID: 1},
					nil,
				)
				tr.EXPECT().GetByColumnID(1).Return(
					[]model.Task{{ID: 1, Name: "Task 1", Index: 1, ColumnID: 1}},
					nil,
				)
				tr.EXPECT().GetByColumnID(task.ColumnID).Return(
					[]model.Task{
						task, {ID: 3, Name: "Task 3", Index: 2, ColumnID: 2},
					},
					nil,
				)
				tr.EXPECT().Update(model.Task{ID: 3, Name: "Task 3", Index: 1, ColumnID: 2}).Return(
					model.Task{ID: 3, Name: "Task 3", Index: 1, ColumnID: 2},
					nil,
				)
				tr.EXPECT().Update(model.Task{ID: 2, Name: "Task 2", Index: 2, ColumnID: 1}).Return(
					model.Task{ID: 2, Name: "Task 2", Index: 2, ColumnID: 1},
					nil,
				)
				s.EXPECT().Tasks().Times(5).Return(tr)
				s.EXPECT().Columns().Times(2).Return(cr)
			},
			task:     model.Task{ID: 2, Name: "Task 2", Index: 1, ColumnID: 2},
			left:     true,
			expError: nil,
		},
		{
			name: "task is moved right",
			mock: func(c *gomock.Controller, s *mock_store.MockStore, task model.Task) {
				cr := mock_store.NewMockColumnRepo(c)
				tr := mock_store.NewMockTaskRepo(c)

				tr.EXPECT().GetByID(task.ID).Return(task, nil)
				cr.EXPECT().GetByID(task.ColumnID).Return(
					model.Column{ID: task.ColumnID, Name: "Column 1", Index: 1, ProjectID: 1},
					nil,
				)
				cr.EXPECT().GetByIndexAndProjectID(2, 1).Return(
					model.Column{ID: 2, Name: "Column 2", Index: 2, ProjectID: 1},
					nil,
				)
				tr.EXPECT().GetByColumnID(2).Return(
					[]model.Task{{ID: 1, Name: "Task 1", Index: 1, ColumnID: 2}},
					nil,
				)
				tr.EXPECT().GetByColumnID(task.ColumnID).Return(
					[]model.Task{
						task, {ID: 3, Name: "Task 3", Index: 2, ColumnID: 1},
					},
					nil,
				)
				tr.EXPECT().Update(model.Task{ID: 3, Name: "Task 3", Index: 1, ColumnID: 1}).Return(
					model.Task{ID: 3, Name: "Task 3", Index: 1, ColumnID: 1},
					nil,
				)
				tr.EXPECT().Update(model.Task{ID: 2, Name: "Task 2", Index: 2, ColumnID: 2}).Return(
					model.Task{ID: 2, Name: "Task 2", Index: 2, ColumnID: 2},
					nil,
				)
				s.EXPECT().Tasks().Times(5).Return(tr)
				s.EXPECT().Columns().Times(2).Return(cr)
			},
			task:     model.Task{ID: 2, Name: "Task 2", Index: 1, ColumnID: 1},
			left:     false,
			expError: nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(c, store, tc.task)
			s := newTaskService(store)
			err := s.MoveToColumnByID(tc.task.ID, tc.left)

			assert.Equal(t, tc.expError, err)
		})
	}
}

func TestTaskService_MoveByID(t *testing.T) {
	testcases := []struct {
		name     string
		mock     func(*gomock.Controller, *mock_store.MockStore, model.Task)
		task     model.Task
		up       bool
		expError error
	}{
		{
			name: "task is moved up",
			mock: func(c *gomock.Controller, s *mock_store.MockStore, t model.Task) {
				tr := mock_store.NewMockTaskRepo(c)

				tr.EXPECT().GetByID(t.ID).Return(t, nil)
				tr.EXPECT().GetByIndexAndColumnID(t.Index-1, t.ColumnID).Return(
					model.Task{ID: 1, Name: "Task 1", Index: 1, ColumnID: t.ColumnID},
					nil,
				)
				tr.EXPECT().Update(
					model.Task{ID: 1, Name: "Task 1", Index: 2, ColumnID: t.ColumnID},
				).Return(
					model.Task{ID: 1, Name: "Task 1", Index: 2, ColumnID: t.ColumnID},
					nil,
				)
				tr.EXPECT().Update(
					model.Task{ID: 2, Name: "Task 2", Index: 1, ColumnID: t.ColumnID},
				).Return(
					model.Task{ID: 2, Name: "Task 2", Index: 1, ColumnID: t.ColumnID},
					nil,
				)
				s.EXPECT().Tasks().Times(4).Return(tr)
			},
			task:     model.Task{ID: 2, Name: "Task 2", Index: 2, ColumnID: 1},
			up:       true,
			expError: nil,
		},
		{
			name: "task is moved down",
			mock: func(c *gomock.Controller, s *mock_store.MockStore, t model.Task) {
				tr := mock_store.NewMockTaskRepo(c)

				tr.EXPECT().GetByID(t.ID).Return(t, nil)
				tr.EXPECT().GetByIndexAndColumnID(t.Index+1, t.ColumnID).Return(
					model.Task{ID: 2, Name: "Task 2", Index: 2, ColumnID: t.ColumnID},
					nil,
				)
				tr.EXPECT().Update(
					model.Task{ID: 2, Name: "Task 2", Index: 1, ColumnID: t.ColumnID},
				).Return(
					model.Task{ID: 2, Name: "Task 2", Index: 1, ColumnID: t.ColumnID},
					nil,
				)
				tr.EXPECT().Update(
					model.Task{ID: 1, Name: "Task 1", Index: 2, ColumnID: t.ColumnID},
				).Return(
					model.Task{ID: 1, Name: "Task 1", Index: 2, ColumnID: t.ColumnID},
					nil,
				)
				s.EXPECT().Tasks().Times(4).Return(tr)
			},
			task:     model.Task{ID: 1, Name: "Task 1", Index: 1, ColumnID: 1},
			up:       false,
			expError: nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(c, store, tc.task)
			s := newTaskService(store)
			err := s.MoveByID(tc.task.ID, tc.up)

			assert.Equal(t, tc.expError, err)
		})
	}
}

func TestTaskService_DeleteByID(t *testing.T) {
	testcases := []struct {
		name     string
		mock     func(*gomock.Controller, *mock_store.MockStore, model.Task)
		task     model.Task
		expError error
	}{
		{
			name: "task is deleted",
			mock: func(c *gomock.Controller, s *mock_store.MockStore, t model.Task) {
				tr := mock_store.NewMockTaskRepo(c)

				tr.EXPECT().GetByID(t.ID).Return(t, nil)
				tr.EXPECT().GetByColumnID(t.ColumnID).Return(
					[]model.Task{
						t, {ID: 2, Name: "Task 2", Index: 2, ColumnID: t.ColumnID}},
					nil,
				)
				tr.EXPECT().Update(
					model.Task{ID: 2, Name: "Task 2", Index: 1, ColumnID: t.ColumnID},
				).Return(
					model.Task{ID: 2, Name: "Task 2", Index: 1, ColumnID: t.ColumnID},
					nil,
				)
				tr.EXPECT().DeleteByID(t.ID).Return(nil)
				s.EXPECT().Tasks().Times(4).Return(tr)
			},
			task:     model.Task{ID: 1, Name: "Task 1", Index: 1, ColumnID: 1},
			expError: nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(c, store, tc.task)
			s := newTaskService(store)
			err := s.DeleteByID(tc.task.ID)

			assert.Equal(t, tc.expError, err)
		})
	}
}

func TestTaskService_Validate(t *testing.T) {
	testcases := []struct {
		name     string
		mock     func(*gomock.Controller, *mock_store.MockStore, model.Task)
		task     model.Task
		expError error
	}{
		{
			name:     "task passes validation",
			mock:     func(c *gomock.Controller, s *mock_store.MockStore, t model.Task) {},
			task:     model.Task{Name: "Task 1", Index: 1, ColumnID: 1},
			expError: nil,
		},
		{
			name:     "task doesn't pass validation because of empty name",
			mock:     func(c *gomock.Controller, s *mock_store.MockStore, t model.Task) {},
			task:     model.Task{},
			expError: ErrNameIsRequired,
		},
		{
			name:     "task doesn't pass validation because of too long name",
			mock:     func(c *gomock.Controller, s *mock_store.MockStore, t model.Task) {},
			task:     model.Task{Name: fixedLengthString(501)},
			expError: ErrNameIsTooLong,
		},
		{
			name:     "task doesn't pass validation because of too long description",
			mock:     func(c *gomock.Controller, s *mock_store.MockStore, t model.Task) {},
			task:     model.Task{Name: "Task 1", Description: fixedLengthString(5001)},
			expError: ErrDescriptionIsTooLong,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(c, store, tc.task)
			s := newTaskService(store)
			err := s.Validate(tc.task)

			assert.Equal(t, tc.expError, err)
		})
	}
}
