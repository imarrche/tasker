package pg

import (
	"database/sql"

	"github.com/imarrche/tasker/internal/model"
	"github.com/imarrche/tasker/internal/store"
)

// commentRepo is the comment repository for PostgreSQL store.
type commentRepo struct {
	db *sql.DB
}

// newCommentRepo creates and returns a new commentRepo instance.
func newCommentRepo(db *sql.DB) *commentRepo { return &commentRepo{db: db} }

// GetByTaskID returns all comments with specific task ID.
func (r *commentRepo) GetByTaskID(id int) ([]model.Comment, error) {
	rows, err := r.db.Query("SELECT * FROM tasks WHERE id = $1;", id)
	if err != nil {
		return nil, err
	} else if !rows.Next() {
		return nil, store.ErrNotFound
	}

	rows, err = r.db.Query("SELECT * FROM comments WHERE task_id = $1;", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cs, c := []model.Comment{}, model.Comment{}
	for rows.Next() {
		if err := rows.Scan(&c.ID, &c.Text, &c.CreatedAt, &c.TaskID); err != nil {
			return nil, err
		}
		cs = append(cs, c)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return cs, nil
}

// Create creates and returns a new comment.
func (r *commentRepo) Create(c model.Comment) (model.Comment, error) {
	query := "INSERT INTO comments (text, created_at, task_id) VALUES ($1, $2, $3) RETURNING id;"
	row := r.db.QueryRow(query, c.Text, c.CreatedAt, c.TaskID)

	var id int
	if err := row.Scan(&id); err != nil {
		return model.Comment{}, err
	}
	c.ID = id

	return c, nil
}

// GetByID returns the comment with specific ID.
func (r *commentRepo) GetByID(id int) (model.Comment, error) {
	row := r.db.QueryRow("SELECT * FROM comments WHERE id = $1;", id)

	var c model.Comment
	err := row.Scan(&c.ID, &c.Text, &c.CreatedAt, &c.TaskID)
	if err == sql.ErrNoRows {
		return model.Comment{}, store.ErrNotFound
	} else if err != nil {
		return model.Comment{}, err
	}

	return c, nil
}

// Update updates the comment.
func (r *commentRepo) Update(c model.Comment) (model.Comment, error) {
	query := "UPDATE comments SET text = $1, created_at = $2, task_id = $3 WHERE id = $4;"
	res, err := r.db.Exec(query, c.Text, c.CreatedAt, c.TaskID, c.ID)

	if err != nil {
		return model.Comment{}, err
	}
	rowsCount, err := res.RowsAffected()
	if err != nil {
		return model.Comment{}, err
	} else if rowsCount == 0 {
		return model.Comment{}, store.ErrNotFound
	}

	return c, err
}

// DeleteByID deletes the comment with specific ID.
func (r *commentRepo) DeleteByID(id int) error {
	res, err := r.db.Exec("DELETE FROM comments WHERE id = $1;", id)

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
