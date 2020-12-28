package pg

import (
	"database/sql"

	"github.com/imarrche/tasker/internal/model"
)

// taskRepo is the task repository for PostgreSQL store.
type taskRepo struct {
	db *sql.DB
}

// newTaskRepo creates and returns a new taskRepo instance.
func newTaskRepo(db *sql.DB) *taskRepo {
	return &taskRepo{db: db}
}

// GetByColumnID returns all tasks with specific column ID.
func (r *taskRepo) GetByColumnID(id int) (ts []model.Task, err error) {
	rows, err := r.db.Query("SELECT * FROM tasks WHERE column_id = $1;", id)
	if err != nil {
		return []model.Task{}, err
	}
	defer rows.Close()

	var t model.Task
	for rows.Next() {
		err := rows.Scan(&t.ID, &t.Name, &t.Description, &t.Index, &t.ColumnID)
		if err != nil {
			return []model.Task{}, err
		}

		ts = append(ts, t)
	}
	err = rows.Err()
	if err != nil {
		return []model.Task{}, err
	}

	return ts, nil
}

// Create creates and returns a new task.
func (r *taskRepo) Create(t model.Task) (model.Task, error) {
	var id int
	query := "INSERT INTO tasks (name, description, index, column_id) VALUES ($1, $2, $3, $4) RETURNING id;"

	row := r.db.QueryRow(query, t.Name, t.Description, t.Index, t.ColumnID)
	if err := row.Scan(&id); err != nil {
		return model.Task{}, err
	}

	t.ID = id
	return t, nil
}

// GetByID returns the task with specifc ID.
func (r *taskRepo) GetByID(id int) (model.Task, error) {
	var t model.Task
	query := "SELECT FROM tasks WHERE id = $1;"

	row := r.db.QueryRow(query, id)
	if err := row.Scan(&t.ID, &t.Name, &t.Description, &t.Index, &t.ColumnID); err != nil {
		return model.Task{}, err
	}

	return t, nil
}

// GetByIndexAndColumnID returns the task with specific index and column ID.
func (r *taskRepo) GetByIndexAndColumnID(index, id int) (model.Task, error) {
	var t model.Task
	query := "SELECT FROM tasks WHERE index = $1 AND column_id = $2;"

	row := r.db.QueryRow(query, index, id)
	if err := row.Scan(&t.ID, &t.Name, &t.Description, &t.Index, &t.ColumnID); err != nil {
		return model.Task{}, err
	}

	return t, nil
}

// Update updates the tasks.
func (r *taskRepo) Update(t model.Task) (model.Task, error) {
	query := "UPDATE tasks SET name = $1, description = $2, index = $3 WHERE id = $4;"

	_, err := r.db.Exec(query, t.Name, t.Description, t.Index, t.ID)
	return t, err
}

// DeleteByID deletes the task with specific ID.
func (r *taskRepo) DeleteByID(id int) error {
	query := "DELETE FROM tasks WHERE id = $1;"

	_, err := r.db.Exec(query, id)
	return err
}
