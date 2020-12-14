package web

import (
	"sort"

	"github.com/imarrche/tasker/internal/model"
	"github.com/imarrche/tasker/internal/service"
	"github.com/imarrche/tasker/internal/store"
)

// CommentService is a web comment service.
type CommentService struct {
	taskRepo    store.TaskRepository
	commentRepo store.CommentRepository
}

// NewCommentService creates and returns a new CommentService instance.
func NewCommentService(tr store.TaskRepository, cr store.CommentRepository) service.CommentService {
	return &CommentService{taskRepo: tr, commentRepo: cr}
}

// GetAll returns all comments sorted by creating date(from newest to oldest).
func (s *CommentService) GetAll() ([]model.Comment, error) {
	cs, err := s.commentRepo.GetAll()
	if err != nil {
		return nil, err
	}

	sort.SliceStable(cs, func(i, j int) bool {
		return cs[i].CreatedAt.After(cs[j].CreatedAt)
	})

	return cs, nil
}

// Create creates a new comment.
func (s *CommentService) Create(c model.Comment) (model.Comment, error) {
	if err := s.Validate(c); err != nil {
		return model.Comment{}, err
	}

	return s.commentRepo.Create(c)
}

// GetByID returns comment with specific ID.
func (s *CommentService) GetByID(id int) (model.Comment, error) {
	return s.commentRepo.GetByID(id)
}

// Update updates a comment.
func (s *CommentService) Update(c model.Comment) error {
	if err := s.Validate(c); err != nil {
		return err
	}

	return s.commentRepo.Update(c)
}

// DeleteByID deletes a comment with specific ID.
func (s *CommentService) DeleteByID(id int) error {
	return s.commentRepo.DeleteByID(id)
}

// Validate validates a comment.
func (s *CommentService) Validate(c model.Comment) error {
	if len(c.Text) == 0 {
		return ErrTextIsRequired
	} else if len(c.Text) > 5000 {
		return ErrTextIsTooLong
	}

	if _, err := s.taskRepo.GetByID(c.Task.ID); err != nil {
		return ErrInvalidTask
	}

	return nil
}
