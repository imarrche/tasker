package memory

import (
	"sync"

	"github.com/imarrche/tasker/internal/model"
	"github.com/imarrche/tasker/internal/store"
)

// CommentRepository is an in memory comment repository.
type CommentRepository struct {
	store *Store
	m     sync.RWMutex
}

// NewCommentRepository creates and returns a new CommentRepository instance.
func NewCommentRepository() *CommentRepository {
	return &CommentRepository{store: NewStore()}
}

// GetAll returns all comments.
func (r *CommentRepository) GetAll() ([]model.Comment, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	comments := []model.Comment{}
	for _, c := range r.store.comments {
		comments = append(comments, c)
	}

	return comments, nil
}

// Create creates and returns a new comment.
func (r *CommentRepository) Create(c model.Comment) (model.Comment, error) {
	r.m.Lock()
	defer r.m.Unlock()

	c.ID = len(r.store.comments) + 1
	r.store.comments[c.ID] = c

	return c, nil
}

// GetByID returns a comment with specific ID.
func (r *CommentRepository) GetByID(id int) (model.Comment, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	if c, ok := r.store.comments[id]; ok {
		return c, nil
	}

	return model.Comment{}, store.ErrNotFound
}

// Update updates a comment.
func (r *CommentRepository) Update(c model.Comment) error {
	r.m.Lock()
	defer r.m.Unlock()

	if _, ok := r.store.comments[c.ID]; ok {
		r.store.comments[c.ID] = c
		return nil
	}

	return store.ErrNotFound
}

// Delete deletes a comment.
func (r *CommentRepository) Delete(c model.Comment) error {
	r.m.Lock()
	defer r.m.Unlock()

	if _, ok := r.store.comments[c.ID]; ok {
		delete(r.store.comments, c.ID)
		return nil
	}

	return store.ErrNotFound
}
