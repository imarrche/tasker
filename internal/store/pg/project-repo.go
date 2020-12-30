package pg

import (
	"database/sql"

	"github.com/imarrche/tasker/internal/store"

	"github.com/imarrche/tasker/internal/model"
)

// projectRepo is the project repository for PostgreSQL store.
type projectRepo struct {
	db *sql.DB
}

// newProjectRepo creates and returns a new projectRepo instance.
func newProjectRepo(db *sql.DB) *projectRepo {
	return &projectRepo{db: db}
}

// GetAll returns all projects.
func (r *projectRepo) GetAll() (ps []model.Project, err error) {
	rows, err := r.db.Query("SELECT * FROM projects;")
	if err != nil {
		return []model.Project{}, err
	}
	defer rows.Close()

	var p model.Project
	for rows.Next() {
		err := rows.Scan(&p.ID, &p.Name, &p.Description)
		if err != nil {
			return []model.Project{}, err
		}

		ps = append(ps, p)
	}
	err = rows.Err()
	if err != nil {
		return []model.Project{}, err
	}

	return ps, nil
}

// Create creates and returns a new project.
func (r *projectRepo) Create(p model.Project) (model.Project, error) {
	var id int
	query := "INSERT INTO projects (name, description) VALUES ($1, $2) RETURNING id;"

	row := r.db.QueryRow(query, p.Name, p.Description)
	if err := row.Scan(&id); err != nil {
		return model.Project{}, err
	}

	p.ID = id
	return p, nil
}

// GetByID returns the project with specific ID.
func (r *projectRepo) GetByID(id int) (model.Project, error) {
	var p model.Project
	query := "SELECT * FROM projects WHERE id = $1;"

	row := r.db.QueryRow(query, id)
	err := row.Scan(&p.ID, &p.Name, &p.Description)
	if err == sql.ErrNoRows {
		return model.Project{}, store.ErrNotFound
	}
	if err != nil {
		return model.Project{}, err
	}

	return p, nil
}

// Update updates the project.
func (r *projectRepo) Update(p model.Project) (model.Project, error) {
	query := "UPDATE projects SET name = $1, description = $2 WHERE id = $3;"

	_, err := r.db.Exec(query, p.Name, p.Description, p.ID)
	return p, err
}

// DeleteByID deletes the project with specific ID.
func (r *projectRepo) DeleteByID(id int) error {
	query := "DELETE FROM projects WHERE id = $1;"

	_, err := r.db.Exec(query, id)
	return err
}
