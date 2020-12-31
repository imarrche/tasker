package pg

import (
	"database/sql"

	"github.com/imarrche/tasker/internal/store"

	"github.com/imarrche/tasker/internal/model"
)

// commentRepo is the comment repository for PostgreSQL store.
type commentRepo struct {
	db *sql.DB
}

// newCommentRepo creates and returns a new commentRepo instance.
func newCommentRepo(db *sql.DB) *commentRepo {
	return &commentRepo{db: db}
}

// GetByTaskID returns all comments with specific task ID.
func (r *commentRepo) GetByTaskID(id int) ([]model.Comment, error) {
	rows, err := r.db.Query("SELECT * FROM comments WHERE task_id = $1;", id)
	if err != nil {
		return []model.Comment{}, err
	}
	defer rows.Close()

	cs := []model.Comment{}
	var c model.Comment
	for rows.Next() {
		err := rows.Scan(&c.ID, &c.Text, &c.CreatedAt, &c.TaskID)
		if err != nil {
			return []model.Comment{}, err
		}

		cs = append(cs, c)
	}
	err = rows.Err()
	if err != nil {
		return []model.Comment{}, err
	}

	return cs, nil
}

// Create creates and returns a new comment.
func (r *commentRepo) Create(c model.Comment) (model.Comment, error) {
	var id int
	query := "INSERT INTO comments (text, created_at, task_id) VALUES ($1, $2, $3) RETURNING id;"

	row := r.db.QueryRow(query, c.Text, c.CreatedAt, c.TaskID)
	if err := row.Scan(&id); err != nil {
		return model.Comment{}, err
	}

	c.ID = id
	return c, nil
}

// GetByID returns the comment with specific ID.
func (r *commentRepo) GetByID(id int) (model.Comment, error) {
	var c model.Comment
	query := "SELECT * FROM comments WHERE id = $1;"

	row := r.db.QueryRow(query, id)
	err := row.Scan(&c.ID, &c.Text, &c.CreatedAt, &c.TaskID)
	if err == sql.ErrNoRows {
		return model.Comment{}, store.ErrNotFound
	}
	if err != nil {
		return model.Comment{}, err
	}

	return c, nil
}

// Update updates the comment.
func (r *commentRepo) Update(c model.Comment) (model.Comment, error) {
	query := "UPDATE comments SET text = $1 WHERE id = $2;"

	_, err := r.db.Exec(query, c.Text, c.ID)
	return c, err
}

// DeleteByID deletes the comment with specific ID.
func (r *commentRepo) DeleteByID(id int) error {
	query := "DELETE FROM comments WHERE id = $1;"

	_, err := r.db.Exec(query, id)
	return err
}
