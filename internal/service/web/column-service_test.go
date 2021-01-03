package web

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/imarrche/tasker/internal/model"
	"github.com/imarrche/tasker/internal/store"
	mock_store "github.com/imarrche/tasker/internal/store/mocks"
)

func TestColumnService_GetByProjectID(t *testing.T) {
	testcases := []struct {
		name       string
		mock       func(*gomock.Controller, *mock_store.MockStore, int, []model.Column)
		projectID  int
		columns    []model.Column
		expColumns []model.Column
		expError   error
	}{
		{
			name: "columns are retrieved and sorted alphabetically",
			mock: func(c *gomock.Controller, s *mock_store.MockStore, id int, columns []model.Column) {
				cr := mock_store.NewMockColumnRepo(c)

				cr.EXPECT().GetByProjectID(id).Return(columns, nil)
				s.EXPECT().Columns().Return(cr)
			},
			projectID: 1,
			columns: []model.Column{
				model.Column{ID: 1, Name: "C1", Index: 3, ProjectID: 1},
				model.Column{ID: 2, Name: "C2", Index: 2, ProjectID: 1},
				model.Column{ID: 3, Name: "C3", Index: 1, ProjectID: 1},
			},
			expColumns: []model.Column{
				model.Column{ID: 3, Name: "C3", Index: 1, ProjectID: 1},
				model.Column{ID: 2, Name: "C2", Index: 2, ProjectID: 1},
				model.Column{ID: 1, Name: "C1", Index: 3, ProjectID: 1},
			},
			expError: nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(c, store, tc.projectID, tc.columns)
			s := newColumnService(store)
			cs, err := s.GetByProjectID(tc.projectID)

			assert.Equal(t, tc.expError, err)
			assert.Equal(t, tc.expColumns, cs)
		})
	}
}

func TestColumnService_Create(t *testing.T) {
	testcases := []struct {
		name      string
		mock      func(*gomock.Controller, *mock_store.MockStore, model.Column)
		column    model.Column
		expColumn model.Column
		expError  error
	}{
		{
			name: "column is created",
			mock: func(c *gomock.Controller, s *mock_store.MockStore, column model.Column) {
				cr := mock_store.NewMockColumnRepo(c)

				cr.EXPECT().GetByProjectID(column.ProjectID).Times(2).Return([]model.Column{}, nil)
				cr.EXPECT().Create(column).Return(
					model.Column{ID: 1, Name: column.Name, Index: column.Index, ProjectID: column.ProjectID},
					nil,
				)
				s.EXPECT().Columns().Times(3).Return(cr)
			},
			column:    model.Column{Name: "Column 1", Index: 1, ProjectID: 1},
			expColumn: model.Column{ID: 1, Name: "Column 1", Index: 1, ProjectID: 1},
			expError:  nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(c, store, tc.column)
			s := newColumnService(store)
			column, err := s.Create(tc.column)

			assert.Equal(t, tc.expError, err)
			assert.Equal(t, tc.expColumn, column)
		})
	}
}

func TestColumnService_GetByID(t *testing.T) {
	testcases := []struct {
		name      string
		mock      func(*gomock.Controller, *mock_store.MockStore, model.Column)
		column    model.Column
		expColumn model.Column
		expError  error
	}{
		{
			name: "column is retrieved",
			mock: func(c *gomock.Controller, s *mock_store.MockStore, column model.Column) {
				cr := mock_store.NewMockColumnRepo(c)

				cr.EXPECT().GetByID(column.ID).Return(column, nil)
				s.EXPECT().Columns().Return(cr)
			},
			column:    model.Column{ID: 1, Name: "Column 1", Index: 1, ProjectID: 1},
			expColumn: model.Column{ID: 1, Name: "Column 1", Index: 1, ProjectID: 1},
			expError:  nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(c, store, tc.column)
			s := newColumnService(store)
			column, err := s.GetByID(tc.column.ID)

			assert.Equal(t, tc.expError, err)
			assert.Equal(t, tc.expColumn, column)
		})
	}
}

func TestColumnService_Update(t *testing.T) {
	testcases := []struct {
		name      string
		mock      func(*gomock.Controller, *mock_store.MockStore, model.Column)
		column    model.Column
		expColumn model.Column
		expError  error
	}{
		{
			name: "column is updated",
			mock: func(c *gomock.Controller, s *mock_store.MockStore, column model.Column) {
				cr := mock_store.NewMockColumnRepo(c)

				cr.EXPECT().GetByID(column.ID).Return(column, nil)
				cr.EXPECT().GetByProjectID(column.ProjectID).Return([]model.Column{}, nil)
				cr.EXPECT().Update(column).Return(column, nil)
				s.EXPECT().Columns().Times(3).Return(cr)
			},
			column:    model.Column{ID: 1, Name: "Column 1", Index: 1, ProjectID: 1},
			expColumn: model.Column{ID: 1, Name: "Column 1", Index: 1, ProjectID: 1},
			expError:  nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(c, store, tc.column)
			s := newColumnService(store)
			column, err := s.Update(tc.column)

			assert.Equal(t, tc.expError, err)
			assert.Equal(t, tc.expColumn, column)
		})
	}
}

func TestColumnService_MoveByID(t *testing.T) {
	testcases := []struct {
		name     string
		mock     func(*gomock.Controller, *mock_store.MockStore, model.Column)
		column   model.Column
		left     bool
		expError error
	}{
		{
			name: "column is moved left",
			mock: func(c *gomock.Controller, s *mock_store.MockStore, column model.Column) {
				cr := mock_store.NewMockColumnRepo(c)

				cr.EXPECT().GetByID(column.ID).Return(column, nil)
				cr.EXPECT().GetByIndexAndProjectID(column.Index-1, column.ProjectID).Return(
					model.Column{ID: 1, Index: column.Index - 1},
					nil,
				)
				cr.EXPECT().Update(model.Column{ID: 2, Index: 1}).Return(
					model.Column{ID: 2, Index: 1},
					nil,
				)
				cr.EXPECT().Update(model.Column{ID: 1, Index: 2}).Return(
					model.Column{ID: 1, Index: 2},
					nil,
				)
				s.EXPECT().Columns().Times(4).Return(cr)
			},
			column:   model.Column{ID: 2, Index: 2},
			left:     true,
			expError: nil,
		},
		{
			name: "column is moved right",
			mock: func(c *gomock.Controller, s *mock_store.MockStore, column model.Column) {
				cr := mock_store.NewMockColumnRepo(c)

				cr.EXPECT().GetByID(column.ID).Return(column, nil)
				cr.EXPECT().GetByIndexAndProjectID(column.Index+1, column.ProjectID).Return(
					model.Column{ID: 2, Index: column.Index + 1},
					nil,
				)
				cr.EXPECT().Update(model.Column{ID: 2, Index: 1}).Return(
					model.Column{ID: 2, Index: 1},
					nil,
				)
				cr.EXPECT().Update(model.Column{ID: 1, Index: 2}).Return(
					model.Column{ID: 1, Index: 2},
					nil,
				)
				s.EXPECT().Columns().Times(4).Return(cr)
			},
			column:   model.Column{ID: 1, Index: 1},
			left:     false,
			expError: nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(c, store, tc.column)
			s := newColumnService(store)
			err := s.MoveByID(tc.column.ID, tc.left)

			assert.Equal(t, tc.expError, err)
		})
	}
}

func TestColumnService_DeleteByID(t *testing.T) {
	testcases := []struct {
		name     string
		mock     func(*gomock.Controller, *mock_store.MockStore, model.Column)
		column   model.Column
		expError error
	}{
		{
			name: "column is deleted",
			mock: func(c *gomock.Controller, s *mock_store.MockStore, column model.Column) {
				cr := mock_store.NewMockColumnRepo(c)
				tr := mock_store.NewMockTaskRepo(c)

				cr.EXPECT().GetByID(column.ID).Return(column, nil)
				cr.EXPECT().GetByProjectID(column.ProjectID).Return(
					[]model.Column{
						column,
						model.Column{ID: 2, Name: "Column 2", Index: 2, ProjectID: column.ProjectID},
					},
					nil,
				)
				tr.EXPECT().GetByColumnID(1).Return(
					[]model.Task{model.Task{ID: 1, Name: "Task 1", Index: 1, ColumnID: 1}},
					nil,
				)
				tr.EXPECT().GetByColumnID(2).Return(
					[]model.Task{model.Task{ID: 2, Name: "Task 2", Index: 1, ColumnID: 2}},
					nil,
				)
				tr.EXPECT().Update(model.Task{ID: 1, Name: "Task 1", Index: 2, ColumnID: 2}).Return(
					model.Task{ID: 1, Name: "Task 1", Index: 2, ColumnID: 2},
					nil,
				)
				cr.EXPECT().Update(
					model.Column{ID: 2, Name: "Column 2", Index: 1, ProjectID: column.ProjectID},
				).Return(
					model.Column{ID: 2, Name: "Column 2", Index: 1, ProjectID: column.ProjectID},
					nil,
				)
				cr.EXPECT().DeleteByID(column.ID).Return(nil)
				s.EXPECT().Columns().Times(4).Return(cr)
				s.EXPECT().Tasks().Times(3).Return(tr)
			},
			column:   model.Column{ID: 1, Name: "Column 1", Index: 1, ProjectID: 1},
			expError: nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(c, store, tc.column)
			s := newColumnService(store)

			err := s.DeleteByID(tc.column.ID)
			assert.Equal(t, tc.expError, err)
		})
	}
}

func TestColumnService_Validate(t *testing.T) {
	testcases := []struct {
		name     string
		mock     func(*gomock.Controller, *mock_store.MockStore, model.Column)
		column   model.Column
		expError error
	}{
		{
			name: "column passes validation",
			mock: func(c *gomock.Controller, s *mock_store.MockStore, column model.Column) {
				cr := mock_store.NewMockColumnRepo(c)

				cr.EXPECT().GetByProjectID(column.ProjectID).Return([]model.Column{}, nil)
				s.EXPECT().Columns().Return(cr)
			},
			column:   model.Column{Name: "Column 1", ProjectID: 1},
			expError: nil,
		},
		{
			name:     "column doesn't pass validation because of empty name",
			mock:     func(c *gomock.Controller, s *mock_store.MockStore, column model.Column) {},
			column:   model.Column{Name: ""},
			expError: ErrNameIsRequired,
		},
		{
			name:     "column doesn't pass validation because of too long name",
			mock:     func(c *gomock.Controller, s *mock_store.MockStore, column model.Column) {},
			column:   model.Column{Name: fixedLengthString(256)},
			expError: ErrNameIsTooLong,
		},
		{
			name: "column doesn't pass validation because of invalid project ID",
			mock: func(c *gomock.Controller, s *mock_store.MockStore, column model.Column) {
				cr := mock_store.NewMockColumnRepo(c)

				cr.EXPECT().GetByProjectID(column.ProjectID).Return([]model.Column{}, store.ErrDbQuery)
				s.EXPECT().Columns().Return(cr)
			},
			column:   model.Column{Name: "Column 1", ProjectID: 1},
			expError: store.ErrDbQuery,
		},
		{
			name: "column doesn't pass validation because column with this name exists",
			mock: func(c *gomock.Controller, s *mock_store.MockStore, column model.Column) {
				cr := mock_store.NewMockColumnRepo(c)

				cr.EXPECT().GetByProjectID(column.ProjectID).Return(
					[]model.Column{
						model.Column{ID: 1, Name: column.Name, Index: 1, ProjectID: column.ProjectID},
					},
					nil,
				)
				s.EXPECT().Columns().Return(cr)
			},
			column:   model.Column{Name: "Column 1", ProjectID: 1},
			expError: ErrColumnAlreadyExists,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(c, store, tc.column)
			s := newColumnService(store)
			err := s.Validate(tc.column)

			assert.Equal(t, tc.expError, err)
		})
	}
}
