package web

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/imarrche/tasker/internal/model"
	"github.com/imarrche/tasker/internal/store"
	mock_store "github.com/imarrche/tasker/internal/store/mocks"
)

func TestTaskService_GetByColumnID(t *testing.T) {
	testcases := []struct {
		name     string
		mock     func(*mock_store.MockStore, *gomock.Controller, model.Column)
		column   model.Column
		expTasks []model.Task
		expError error
	}{
		{
			name: "Tasks are retrieved and sorted by index",
			mock: func(s *mock_store.MockStore, c *gomock.Controller, column model.Column) {
				tr := mock_store.NewMockTaskRepo(c)

				tr.EXPECT().GetByColumnID(column.ID).Return(
					[]model.Task{
						model.Task{ID: 1, Index: 3, Name: "T1", ColumnID: 1},
						model.Task{ID: 2, Index: 2, Name: "T2", ColumnID: 1},
						model.Task{ID: 3, Index: 1, Name: "T3", ColumnID: 1},
					},
					nil,
				)
				s.EXPECT().Tasks().Return(tr)
			},
			expTasks: []model.Task{
				model.Task{ID: 3, Index: 1, Name: "T3", ColumnID: 1},
				model.Task{ID: 2, Index: 2, Name: "T2", ColumnID: 1},
				model.Task{ID: 1, Index: 3, Name: "T1", ColumnID: 1},
			},
			expError: nil,
		},
		{
			name: "Error occures while retrieving tasks",
			mock: func(s *mock_store.MockStore, c *gomock.Controller, column model.Column) {
				tr := mock_store.NewMockTaskRepo(c)

				tr.EXPECT().GetByColumnID(column.ID).Return(nil, store.ErrDbQuery)
				s.EXPECT().Tasks().Return(tr)
			},
			expTasks: nil,
			expError: store.ErrDbQuery,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(store, c, tc.column)
			s := newTaskService(store)

			ts, err := s.GetByColumnID(tc.column.ID)
			assert.Equal(t, tc.expError, err)
			assert.Equal(t, tc.expTasks, ts)
		})
	}
}

func TestTaskService_Create(t *testing.T) {
	testcases := []struct {
		name     string
		mock     func(*mock_store.MockStore, *gomock.Controller, model.Task)
		task     model.Task
		expTask  model.Task
		expError error
	}{
		{
			name: "Task is created",
			mock: func(s *mock_store.MockStore, c *gomock.Controller, t model.Task) {
				cr := mock_store.NewMockColumnRepo(c)
				tr := mock_store.NewMockTaskRepo(c)

				cr.EXPECT().GetByID(t.ColumnID).Return(
					model.Column{ID: t.ColumnID, Name: "C", ProjectID: 1},
					nil,
				)
				tr.EXPECT().Create(t).Return(
					model.Task{
						ID:       1,
						Name:     t.Name,
						Index:    t.Index,
						ColumnID: t.ColumnID,
					},
					nil,
				)
				s.EXPECT().Columns().Return(cr)
				s.EXPECT().Tasks().Return(tr)
			},
			task:     model.Task{Name: "T", Index: 1, ColumnID: 1},
			expTask:  model.Task{ID: 1, Name: "T", Index: 1, ColumnID: 1},
			expError: nil,
		},
		{
			name:     "Task doesn't pass validation",
			mock:     func(s *mock_store.MockStore, c *gomock.Controller, t model.Task) {},
			task:     model.Task{},
			expTask:  model.Task{},
			expError: ErrNameIsRequired,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(store, c, tc.task)
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
		mock     func(*mock_store.MockStore, *gomock.Controller, model.Task)
		task     model.Task
		expTask  model.Task
		expError error
	}{
		{
			name: "Task is retrieved",
			mock: func(s *mock_store.MockStore, c *gomock.Controller, t model.Task) {
				tr := mock_store.NewMockTaskRepo(c)

				tr.EXPECT().GetByID(t.ID).Return(t, nil)
				s.EXPECT().Tasks().Return(tr)
			},
			task:     model.Task{ID: 1, Name: "C", Index: 1, ColumnID: 1},
			expTask:  model.Task{ID: 1, Name: "C", Index: 1, ColumnID: 1},
			expError: nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(store, c, tc.task)
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
		mock     func(*mock_store.MockStore, *gomock.Controller, model.Task)
		task     model.Task
		expTask  model.Task
		expError error
	}{
		{
			name: "Task is updated",
			mock: func(s *mock_store.MockStore, c *gomock.Controller, t model.Task) {
				tr := mock_store.NewMockTaskRepo(c)

				tr.EXPECT().Update(t).Return(t, nil)
				s.EXPECT().Tasks().Return(tr)
			},
			task:     model.Task{ID: 1, Name: "T1", Index: 1, ColumnID: 1},
			expTask:  model.Task{ID: 1, Name: "T1", Index: 1, ColumnID: 1},
			expError: nil,
		},
		{
			name:     "Column doesn't pass validation",
			mock:     func(s *mock_store.MockStore, c *gomock.Controller, t model.Task) {},
			task:     model.Task{},
			expTask:  model.Task{},
			expError: ErrNameIsRequired,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(store, c, tc.task)
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
		mock     func(*mock_store.MockStore, *gomock.Controller, model.Task)
		task     model.Task
		left     bool
		expError error
	}{
		{
			name: "Task is moved left",
			mock: func(s *mock_store.MockStore, c *gomock.Controller, task model.Task) {
				cr := mock_store.NewMockColumnRepo(c)
				tr := mock_store.NewMockTaskRepo(c)

				tr.EXPECT().GetByID(task.ID).Return(task, nil)
				cr.EXPECT().GetByID(task.ColumnID).Return(
					model.Column{ID: task.ColumnID, Index: 2, ProjectID: 1},
					nil,
				)
				cr.EXPECT().GetByIndexAndProjectID(1, 1).Return(
					model.Column{ID: 1, Index: 1},
					nil,
				)
				tr.EXPECT().GetByColumnID(1).Return(
					[]model.Task{model.Task{ID: 1, Index: 1}},
					nil,
				)
				tr.EXPECT().GetByColumnID(task.ColumnID).Return(
					[]model.Task{
						task,
						model.Task{ID: 3, Index: 2, ColumnID: 2},
					},
					nil,
				)
				tr.EXPECT().Update(model.Task{ID: 3, Index: 1, ColumnID: 2}).Return(
					model.Task{ID: 3, Index: 1, ColumnID: 2},
					nil,
				)
				tr.EXPECT().Update(model.Task{ID: 2, Index: 2, ColumnID: 1}).Return(
					model.Task{ID: 2, Index: 2, ColumnID: 1},
					nil,
				)
				s.EXPECT().Tasks().Times(5).Return(tr)
				s.EXPECT().Columns().Times(2).Return(cr)
			},
			task:     model.Task{ID: 2, Index: 1, ColumnID: 2},
			left:     true,
			expError: nil,
		},
		{
			name: "Task is moved right",
			mock: func(s *mock_store.MockStore, c *gomock.Controller, task model.Task) {
				cr := mock_store.NewMockColumnRepo(c)
				tr := mock_store.NewMockTaskRepo(c)

				tr.EXPECT().GetByID(task.ID).Return(task, nil)
				cr.EXPECT().GetByID(task.ColumnID).Return(
					model.Column{ID: task.ColumnID, Index: 1, ProjectID: 1},
					nil,
				)
				cr.EXPECT().GetByIndexAndProjectID(2, 1).Return(
					model.Column{ID: 2, Index: 2},
					nil,
				)
				tr.EXPECT().GetByColumnID(2).Return(
					[]model.Task{model.Task{ID: 1, Index: 1}},
					nil,
				)
				tr.EXPECT().GetByColumnID(task.ColumnID).Return(
					[]model.Task{
						task,
						model.Task{ID: 3, Index: 2, ColumnID: 1},
					},
					nil,
				)
				tr.EXPECT().Update(model.Task{ID: 3, Index: 1, ColumnID: 1}).Return(
					model.Task{ID: 3, Index: 1, ColumnID: 1},
					nil,
				)
				tr.EXPECT().Update(model.Task{ID: 2, Index: 2, ColumnID: 2}).Return(
					model.Task{ID: 2, Index: 2, ColumnID: 2},
					nil,
				)
				s.EXPECT().Tasks().Times(5).Return(tr)
				s.EXPECT().Columns().Times(2).Return(cr)
			},
			task:     model.Task{ID: 2, Index: 1, ColumnID: 1},
			left:     false,
			expError: nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(store, c, tc.task)
			s := newTaskService(store)

			err := s.MoveToColumnByID(tc.task.ID, tc.left)
			assert.Equal(t, tc.expError, err)
		})
	}
}

func TestTaskService_MoveByID(t *testing.T) {
	testcases := []struct {
		name     string
		mock     func(*mock_store.MockStore, *gomock.Controller, model.Task)
		task     model.Task
		up       bool
		expError error
	}{
		{
			name: "Task is moved up",
			mock: func(s *mock_store.MockStore, c *gomock.Controller, t model.Task) {
				tr := mock_store.NewMockTaskRepo(c)

				tr.EXPECT().GetByID(t.ID).Return(t, nil)
				tr.EXPECT().GetByIndexAndColumnID(t.Index-1, t.ColumnID).Return(
					model.Task{ID: 1, Index: 1, ColumnID: t.ColumnID},
					nil,
				)
				tr.EXPECT().Update(
					model.Task{ID: 1, Index: 2, ColumnID: t.ColumnID},
				).Return(
					model.Task{ID: 1, Index: 2, ColumnID: t.ColumnID},
					nil,
				)
				tr.EXPECT().Update(
					model.Task{ID: 2, Index: 1, ColumnID: t.ColumnID},
				).Return(
					model.Task{ID: 2, Index: 1, ColumnID: t.ColumnID},
					nil,
				)
				s.EXPECT().Tasks().Times(4).Return(tr)
			},
			task:     model.Task{ID: 2, Index: 2, ColumnID: 1},
			up:       true,
			expError: nil,
		},
		{
			name: "Task is moved down",
			mock: func(s *mock_store.MockStore, c *gomock.Controller, t model.Task) {
				tr := mock_store.NewMockTaskRepo(c)

				tr.EXPECT().GetByID(t.ID).Return(t, nil)
				tr.EXPECT().GetByIndexAndColumnID(t.Index+1, t.ColumnID).Return(
					model.Task{ID: 2, Index: 2, ColumnID: t.ColumnID},
					nil,
				)
				tr.EXPECT().Update(
					model.Task{ID: 2, Index: 1, ColumnID: t.ColumnID},
				).Return(
					model.Task{ID: 2, Index: 1, ColumnID: t.ColumnID},
					nil,
				)
				tr.EXPECT().Update(
					model.Task{ID: 1, Index: 2, ColumnID: t.ColumnID},
				).Return(
					model.Task{ID: 1, Index: 2, ColumnID: t.ColumnID},
					nil,
				)
				s.EXPECT().Tasks().Times(4).Return(tr)
			},
			task:     model.Task{ID: 1, Index: 1, ColumnID: 1},
			up:       false,
			expError: nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(store, c, tc.task)
			s := newTaskService(store)

			err := s.MoveByID(tc.task.ID, tc.up)
			assert.Equal(t, tc.expError, err)
		})
	}
}

func TestTaskService_DeleteByID(t *testing.T) {
	testcases := []struct {
		name     string
		mock     func(*mock_store.MockStore, *gomock.Controller, model.Task)
		task     model.Task
		expError error
	}{
		{
			name: "Task is deleted",
			mock: func(s *mock_store.MockStore, c *gomock.Controller, t model.Task) {
				tr := mock_store.NewMockTaskRepo(c)

				tr.EXPECT().GetByID(t.ID).Return(t, nil)
				tr.EXPECT().GetByColumnID(t.ColumnID).Return(
					[]model.Task{t, model.Task{Index: 2}},
					nil,
				)
				tr.EXPECT().Update(model.Task{Index: 1}).Return(model.Task{Index: 1}, nil)
				tr.EXPECT().DeleteByID(t.ID).Return(nil)
				s.EXPECT().Tasks().Times(4).Return(tr)
			},
			task:     model.Task{ID: 1, Name: "T", Index: 1, ColumnID: 1},
			expError: nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(store, c, tc.task)
			s := newTaskService(store)

			err := s.DeleteByID(tc.task.ID)
			assert.Equal(t, tc.expError, err)
		})
	}
}

func TestTaskService_Validate(t *testing.T) {
	testcases := []struct {
		name     string
		mock     func(*mock_store.MockStore, *gomock.Controller, model.Task)
		task     model.Task
		expError error
	}{
		{
			name:     "Task passes validation",
			mock:     func(s *mock_store.MockStore, c *gomock.Controller, t model.Task) {},
			task:     model.Task{Name: "T", Index: 1, ColumnID: 1},
			expError: nil,
		},
		{
			name:     "Task doesn't pass validation because of empty name",
			mock:     func(s *mock_store.MockStore, c *gomock.Controller, t model.Task) {},
			task:     model.Task{},
			expError: ErrNameIsRequired,
		},
		{
			name:     "Task doesn't pass validation because of too long name",
			mock:     func(s *mock_store.MockStore, c *gomock.Controller, t model.Task) {},
			task:     model.Task{Name: fixedLengthString(501)},
			expError: ErrNameIsTooLong,
		},
		{
			name:     "Task doesn't pass validation because of too long description",
			mock:     func(s *mock_store.MockStore, c *gomock.Controller, t model.Task) {},
			task:     model.Task{Name: "T", Description: fixedLengthString(5001)},
			expError: ErrDescriptionIsTooLong,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(store, c, tc.task)
			s := newTaskService(store)

			err := s.Validate(tc.task)
			assert.Equal(t, tc.expError, err)
		})
	}
}
