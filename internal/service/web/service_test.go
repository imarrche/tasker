package web

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	mock_store "github.com/imarrche/tasker/internal/store/mocks"
)

func TestService_Projects(t *testing.T) {
	c := gomock.NewController(t)
	c.Finish()

	store := mock_store.NewMockStore(c)
	s := NewService(store)

	assert.Equal(t, newProjectService(store), s.Projects())
}

func TestService_Columns(t *testing.T) {
	c := gomock.NewController(t)
	c.Finish()

	store := mock_store.NewMockStore(c)
	s := NewService(store)

	assert.Equal(t, newColumnService(store), s.Columns())
}

func TestService_Tasks(t *testing.T) {
	c := gomock.NewController(t)
	c.Finish()

	store := mock_store.NewMockStore(c)
	s := NewService(store)

	assert.Equal(t, newTaskService(store), s.Tasks())
}

func TestService_Comments(t *testing.T) {
	c := gomock.NewController(t)
	c.Finish()

	store := mock_store.NewMockStore(c)
	s := NewService(store)

	assert.Equal(t, newCommentService(store), s.Comments())
}
