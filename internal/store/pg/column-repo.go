package pg

import (
	"database/sql"

	"github.com/imarrche/tasker/internal/model"
	"github.com/imarrche/tasker/internal/store"
)

// columnRepo is the column repository for PostgreSQL store.
type columnRepo struct {
	db *sql.DB
}

// newColumnRepo creates and returns a new columnRepo instance.
func newColumnRepo(db *sql.DB) *columnRepo {
	return &columnRepo{db: db}
}

// GetByProjectID returns all columns with specific project ID.
func (r *columnRepo) GetByProjectID(id int) (cs []model.Column, err error) {
	rows, err := r.db.Query("SELECT * FROM columns WHERE project_id = $1;", id)
	if err != nil {
		return []model.Column{}, err
	}
	defer rows.Close()

	var c model.Column
	for rows.Next() {
		err := rows.Scan(&c.ID, &c.Name, &c.Index, &c.ProjectID)
		if err != nil {
			return []model.Column{}, err
		}

		cs = append(cs, c)
	}
	err = rows.Err()
	if err != nil {
		return []model.Column{}, err
	}

	return cs, nil
}

// Create creates and returns a new column.
func (r *columnRepo) Create(c model.Column) (model.Column, error) {
	var id int
	query := "INSERT INTO columns (name, index, project_id) VALUES ($1, $2, $3) RETURNING id;"

	row := r.db.QueryRow(query, c.Name, c.Index, c.ProjectID)
	if err := row.Scan(&id); err != nil {
		return model.Column{}, err
	}

	c.ID = id
	return c, nil
}

// GetByID returns the column with specifc ID.
func (r *columnRepo) GetByID(id int) (model.Column, error) {
	var c model.Column
	query := "SELECT * FROM columns WHERE id = $1;"

	row := r.db.QueryRow(query, id)
	err := row.Scan(&c.ID, &c.Name, &c.Index, &c.ProjectID)
	if err == sql.ErrNoRows {
		return model.Column{}, store.ErrNotFound
	}
	if err != nil {
		return model.Column{}, err
	}

	return c, nil
}

// GetByIndexAndProjectID returns the column with specific index and project ID.
func (r *columnRepo) GetByIndexAndProjectID(index, id int) (model.Column, error) {
	var c model.Column
	query := "SELECT * FROM columns WHERE index = $1 AND project_id = $2;"

	row := r.db.QueryRow(query, index, id)
	if err := row.Scan(&c.ID, &c.Name, &c.Index, &c.ProjectID); err != nil {
		return model.Column{}, err
	}

	return c, nil
}

// Update updates the column.
func (r *columnRepo) Update(c model.Column) (model.Column, error) {
	query := "UPDATE columns SET name = $1, index = $2 WHERE id = $3;"

	_, err := r.db.Exec(query, c.Name, c.Index, c.ID)
	return c, err
}

// DeleteByID deletes the column with specific ID.
func (r *columnRepo) DeleteByID(id int) error {
	query := "DELETE FROM columns WHERE id = $1;"

	_, err := r.db.Exec(query, id)
	return err
}
