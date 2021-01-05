package web

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	mock_store "github.com/imarrche/tasker/internal/store/mocks"
)

func TestService_Projects(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()

	store := mock_store.NewMockStore(c)

	assert.Equal(t, newProjectService(store), NewService(store).Projects())
}

func TestService_Columns(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()

	store := mock_store.NewMockStore(c)

	assert.Equal(t, newColumnService(store), NewService(store).Columns())
}

func TestService_Tasks(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()

	store := mock_store.NewMockStore(c)

	assert.Equal(t, newTaskService(store), NewService(store).Tasks())
}

func TestService_Comments(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()

	store := mock_store.NewMockStore(c)

	assert.Equal(t, newCommentService(store), NewService(store).Comments())
}
