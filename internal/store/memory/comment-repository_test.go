package memory

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/imarrche/tasker/internal/model"
	"github.com/imarrche/tasker/internal/store"
)

func TestCommentRepository_GetAll(t *testing.T) {
	s := NewStore()
	c1 := model.Comment{Text: "Comment 1"}
	c2 := model.Comment{Text: "Comment 2"}
	s.db.comments[1] = c1
	s.db.comments[2] = c2

	cs, err := s.Comments().GetAll()

	assert.NoError(t, err)
	assert.Equal(t, 2, len(cs))
}

func TestCommentRepository_Create(t *testing.T) {
	s := NewStore()
	c1 := model.Comment{Text: "Comment 1"}
	c2 := model.Comment{Text: "Comment 2"}

	c1FromRepo, err1 := s.Comments().Create(c1)
	c2FromRepo, err2 := s.Comments().Create(c2)

	assert.NoError(t, err1)
	assert.Equal(t, 1, c1FromRepo.ID)
	assert.Equal(t, c1.Text, c1FromRepo.Text)
	assert.NoError(t, err2)
	assert.Equal(t, 2, c2FromRepo.ID)
	assert.Equal(t, c2.Text, c2FromRepo.Text)
}

func TestCommentRepository_GetByID(t *testing.T) {
	s := NewStore()
	c1 := model.Comment{ID: 1, Text: "Comment 1"}
	s.db.comments[c1.ID] = c1

	c1FromRepo, err1 := s.Comments().GetByID(c1.ID)
	_, err2 := s.Comments().GetByID(2)

	assert.NoError(t, err1)
	assert.Equal(t, c1.ID, c1FromRepo.ID)
	assert.Equal(t, c1.Text, c1FromRepo.Text)
	assert.Equal(t, store.ErrNotFound, err2)
}

func TestCommentRepository_Update(t *testing.T) {
	s := NewStore()
	c1 := model.Comment{ID: 1, Text: "Comment 1"}
	c2 := model.Comment{ID: 2, Text: "Comment 2"}
	s.db.comments[c1.ID] = c1

	c1.Text = "Updated text"
	err1 := s.Comments().Update(c1)
	err2 := s.Comments().Update(c2)

	assert.NoError(t, err1)
	assert.Equal(t, c1.Text, s.db.comments[c1.ID].Text)
	assert.Equal(t, store.ErrNotFound, err2)
}

func TestCommentRepository_Delete(t *testing.T) {
	s := NewStore()
	c1 := model.Comment{ID: 1, Text: "Comment 1"}
	c2 := model.Comment{ID: 2, Text: "Comment 2"}
	s.db.comments[c1.ID] = c1

	err1 := s.Comments().DeleteByID(c1.ID)
	err2 := s.Comments().DeleteByID(c2.ID)

	assert.NoError(t, err1)
	assert.Equal(t, 0, len(s.db.comments))
	assert.Equal(t, store.ErrNotFound, err2)
}
