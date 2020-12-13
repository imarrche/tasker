package memory

import (
	"sync"

	"github.com/imarrche/tasker/internal/model"
	"github.com/imarrche/tasker/internal/store"
)

// ProjectRepository is an in memory project repository.
type ProjectRepository struct {
	db *inMemoryDb
	m  sync.RWMutex
}

// NewProjectRepository creates and returns a new ProjectRepository instance.
func NewProjectRepository(db *inMemoryDb) *ProjectRepository {
	return &ProjectRepository{db: db}
}

// GetAll returns all projects.
func (r *ProjectRepository) GetAll() ([]model.Project, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	projects := []model.Project{}
	for _, p := range r.db.projects {
		projects = append(projects, p)
	}

	return projects, nil
}

// Create creates and returns a new project.
func (r *ProjectRepository) Create(p model.Project) (model.Project, error) {
	r.m.Lock()
	defer r.m.Unlock()

	p.ID = len(r.db.projects) + 1
	r.db.projects[p.ID] = p

	return p, nil
}

// GetByID returns a project with specific ID.
func (r *ProjectRepository) GetByID(id int) (model.Project, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	if p, ok := r.db.projects[id]; ok {
		return p, nil
	}

	return model.Project{}, store.ErrNotFound
}

// Update updates a project.
func (r *ProjectRepository) Update(p model.Project) error {
	r.m.Lock()
	defer r.m.Unlock()

	if _, ok := r.db.projects[p.ID]; ok {
		r.db.projects[p.ID] = p
		return nil
	}

	return store.ErrNotFound
}

// DeleteByID deletes a project with specific ID.
func (r *ProjectRepository) DeleteByID(id int) error {
	r.m.Lock()
	defer r.m.Unlock()

	if _, ok := r.db.projects[id]; ok {
		delete(r.db.projects, id)
		return nil
	}

	return store.ErrNotFound
}
