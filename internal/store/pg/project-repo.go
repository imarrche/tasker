package pg

import (
	"database/sql"

	"github.com/imarrche/tasker/internal/model"
	"github.com/imarrche/tasker/internal/store"
)

// projectRepo is the project repository for PostgreSQL store.
type projectRepo struct {
	db *sql.DB
}

// newProjectRepo creates and returns a new projectRepo instance.
func newProjectRepo(db *sql.DB) *projectRepo { return &projectRepo{db: db} }

// GetAll returns all projects.
func (r *projectRepo) GetAll() ([]model.Project, error) {
	rows, err := r.db.Query("SELECT * FROM projects;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ps, p := []model.Project{}, model.Project{}
	for rows.Next() {
		if err = rows.Scan(&p.ID, &p.Name, &p.Description); err != nil {
			return nil, err
		}
		ps = append(ps, p)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return ps, nil
}

// Create creates and returns a new project.
func (r *projectRepo) Create(p model.Project) (model.Project, error) {
	query := "INSERT INTO projects (name, description) VALUES ($1, $2) RETURNING id;"
	row := r.db.QueryRow(query, p.Name, p.Description)

	var id int
	if err := row.Scan(&id); err != nil {
		return model.Project{}, err
	}
	p.ID = id

	return p, nil
}

// GetByID returns the project with specific ID.
func (r *projectRepo) GetByID(id int) (model.Project, error) {
	row := r.db.QueryRow("SELECT * FROM projects WHERE id = $1;", id)

	var p model.Project
	err := row.Scan(&p.ID, &p.Name, &p.Description)
	if err == sql.ErrNoRows {
		return model.Project{}, store.ErrNotFound
	} else if err != nil {
		return model.Project{}, err
	}

	return p, nil
}

// Update updates the project.
func (r *projectRepo) Update(p model.Project) (model.Project, error) {
	query := "UPDATE projects SET name = $1, description = $2 WHERE id = $3;"
	res, err := r.db.Exec(query, p.Name, p.Description, p.ID)

	if err != nil {
		return model.Project{}, err
	}
	rowsCount, err := res.RowsAffected()
	if err != nil {
		return model.Project{}, err
	} else if rowsCount == 0 {
		return model.Project{}, store.ErrNotFound
	}

	return p, nil
}

// DeleteByID deletes the project with specific ID.
func (r *projectRepo) DeleteByID(id int) error {
	res, err := r.db.Exec("DELETE FROM projects WHERE id = $1;", id)

	if err != nil {
		return err
	}
	rowsCount, err := res.RowsAffected()
	if err != nil {
		return err
	} else if rowsCount == 0 {
		return store.ErrNotFound
	}

	return nil
}
