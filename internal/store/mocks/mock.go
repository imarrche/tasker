// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock_store is a generated GoMock package.
package mock_store

import (
	gomock "github.com/golang/mock/gomock"
	model "github.com/imarrche/tasker/internal/model"
	store "github.com/imarrche/tasker/internal/store"
	reflect "reflect"
)

// MockStore is a mock of Store interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// Open mocks base method
func (m *MockStore) Open() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Open")
	ret0, _ := ret[0].(error)
	return ret0
}

// Open indicates an expected call of Open
func (mr *MockStoreMockRecorder) Open() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockStore)(nil).Open))
}

// Projects mocks base method
func (m *MockStore) Projects() store.ProjectRepo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Projects")
	ret0, _ := ret[0].(store.ProjectRepo)
	return ret0
}

// Projects indicates an expected call of Projects
func (mr *MockStoreMockRecorder) Projects() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Projects", reflect.TypeOf((*MockStore)(nil).Projects))
}

// Columns mocks base method
func (m *MockStore) Columns() store.ColumnRepo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Columns")
	ret0, _ := ret[0].(store.ColumnRepo)
	return ret0
}

// Columns indicates an expected call of Columns
func (mr *MockStoreMockRecorder) Columns() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Columns", reflect.TypeOf((*MockStore)(nil).Columns))
}

// Tasks mocks base method
func (m *MockStore) Tasks() store.TaskRepo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tasks")
	ret0, _ := ret[0].(store.TaskRepo)
	return ret0
}

// Tasks indicates an expected call of Tasks
func (mr *MockStoreMockRecorder) Tasks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tasks", reflect.TypeOf((*MockStore)(nil).Tasks))
}

// Comments mocks base method
func (m *MockStore) Comments() store.CommentRepo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Comments")
	ret0, _ := ret[0].(store.CommentRepo)
	return ret0
}

// Comments indicates an expected call of Comments
func (mr *MockStoreMockRecorder) Comments() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Comments", reflect.TypeOf((*MockStore)(nil).Comments))
}

// Close mocks base method
func (m *MockStore) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockStoreMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStore)(nil).Close))
}

// MockProjectRepo is a mock of ProjectRepo interface
type MockProjectRepo struct {
	ctrl     *gomock.Controller
	recorder *MockProjectRepoMockRecorder
}

// MockProjectRepoMockRecorder is the mock recorder for MockProjectRepo
type MockProjectRepoMockRecorder struct {
	mock *MockProjectRepo
}

// NewMockProjectRepo creates a new mock instance
func NewMockProjectRepo(ctrl *gomock.Controller) *MockProjectRepo {
	mock := &MockProjectRepo{ctrl: ctrl}
	mock.recorder = &MockProjectRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProjectRepo) EXPECT() *MockProjectRepoMockRecorder {
	return m.recorder
}

// GetAll mocks base method
func (m *MockProjectRepo) GetAll() ([]model.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]model.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockProjectRepoMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockProjectRepo)(nil).GetAll))
}

// Create mocks base method
func (m *MockProjectRepo) Create(arg0 model.Project) (model.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(model.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockProjectRepoMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProjectRepo)(nil).Create), arg0)
}

// GetByID mocks base method
func (m *MockProjectRepo) GetByID(arg0 int) (model.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0)
	ret0, _ := ret[0].(model.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockProjectRepoMockRecorder) GetByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockProjectRepo)(nil).GetByID), arg0)
}

// Update mocks base method
func (m *MockProjectRepo) Update(arg0 model.Project) (model.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(model.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockProjectRepoMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockProjectRepo)(nil).Update), arg0)
}

// DeleteByID mocks base method
func (m *MockProjectRepo) DeleteByID(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID
func (mr *MockProjectRepoMockRecorder) DeleteByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockProjectRepo)(nil).DeleteByID), arg0)
}

// MockColumnRepo is a mock of ColumnRepo interface
type MockColumnRepo struct {
	ctrl     *gomock.Controller
	recorder *MockColumnRepoMockRecorder
}

// MockColumnRepoMockRecorder is the mock recorder for MockColumnRepo
type MockColumnRepoMockRecorder struct {
	mock *MockColumnRepo
}

// NewMockColumnRepo creates a new mock instance
func NewMockColumnRepo(ctrl *gomock.Controller) *MockColumnRepo {
	mock := &MockColumnRepo{ctrl: ctrl}
	mock.recorder = &MockColumnRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockColumnRepo) EXPECT() *MockColumnRepoMockRecorder {
	return m.recorder
}

// GetByProjectID mocks base method
func (m *MockColumnRepo) GetByProjectID(arg0 int) ([]model.Column, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByProjectID", arg0)
	ret0, _ := ret[0].([]model.Column)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByProjectID indicates an expected call of GetByProjectID
func (mr *MockColumnRepoMockRecorder) GetByProjectID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByProjectID", reflect.TypeOf((*MockColumnRepo)(nil).GetByProjectID), arg0)
}

// Create mocks base method
func (m *MockColumnRepo) Create(arg0 model.Column) (model.Column, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(model.Column)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockColumnRepoMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockColumnRepo)(nil).Create), arg0)
}

// GetByID mocks base method
func (m *MockColumnRepo) GetByID(arg0 int) (model.Column, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0)
	ret0, _ := ret[0].(model.Column)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockColumnRepoMockRecorder) GetByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockColumnRepo)(nil).GetByID), arg0)
}

// GetByIndexAndProjectID mocks base method
func (m *MockColumnRepo) GetByIndexAndProjectID(arg0, arg1 int) (model.Column, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIndexAndProjectID", arg0, arg1)
	ret0, _ := ret[0].(model.Column)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIndexAndProjectID indicates an expected call of GetByIndexAndProjectID
func (mr *MockColumnRepoMockRecorder) GetByIndexAndProjectID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIndexAndProjectID", reflect.TypeOf((*MockColumnRepo)(nil).GetByIndexAndProjectID), arg0, arg1)
}

// Update mocks base method
func (m *MockColumnRepo) Update(arg0 model.Column) (model.Column, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(model.Column)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockColumnRepoMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockColumnRepo)(nil).Update), arg0)
}

// DeleteByID mocks base method
func (m *MockColumnRepo) DeleteByID(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID
func (mr *MockColumnRepoMockRecorder) DeleteByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockColumnRepo)(nil).DeleteByID), arg0)
}

// MockTaskRepo is a mock of TaskRepo interface
type MockTaskRepo struct {
	ctrl     *gomock.Controller
	recorder *MockTaskRepoMockRecorder
}

// MockTaskRepoMockRecorder is the mock recorder for MockTaskRepo
type MockTaskRepoMockRecorder struct {
	mock *MockTaskRepo
}

// NewMockTaskRepo creates a new mock instance
func NewMockTaskRepo(ctrl *gomock.Controller) *MockTaskRepo {
	mock := &MockTaskRepo{ctrl: ctrl}
	mock.recorder = &MockTaskRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTaskRepo) EXPECT() *MockTaskRepoMockRecorder {
	return m.recorder
}

// GetByColumnID mocks base method
func (m *MockTaskRepo) GetByColumnID(arg0 int) ([]model.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByColumnID", arg0)
	ret0, _ := ret[0].([]model.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByColumnID indicates an expected call of GetByColumnID
func (mr *MockTaskRepoMockRecorder) GetByColumnID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByColumnID", reflect.TypeOf((*MockTaskRepo)(nil).GetByColumnID), arg0)
}

// Create mocks base method
func (m *MockTaskRepo) Create(arg0 model.Task) (model.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(model.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockTaskRepoMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTaskRepo)(nil).Create), arg0)
}

// GetByID mocks base method
func (m *MockTaskRepo) GetByID(arg0 int) (model.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0)
	ret0, _ := ret[0].(model.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockTaskRepoMockRecorder) GetByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockTaskRepo)(nil).GetByID), arg0)
}

// GetByIndexAndColumnID mocks base method
func (m *MockTaskRepo) GetByIndexAndColumnID(arg0, arg1 int) (model.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIndexAndColumnID", arg0, arg1)
	ret0, _ := ret[0].(model.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIndexAndColumnID indicates an expected call of GetByIndexAndColumnID
func (mr *MockTaskRepoMockRecorder) GetByIndexAndColumnID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIndexAndColumnID", reflect.TypeOf((*MockTaskRepo)(nil).GetByIndexAndColumnID), arg0, arg1)
}

// Update mocks base method
func (m *MockTaskRepo) Update(arg0 model.Task) (model.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(model.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockTaskRepoMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTaskRepo)(nil).Update), arg0)
}

// DeleteByID mocks base method
func (m *MockTaskRepo) DeleteByID(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID
func (mr *MockTaskRepoMockRecorder) DeleteByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockTaskRepo)(nil).DeleteByID), arg0)
}

// MockCommentRepo is a mock of CommentRepo interface
type MockCommentRepo struct {
	ctrl     *gomock.Controller
	recorder *MockCommentRepoMockRecorder
}

// MockCommentRepoMockRecorder is the mock recorder for MockCommentRepo
type MockCommentRepoMockRecorder struct {
	mock *MockCommentRepo
}

// NewMockCommentRepo creates a new mock instance
func NewMockCommentRepo(ctrl *gomock.Controller) *MockCommentRepo {
	mock := &MockCommentRepo{ctrl: ctrl}
	mock.recorder = &MockCommentRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCommentRepo) EXPECT() *MockCommentRepoMockRecorder {
	return m.recorder
}

// GetByTaskID mocks base method
func (m *MockCommentRepo) GetByTaskID(arg0 int) ([]model.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByTaskID", arg0)
	ret0, _ := ret[0].([]model.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByTaskID indicates an expected call of GetByTaskID
func (mr *MockCommentRepoMockRecorder) GetByTaskID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByTaskID", reflect.TypeOf((*MockCommentRepo)(nil).GetByTaskID), arg0)
}

// Create mocks base method
func (m *MockCommentRepo) Create(arg0 model.Comment) (model.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(model.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockCommentRepoMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCommentRepo)(nil).Create), arg0)
}

// GetByID mocks base method
func (m *MockCommentRepo) GetByID(arg0 int) (model.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0)
	ret0, _ := ret[0].(model.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockCommentRepoMockRecorder) GetByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockCommentRepo)(nil).GetByID), arg0)
}

// Update mocks base method
func (m *MockCommentRepo) Update(arg0 model.Comment) (model.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(model.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockCommentRepoMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCommentRepo)(nil).Update), arg0)
}

// DeleteByID mocks base method
func (m *MockCommentRepo) DeleteByID(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID
func (mr *MockCommentRepoMockRecorder) DeleteByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockCommentRepo)(nil).DeleteByID), arg0)
}
