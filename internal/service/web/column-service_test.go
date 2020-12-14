package web

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/imarrche/tasker/internal/model"
	"github.com/imarrche/tasker/internal/store"
	mock_store "github.com/imarrche/tasker/internal/store/mocks"
)

func TestColumnService_GetAll(t *testing.T) {
	testcases := [...]struct {
		name            string
		mock            func(*mock_store.MockColumnRepository)
		expectedColumns []model.Column
		expectedError   error
	}{
		{
			name: "Columns are retrieved and sorted alphabetically",
			mock: func(cr *mock_store.MockColumnRepository) {
				cr.EXPECT().GetAll().Return(
					[]model.Column{
						model.Column{Index: 3},
						model.Column{Index: 2},
						model.Column{Index: 1},
					},
					nil,
				)
			},
			expectedColumns: []model.Column{
				model.Column{Index: 1},
				model.Column{Index: 2},
				model.Column{Index: 3},
			},
			expectedError: nil,
		},
		{
			name: "Error occured while retrieving columns",
			mock: func(cr *mock_store.MockColumnRepository) {
				cr.EXPECT().GetAll().Return(nil, errors.New("couldn't get columns"))
			},
			expectedColumns: nil,
			expectedError:   errors.New("couldn't get columns"),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			cr := mock_store.NewMockColumnRepository(c)
			tc.mock(cr)
			s := NewColumnService(cr, nil)

			cs, err := s.GetAll()
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedColumns, cs)
		})
	}
}

func TestColumnService_Create(t *testing.T) {
	testcases := [...]struct {
		name           string
		mock           func(*mock_store.MockColumnRepository, model.Column)
		column         model.Column
		expectedColumn model.Column
		expectedError  error
	}{
		{
			name: "Column is created.",
			mock: func(cr *mock_store.MockColumnRepository, c model.Column) {
				cr.EXPECT().Create(c).Return(c, nil)
				cr.EXPECT().GetAllByProjectID(c.Project.ID).Return([]model.Column{}, nil)
			},
			column:         model.Column{Name: "C1", Project: model.Project{ID: 1}},
			expectedColumn: model.Column{Name: "C1", Project: model.Project{ID: 1}},
			expectedError:  nil,
		},
		{
			name:           "Column didn't pass validation.",
			mock:           func(cr *mock_store.MockColumnRepository, c model.Column) {},
			column:         model.Column{Name: ""},
			expectedColumn: model.Column{},
			expectedError:  ErrNameIsRequired,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			cr := mock_store.NewMockColumnRepository(c)
			tc.mock(cr, tc.column)
			s := NewColumnService(cr, nil)

			column, err := s.Create(tc.column)
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedColumn, column)
		})
	}
}

func TestColumnService_GetByID(t *testing.T) {
	testcases := [...]struct {
		name           string
		mock           func(*mock_store.MockColumnRepository, model.Column)
		column         model.Column
		expectedColumn model.Column
		expectedError  error
	}{
		{
			name: "Column is retrieved by ID.",
			mock: func(cr *mock_store.MockColumnRepository, c model.Column) {
				cr.EXPECT().GetByID(c.ID).Return(c, nil)
			},
			column:         model.Column{ID: 1},
			expectedColumn: model.Column{ID: 1},
			expectedError:  nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			cr := mock_store.NewMockColumnRepository(c)
			tc.mock(cr, tc.column)
			s := NewColumnService(cr, nil)

			column, err := s.GetByID(tc.column.ID)
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedColumn, column)
		})
	}
}

func TestColumnService_Update(t *testing.T) {
	testcases := [...]struct {
		name           string
		mock           func(*mock_store.MockColumnRepository, model.Column)
		column         model.Column
		expectedColumn model.Column
		expectedError  error
	}{
		{
			name: "Column is updated.",
			mock: func(cr *mock_store.MockColumnRepository, c model.Column) {
				cr.EXPECT().Update(c).Return(nil)
				cr.EXPECT().GetAllByProjectID(c.Project.ID).Return(
					[]model.Column{},
					nil,
				)
			},
			column:        model.Column{Name: "C1", Project: model.Project{ID: 1}},
			expectedError: nil,
		},
		{
			name:          "Column didn't pass validation.",
			mock:          func(cr *mock_store.MockColumnRepository, c model.Column) {},
			column:        model.Column{},
			expectedError: ErrNameIsRequired,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			cr := mock_store.NewMockColumnRepository(c)
			tc.mock(cr, tc.column)
			s := NewColumnService(cr, nil)

			err := s.Update(tc.column)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestColumnService_DeleteByID(t *testing.T) {
	testcases := [...]struct {
		name string
		mock func(
			*mock_store.MockColumnRepository,
			*mock_store.MockTaskRepository,
			model.Column,
		)
		column        model.Column
		expectedError error
	}{
		{
			name: "Column is deleted.",
			mock: func(
				cr *mock_store.MockColumnRepository,
				tr *mock_store.MockTaskRepository,
				c model.Column,
			) {
				cr.EXPECT().GetByID(c.ID).Return(c, nil)
				cr.EXPECT().GetAllByProjectID(c.Project.ID).Return(
					[]model.Column{model.Column{Name: "C1"}, model.Column{Name: "C2"}},
					nil,
				)
				tr.EXPECT().GetAllByColumnID(c.ID).Return([]model.Task{}, nil)
				cr.EXPECT().DeleteByID(c.ID).Return(nil)
			},
			column:        model.Column{ID: 1, Name: "C1", Project: model.Project{ID: 1}},
			expectedError: nil,
		},
		{
			name: "Column was not found.",
			mock: func(
				cr *mock_store.MockColumnRepository,
				tr *mock_store.MockTaskRepository,
				c model.Column,
			) {
				cr.EXPECT().GetByID(c.ID).Return(model.Column{}, store.ErrNotFound)
			},
			column:        model.Column{ID: 1, Name: "C1"},
			expectedError: store.ErrNotFound,
		},
		{
			name: "Error occured while retrieving all project's columns.",
			mock: func(
				cr *mock_store.MockColumnRepository,
				tr *mock_store.MockTaskRepository,
				c model.Column,
			) {
				cr.EXPECT().GetByID(c.ID).Return(c, nil)
				cr.EXPECT().GetAllByProjectID(c.Project.ID).Return(
					[]model.Column{},
					errors.New("couldn't get columns"),
				)
			},
			column:        model.Column{ID: 1, Name: "C1"},
			expectedError: errors.New("couldn't get columns"),
		},
		{
			name: "Error occured while deleting last column.",
			mock: func(
				cr *mock_store.MockColumnRepository,
				tr *mock_store.MockTaskRepository,
				c model.Column,
			) {
				cr.EXPECT().GetByID(c.ID).Return(c, nil)
				cr.EXPECT().GetAllByProjectID(c.Project.ID).Return(
					[]model.Column{model.Column{ID: 1, Name: "C1"}},
					nil,
				)
			},
			column:        model.Column{ID: 1, Name: "C1"},
			expectedError: ErrLastColumn,
		},
		{
			name: "Error occured while retrieving all column's tasks.",
			mock: func(
				cr *mock_store.MockColumnRepository,
				tr *mock_store.MockTaskRepository,
				c model.Column,
			) {
				cr.EXPECT().GetByID(c.ID).Return(c, nil)
				cr.EXPECT().GetAllByProjectID(c.Project.ID).Return(
					[]model.Column{model.Column{Name: "C1"}, model.Column{Name: "C2"}},
					nil,
				)
				tr.EXPECT().GetAllByColumnID(c.ID).Return(
					[]model.Task{}, errors.New("couldn't get tasks"),
				)
			},
			column:        model.Column{ID: 1, Name: "C1"},
			expectedError: errors.New("couldn't get tasks"),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			cr := mock_store.NewMockColumnRepository(c)
			tr := mock_store.NewMockTaskRepository(c)
			tc.mock(cr, tr, tc.column)
			s := NewColumnService(cr, tr)

			err := s.DeleteByID(tc.column.ID)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestColumnService_Validaete(t *testing.T) {
	testcases := [...]struct {
		name          string
		mock          func(*mock_store.MockColumnRepository, model.Column)
		column        model.Column
		expectedError error
	}{
		{
			name: "Column passes validation.",
			mock: func(cr *mock_store.MockColumnRepository, c model.Column) {
				cr.EXPECT().GetAllByProjectID(c.Project.ID).Return([]model.Column{}, nil)
			},
			column:        model.Column{Name: "C1", Project: model.Project{ID: 1}},
			expectedError: nil,
		},
		{
			name:          "Column name is not provided.",
			mock:          func(cr *mock_store.MockColumnRepository, c model.Column) {},
			column:        model.Column{Name: "", Project: model.Project{ID: 1}},
			expectedError: ErrNameIsRequired,
		},
		{
			name:          "Column name is too long.",
			mock:          func(cr *mock_store.MockColumnRepository, c model.Column) {},
			column:        model.Column{Name: fixedLengthString(256), Project: model.Project{ID: 1}},
			expectedError: ErrNameIsTooLong,
		},
		{
			name:          "Column project is not provided.",
			mock:          func(cr *mock_store.MockColumnRepository, c model.Column) {},
			column:        model.Column{Name: "C1"},
			expectedError: ErrProjectIsRequired,
		},
		{
			name: "Error occures while retrieving all project's columns.",
			mock: func(cr *mock_store.MockColumnRepository, c model.Column) {
				cr.EXPECT().GetAllByProjectID(c.Project.ID).Return(
					[]model.Column{}, errors.New("couldn't get columns"),
				)
			},
			column:        model.Column{Name: "C1", Project: model.Project{ID: 1}},
			expectedError: errors.New("couldn't get columns"),
		},
		{
			name: "Column with this name already exists.",
			mock: func(cr *mock_store.MockColumnRepository, c model.Column) {
				cr.EXPECT().GetAllByProjectID(c.Project.ID).Return(
					[]model.Column{model.Column{Name: "C1"}}, nil,
				)
			},
			column:        model.Column{Name: "C1", Project: model.Project{ID: 1}},
			expectedError: ErrColumnAlreadyExists,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			cr := mock_store.NewMockColumnRepository(c)
			tc.mock(cr, tc.column)
			s := NewColumnService(cr, nil)

			err := s.Validate(tc.column)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
