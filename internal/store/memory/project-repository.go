package memory

import (
	"sync"

	"github.com/imarrche/tasker/internal/model"
	"github.com/imarrche/tasker/internal/store"
)

// ProjectRepository is an in memory project repository.
type ProjectRepository struct {
	store *Store
	m     sync.RWMutex
}

// NewProjectRepository creates and returns a new ProjectRepository instance.
func NewProjectRepository() *ProjectRepository {
	return &ProjectRepository{store: NewStore()}
}

// GetAll returns all projects.
func (r *ProjectRepository) GetAll() ([]model.Project, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	projects := []model.Project{}
	for _, p := range r.store.projects {
		projects = append(projects, p)
	}

	return projects, nil
}

// Create creates and returns a new project.
func (r *ProjectRepository) Create(p model.Project) (model.Project, error) {
	r.m.Lock()
	defer r.m.Unlock()

	p.ID = len(r.store.projects) + 1
	r.store.projects[p.ID] = p

	return p, nil
}

// GetByID returns a project with specific ID.
func (r *ProjectRepository) GetByID(id int) (model.Project, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	if p, ok := r.store.projects[id]; ok {
		return p, nil
	}

	return model.Project{}, store.ErrNotFound
}

// Update updates a project.
func (r *ProjectRepository) Update(p model.Project) error {
	r.m.Lock()
	defer r.m.Unlock()

	if _, ok := r.store.projects[p.ID]; ok {
		r.store.projects[p.ID] = p
		return nil
	}

	return store.ErrNotFound
}

// Delete deletes a project.
func (r *ProjectRepository) Delete(p model.Project) error {
	r.m.Lock()
	defer r.m.Unlock()

	if _, ok := r.store.projects[p.ID]; ok {
		delete(r.store.projects, p.ID)
		return nil
	}

	return store.ErrNotFound
}
