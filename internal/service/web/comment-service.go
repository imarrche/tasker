package web

import (
	"sort"
	"time"

	"github.com/imarrche/tasker/internal/model"
	"github.com/imarrche/tasker/internal/store"
)

// commentService is the web comment service.
type commentService struct {
	store store.Store
}

// newCommentService creates and returns a new commentService instance.
func newCommentService(s store.Store) *commentService {
	return &commentService{store: s}
}

// GetByTaskID returns all comments with specific task ID sorted by creation time.
func (s *commentService) GetByTaskID(id int) ([]model.Comment, error) {
	cs, err := s.store.Comments().GetByTaskID(id)
	if err != nil {
		return nil, err
	}

	sort.SliceStable(cs, func(i, j int) bool {
		return cs[i].CreatedAt.After(cs[j].CreatedAt)
	})

	return cs, nil
}

// Create creates a new comment.
func (s *commentService) Create(c model.Comment) (model.Comment, error) {
	if err := s.Validate(c); err != nil {
		return model.Comment{}, err
	}
	if _, err := s.store.Tasks().GetByID(c.TaskID); err != nil {
		return model.Comment{}, err
	}

	c.CreatedAt = time.Now()

	return s.store.Comments().Create(c)
}

// GetByID returns the comment with specific ID.
func (s *commentService) GetByID(id int) (model.Comment, error) {
	return s.store.Comments().GetByID(id)
}

// Update updates a comment.
func (s *commentService) Update(c model.Comment) (model.Comment, error) {
	if err := s.Validate(c); err != nil {
		return model.Comment{}, err
	}

	return s.store.Comments().Update(c)
}

// DeleteByID deletes the comment with specific ID.
func (s *commentService) DeleteByID(id int) error {
	return s.store.Comments().DeleteByID(id)
}

// Validate validates a comment.
func (s *commentService) Validate(c model.Comment) error {
	if len(c.Text) == 0 {
		return ErrTextIsRequired
	} else if len(c.Text) > 5000 {
		return ErrTextIsTooLong
	}

	return nil
}
