package inmem

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/imarrche/tasker/internal/model"
	"github.com/imarrche/tasker/internal/store"
)

func TestCommentRepo_GetByTaskID(t *testing.T) {
	s := TestStoreWithFixtures()

	cs, err := s.Comments().GetByTaskID(1)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(cs))
}

func TestCommentRepo_Create(t *testing.T) {
	s := TestStoreWithFixtures()

	c, err := s.Comments().Create(model.Comment{Text: "Comment 1", TaskID: 2})

	assert.NoError(t, err)
	assert.Equal(t, len(s.db.comments), c.ID)
	assert.Equal(t, "Comment 1", c.Text)
}

func TestCommentRepo_GetByID(t *testing.T) {
	s := TestStoreWithFixtures()

	c1, err1 := s.Comments().GetByID(1)
	_, err2 := s.Comments().GetByID(4)

	assert.NoError(t, err1)
	assert.Equal(t, 1, c1.ID)
	assert.Equal(t, store.ErrNotFound, err2)
}

func TestCommentRepo_Update(t *testing.T) {
	s := TestStoreWithFixtures()

	err1 := s.Comments().Update(model.Comment{ID: 1, Text: "Updated comment 1"})
	err2 := s.Comments().Update(model.Comment{ID: 4})

	assert.NoError(t, err1)
	assert.Equal(t, "Updated comment 1", s.db.comments[1].Text)
	assert.Equal(t, store.ErrNotFound, err2)
}

func TestCommentRepo_DeleteByID(t *testing.T) {
	s := TestStoreWithFixtures()

	err1 := s.Comments().DeleteByID(1)
	err2 := s.Comments().DeleteByID(4)

	assert.NoError(t, err1)
	assert.Equal(t, 2, len(s.db.comments))
	assert.Equal(t, store.ErrNotFound, err2)
}
