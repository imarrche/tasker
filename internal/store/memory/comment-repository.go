package memory

import (
	"sync"

	"github.com/imarrche/tasker/internal/model"
	"github.com/imarrche/tasker/internal/store"
)

// CommentRepository is an in memory comment repository.
type CommentRepository struct {
	db *inMemoryDb
	m  sync.RWMutex
}

// NewCommentRepository creates and returns a new CommentRepository instance.
func NewCommentRepository(db *inMemoryDb) *CommentRepository {
	return &CommentRepository{db: db}
}

// GetAll returns all comments.
func (r *CommentRepository) GetAll() ([]model.Comment, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	comments := []model.Comment{}
	for _, c := range r.db.comments {
		comments = append(comments, c)
	}

	return comments, nil
}

// Create creates and returns a new comment.
func (r *CommentRepository) Create(c model.Comment) (model.Comment, error) {
	r.m.Lock()
	defer r.m.Unlock()

	c.ID = len(r.db.comments) + 1
	r.db.comments[c.ID] = c

	return c, nil
}

// GetByID returns a comment with specific ID.
func (r *CommentRepository) GetByID(id int) (model.Comment, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	if c, ok := r.db.comments[id]; ok {
		return c, nil
	}

	return model.Comment{}, store.ErrNotFound
}

// Update updates a comment.
func (r *CommentRepository) Update(c model.Comment) error {
	r.m.Lock()
	defer r.m.Unlock()

	if _, ok := r.db.comments[c.ID]; ok {
		r.db.comments[c.ID] = c
		return nil
	}

	return store.ErrNotFound
}

// Delete deletes a comment.
func (r *CommentRepository) Delete(c model.Comment) error {
	r.m.Lock()
	defer r.m.Unlock()

	if _, ok := r.db.comments[c.ID]; ok {
		delete(r.db.comments, c.ID)
		return nil
	}

	return store.ErrNotFound
}
