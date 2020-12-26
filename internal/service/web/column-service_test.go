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
		mock       func(*mock_store.MockStore, *gomock.Controller, model.Project)
		project    model.Project
		expColumns []model.Column
		expError   error
	}{
		{
			name: "Columns are retrieved and sorted alphabetically",
			mock: func(s *mock_store.MockStore, c *gomock.Controller, p model.Project) {
				cr := mock_store.NewMockColumnRepo(c)

				cr.EXPECT().GetByProjectID(p.ID).Return(
					[]model.Column{
						model.Column{ID: 1, Name: "C1", Index: 3, ProjectID: 1},
						model.Column{ID: 2, Name: "C2", Index: 2, ProjectID: 1},
						model.Column{ID: 3, Name: "C3", Index: 1, ProjectID: 1},
					},
					nil,
				)
				s.EXPECT().Columns().Return(cr)
			},
			project: model.Project{ID: 1, Name: "P"},
			expColumns: []model.Column{
				model.Column{ID: 3, Name: "C3", Index: 1, ProjectID: 1},
				model.Column{ID: 2, Name: "C2", Index: 2, ProjectID: 1},
				model.Column{ID: 1, Name: "C1", Index: 3, ProjectID: 1},
			},
			expError: nil,
		},
		{
			name: "Error occures while retrieving columns",
			mock: func(s *mock_store.MockStore, c *gomock.Controller, p model.Project) {
				cr := mock_store.NewMockColumnRepo(c)

				cr.EXPECT().GetByProjectID(p.ID).Return(nil, store.ErrDbQuery)
				s.EXPECT().Columns().Return(cr)
			},
			expColumns: nil,
			expError:   store.ErrDbQuery,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(store, c, tc.project)
			s := newColumnService(store)

			cs, err := s.GetByProjectID(tc.project.ID)
			assert.Equal(t, tc.expError, err)
			assert.Equal(t, tc.expColumns, cs)
		})
	}
}

func TestColumnService_Create(t *testing.T) {
	testcases := []struct {
		name      string
		mock      func(*mock_store.MockStore, *gomock.Controller, model.Column)
		column    model.Column
		expColumn model.Column
		expError  error
	}{
		{
			name: "Column is created",
			mock: func(s *mock_store.MockStore, c *gomock.Controller, column model.Column) {
				pr := mock_store.NewMockProjectRepo(c)
				cr := mock_store.NewMockColumnRepo(c)

				pr.EXPECT().GetByID(column.ProjectID).Return(model.Project{}, nil)
				cr.EXPECT().GetByProjectID(column.ProjectID).Return([]model.Column{}, nil)
				cr.EXPECT().Create(column).Return(
					model.Column{
						ID:        1,
						Name:      column.Name,
						Index:     column.Index,
						ProjectID: column.ProjectID,
					},
					nil,
				)
				s.EXPECT().Columns().Times(2).Return(cr)
				s.EXPECT().Projects().Return(pr)
			},
			column:    model.Column{Name: "C", Index: 1, ProjectID: 1},
			expColumn: model.Column{ID: 1, Name: "C", Index: 1, ProjectID: 1},
			expError:  nil,
		},
		{
			name:      "Column doesn't pass validation",
			mock:      func(s *mock_store.MockStore, c *gomock.Controller, column model.Column) {},
			column:    model.Column{Name: ""},
			expColumn: model.Column{},
			expError:  ErrNameIsRequired,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(store, c, tc.column)
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
		mock      func(*mock_store.MockStore, *gomock.Controller, model.Column)
		column    model.Column
		expColumn model.Column
		expError  error
	}{
		{
			name: "Column is retrieved",
			mock: func(s *mock_store.MockStore, c *gomock.Controller, column model.Column) {
				cr := mock_store.NewMockColumnRepo(c)

				cr.EXPECT().GetByID(column.ID).Return(column, nil)
				s.EXPECT().Columns().Return(cr)
			},
			column:    model.Column{ID: 1, Name: "C", Index: 1, ProjectID: 1},
			expColumn: model.Column{ID: 1, Name: "C", Index: 1, ProjectID: 1},
			expError:  nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(store, c, tc.column)
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
		mock      func(*mock_store.MockStore, *gomock.Controller, model.Column)
		column    model.Column
		expColumn model.Column
		expError  error
	}{
		{
			name: "Column is updated",
			mock: func(s *mock_store.MockStore, c *gomock.Controller, column model.Column) {
				cr := mock_store.NewMockColumnRepo(c)

				cr.EXPECT().GetByProjectID(column.ProjectID).Return([]model.Column{}, nil)
				cr.EXPECT().Update(column).Return(column, nil)
				s.EXPECT().Columns().Times(2).Return(cr)
			},
			column:    model.Column{ID: 1, Name: "C", Index: 1, ProjectID: 1},
			expColumn: model.Column{ID: 1, Name: "C", Index: 1, ProjectID: 1},
			expError:  nil,
		},
		{
			name:      "Column doesn't pass validation",
			mock:      func(s *mock_store.MockStore, c *gomock.Controller, column model.Column) {},
			column:    model.Column{ID: 1},
			expColumn: model.Column{},
			expError:  ErrNameIsRequired,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(store, c, tc.column)
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
		mock     func(*mock_store.MockStore, *gomock.Controller, model.Column)
		column   model.Column
		left     bool
		expError error
	}{
		{
			name: "Column is moved left",
			mock: func(s *mock_store.MockStore, c *gomock.Controller, column model.Column) {
				cr := mock_store.NewMockColumnRepo(c)

				cr.EXPECT().GetByID(column.ID).Return(column, nil)
				cr.EXPECT().GetByIndexAndProjectID(column.Index-1, column.ProjectID).Return(
					model.Column{ID: 1, Index: 1},
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
			name: "Column is moved right",
			mock: func(s *mock_store.MockStore, c *gomock.Controller, column model.Column) {
				cr := mock_store.NewMockColumnRepo(c)

				cr.EXPECT().GetByID(column.ID).Return(column, nil)
				cr.EXPECT().GetByIndexAndProjectID(column.Index+1, column.ProjectID).Return(
					model.Column{ID: 2, Index: 2},
					nil,
				)
				cr.EXPECT().Update(model.Column{ID: 1, Index: 2}).Return(
					model.Column{ID: 1, Index: 2},
					nil,
				)
				cr.EXPECT().Update(model.Column{ID: 2, Index: 1}).Return(
					model.Column{ID: 2, Index: 1},
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
			tc.mock(store, c, tc.column)
			s := newColumnService(store)

			err := s.MoveByID(tc.column.ID, tc.left)
			assert.Equal(t, tc.expError, err)
		})
	}
}

func TestColumnService_DeleteByID(t *testing.T) {
	testcases := []struct {
		name     string
		mock     func(*mock_store.MockStore, *gomock.Controller, model.Column)
		column   model.Column
		expError error
	}{
		{
			name: "Column is deleted",
			mock: func(s *mock_store.MockStore, c *gomock.Controller, column model.Column) {
				cr := mock_store.NewMockColumnRepo(c)
				tr := mock_store.NewMockTaskRepo(c)

				cr.EXPECT().GetByID(column.ID).Return(column, nil)
				cr.EXPECT().GetByProjectID(column.ProjectID).Return(
					[]model.Column{
						model.Column{ID: 1, Name: "C", Index: 1, ProjectID: 1},
						model.Column{ID: 2, Name: "C", Index: 2, ProjectID: 1},
					},
					nil,
				)
				tr.EXPECT().GetByColumnID(1).Return(
					[]model.Task{model.Task{ID: 1, Index: 1, ColumnID: 1}},
					nil,
				)
				tr.EXPECT().GetByColumnID(2).Return(
					[]model.Task{model.Task{ID: 2, Index: 1, ColumnID: 2}},
					nil,
				)
				tr.EXPECT().Update(model.Task{ID: 1, Index: 2, ColumnID: 2}).Return(
					model.Task{ID: 1, Index: 2, ColumnID: 2},
					nil,
				)
				cr.EXPECT().DeleteByID(column.ID).Return(nil)
				cr.EXPECT().Update(
					model.Column{ID: 2, Name: "C", Index: 1, ProjectID: 1},
				).Return(
					model.Column{ID: 2, Name: "C", Index: 1, ProjectID: 1},
					nil,
				)
				s.EXPECT().Columns().Times(4).Return(cr)
				s.EXPECT().Tasks().Times(3).Return(tr)
			},
			column:   model.Column{ID: 1, Name: "C", Index: 1, ProjectID: 1},
			expError: nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(store, c, tc.column)
			s := newColumnService(store)

			err := s.DeleteByID(tc.column.ID)
			assert.Equal(t, tc.expError, err)
		})
	}
}

func TestColumnService_Validate(t *testing.T) {
	testcases := []struct {
		name     string
		mock     func(*mock_store.MockStore, *gomock.Controller, model.Column)
		column   model.Column
		expError error
	}{
		{
			name: "Column passes validation",
			mock: func(s *mock_store.MockStore, c *gomock.Controller, column model.Column) {
				cr := mock_store.NewMockColumnRepo(c)

				cr.EXPECT().GetByProjectID(column.ProjectID).Return([]model.Column{}, nil)
				s.EXPECT().Columns().Return(cr)
			},
			column:   model.Column{Name: "C1", ProjectID: 1},
			expError: nil,
		},
		{
			name:     "Column doesn't pass validation because of empty name",
			mock:     func(s *mock_store.MockStore, c *gomock.Controller, column model.Column) {},
			column:   model.Column{Name: ""},
			expError: ErrNameIsRequired,
		},
		{
			name:     "Column doesn't pass validation because of too long name",
			mock:     func(s *mock_store.MockStore, c *gomock.Controller, column model.Column) {},
			column:   model.Column{Name: fixedLengthString(256)},
			expError: ErrNameIsTooLong,
		},
		{
			name: "Column doesn't pass validation because of invalid project ID",
			mock: func(s *mock_store.MockStore, c *gomock.Controller, column model.Column) {
				cr := mock_store.NewMockColumnRepo(c)

				cr.EXPECT().GetByProjectID(column.ProjectID).Return(
					[]model.Column{},
					store.ErrDbQuery,
				)
				s.EXPECT().Columns().Return(cr)
			},
			column:   model.Column{Name: "C1", ProjectID: 1},
			expError: store.ErrDbQuery,
		},
		{
			name: "Column doesn't pass validation because column with this name exists",
			mock: func(s *mock_store.MockStore, c *gomock.Controller, column model.Column) {
				cr := mock_store.NewMockColumnRepo(c)

				cr.EXPECT().GetByProjectID(column.ProjectID).Return(
					[]model.Column{model.Column{ID: 1, Name: "C", ProjectID: column.ProjectID}},
					nil,
				)
				s.EXPECT().Columns().Return(cr)
			},
			column:   model.Column{Name: "C", ProjectID: 1},
			expError: ErrColumnAlreadyExists,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			store := mock_store.NewMockStore(c)
			tc.mock(store, c, tc.column)
			s := newColumnService(store)

			err := s.Validate(tc.column)
			assert.Equal(t, tc.expError, err)
		})
	}
}
