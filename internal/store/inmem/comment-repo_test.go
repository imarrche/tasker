package inmem

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/imarrche/tasker/internal/model"
)

func TestCommentRepo_GetByTaskID(t *testing.T) {
	s := TestStoreWithFixtures()

	cs, err := s.Comments().GetByTaskID(1)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(cs))
}

func TestCommentRepo_Create(t *testing.T) {
	s := TestStoreWithFixtures()

	c, err := s.Comments().Create(model.Comment{Text: "Comment 4", TaskID: 2})

	assert.NoError(t, err)
	assert.Equal(t, model.Comment{ID: 4, Text: "Comment 4", TaskID: 2}, c)
}

func TestCommentRepo_GetByID(t *testing.T) {
	s := TestStoreWithFixtures()

	c, err := s.Comments().GetByID(1)

	assert.NoError(t, err)
	assert.Equal(t, 1, c.ID)
}

func TestCommentRepo_Update(t *testing.T) {
	s := TestStoreWithFixtures()
	comment := model.Comment{ID: 1, Text: "Updated comment 1", TaskID: 1}

	c, err := s.Comments().Update(comment)

	assert.NoError(t, err)
	assert.Equal(t, comment, c)
}

func TestCommentRepo_DeleteByID(t *testing.T) {
	s := TestStoreWithFixtures()

	err := s.Comments().DeleteByID(1)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(s.db.comments))
}
