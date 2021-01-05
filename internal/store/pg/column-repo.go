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
func newColumnRepo(db *sql.DB) *columnRepo { return &columnRepo{db: db} }

// GetByProjectID returns all columns with specific project ID.
func (r *columnRepo) GetByProjectID(id int) ([]model.Column, error) {
	rows, err := r.db.Query("SELECT * FROM projects WHERE id = $1;", id)
	if err != nil {
		return nil, err
	} else if !rows.Next() {
		return nil, store.ErrNotFound
	}

	rows, err = r.db.Query("SELECT * FROM columns WHERE project_id = $1;", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cs, c := []model.Column{}, model.Column{}
	for rows.Next() {
		if err = rows.Scan(&c.ID, &c.Name, &c.Index, &c.ProjectID); err != nil {
			return nil, err
		}
		cs = append(cs, c)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return cs, nil
}

// Create creates and returns a new column.
func (r *columnRepo) Create(c model.Column) (model.Column, error) {
	query := "INSERT INTO columns (name, index, project_id) VALUES ($1, $2, $3) RETURNING id;"
	row := r.db.QueryRow(query, c.Name, c.Index, c.ProjectID)

	var id int
	if err := row.Scan(&id); err != nil {
		return model.Column{}, err
	}
	c.ID = id

	return c, nil
}

// GetByID returns the column with specifc ID.
func (r *columnRepo) GetByID(id int) (model.Column, error) {
	row := r.db.QueryRow("SELECT * FROM columns WHERE id = $1;", id)

	var c model.Column
	err := row.Scan(&c.ID, &c.Name, &c.Index, &c.ProjectID)
	if err == sql.ErrNoRows {
		return model.Column{}, store.ErrNotFound
	} else if err != nil {
		return model.Column{}, err
	}

	return c, nil
}

// GetByIndexAndProjectID returns the column with specific index and project ID.
func (r *columnRepo) GetByIndexAndProjectID(index, id int) (model.Column, error) {
	query := "SELECT * FROM columns WHERE index = $1 AND project_id = $2;"
	row := r.db.QueryRow(query, index, id)

	var c model.Column
	err := row.Scan(&c.ID, &c.Name, &c.Index, &c.ProjectID)
	if err == sql.ErrNoRows {
		return model.Column{}, store.ErrNotFound
	} else if err != nil {
		return model.Column{}, err
	}

	return c, nil
}

// Update updates the column.
func (r *columnRepo) Update(c model.Column) (model.Column, error) {
	query := "UPDATE columns SET name = $1, index = $2, project_id = $3 WHERE id = $4;"
	res, err := r.db.Exec(query, c.Name, c.Index, c.ProjectID, c.ID)

	if err != nil {
		return model.Column{}, err
	}
	rowsCount, err := res.RowsAffected()
	if err != nil {
		return model.Column{}, err
	} else if rowsCount == 0 {
		return model.Column{}, store.ErrNotFound
	}

	return c, nil
}

// DeleteByID deletes the column with specific ID.
func (r *columnRepo) DeleteByID(id int) error {
	res, err := r.db.Exec("DELETE FROM columns WHERE id = $1;", id)

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
