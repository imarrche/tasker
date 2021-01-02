package inmem

import (
	"sync"

	"github.com/imarrche/tasker/internal/model"
	"github.com/imarrche/tasker/internal/store"
)

// projectRepo is the project repository for in memory store.
type projectRepo struct {
	db *inMemoryDb
	m  sync.RWMutex
}

// newProjectRepo creates and returns a new projectRepo instance.
func newProjectRepo(db *inMemoryDb) *projectRepo { return &projectRepo{db: db} }

// GetAll returns all projects.
func (r *projectRepo) GetAll() ([]model.Project, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	ps := []model.Project{}
	for _, p := range r.db.projects {
		ps = append(ps, p)
	}

	return ps, nil
}

// Create creates and returns a new project.
func (r *projectRepo) Create(p model.Project) (model.Project, error) {
	r.m.Lock()
	defer r.m.Unlock()

	p.ID = len(r.db.projects) + 1
	r.db.projects[p.ID] = p

	return p, nil
}

// GetByID returns the project with specific ID.
func (r *projectRepo) GetByID(id int) (model.Project, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	if p, ok := r.db.projects[id]; ok {
		return p, nil
	}

	return model.Project{}, store.ErrNotFound
}

// Update updates the project.
func (r *projectRepo) Update(p model.Project) (model.Project, error) {
	r.m.Lock()
	defer r.m.Unlock()

	if _, ok := r.db.projects[p.ID]; !ok {
		return model.Project{}, store.ErrNotFound
	}

	r.db.projects[p.ID] = p

	return p, nil
}

// DeleteByID deletes the project with specific ID.
func (r *projectRepo) DeleteByID(id int) error {
	r.m.Lock()
	defer r.m.Unlock()

	if _, ok := r.db.projects[id]; !ok {
		return store.ErrNotFound
	}

	for columnID, column := range r.db.columns {
		if column.ProjectID == id {
			for taskID, task := range r.db.tasks {
				if task.ColumnID == columnID {
					for commentID, comment := range r.db.comments {
						if comment.TaskID == taskID {
							delete(r.db.comments, commentID)
						}
					}
					delete(r.db.tasks, taskID)
				}
			}
			delete(r.db.columns, columnID)
		}
	}
	delete(r.db.projects, id)

	return nil
}
