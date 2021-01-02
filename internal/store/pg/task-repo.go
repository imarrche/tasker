package pg

import (
	"database/sql"

	"github.com/imarrche/tasker/internal/model"
	"github.com/imarrche/tasker/internal/store"
)

// taskRepo is the task repository for PostgreSQL store.
type taskRepo struct {
	db *sql.DB
}

// newTaskRepo creates and returns a new taskRepo instance.
func newTaskRepo(db *sql.DB) *taskRepo { return &taskRepo{db: db} }

// GetByColumnID returns all tasks with specific column ID.
func (r *taskRepo) GetByColumnID(id int) ([]model.Task, error) {
	rows, err := r.db.Query("SELECT * FROM columns WHERE id = $1;", id)
	if err != nil {
		return nil, err
	} else if !rows.Next() {
		return nil, store.ErrNotFound
	}

	rows, err = r.db.Query("SELECT * FROM tasks WHERE column_id = $1;", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ts, t := []model.Task{}, model.Task{}
	for rows.Next() {
		if err = rows.Scan(&t.ID, &t.Name, &t.Description, &t.Index, &t.ColumnID); err != nil {
			return nil, err
		}
		ts = append(ts, t)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return ts, nil
}

// Create creates and returns a new task.
func (r *taskRepo) Create(t model.Task) (model.Task, error) {
	query := "INSERT INTO tasks (name, description, index, column_id) VALUES ($1, $2, $3, $4) RETURNING id;"
	row := r.db.QueryRow(query, t.Name, t.Description, t.Index, t.ColumnID)

	var id int
	if err := row.Scan(&id); err != nil {
		return model.Task{}, err
	}
	t.ID = id

	return t, nil
}

// GetByID returns the task with specifc ID.
func (r *taskRepo) GetByID(id int) (model.Task, error) {
	row := r.db.QueryRow("SELECT * FROM tasks WHERE id = $1;", id)

	var t model.Task
	err := row.Scan(&t.ID, &t.Name, &t.Description, &t.Index, &t.ColumnID)
	if err == sql.ErrNoRows {
		return model.Task{}, store.ErrNotFound
	} else if err != nil {
		return model.Task{}, err
	}

	return t, nil
}

// GetByIndexAndColumnID returns the task with specific index and column ID.
func (r *taskRepo) GetByIndexAndColumnID(index, id int) (model.Task, error) {
	row := r.db.QueryRow("SELECT * FROM tasks WHERE index = $1 AND column_id = $2;", index, id)

	var t model.Task
	err := row.Scan(&t.ID, &t.Name, &t.Description, &t.Index, &t.ColumnID)
	if err == sql.ErrNoRows {
		return model.Task{}, store.ErrNotFound
	} else if err != nil {
		return model.Task{}, err
	}

	return t, nil
}

// Update updates the tasks.
func (r *taskRepo) Update(t model.Task) (model.Task, error) {
	query := "UPDATE tasks SET name = $1, description = $2, index = $3, column_id = $4 WHERE id = $5;"
	res, err := r.db.Exec(query, t.Name, t.Description, t.Index, t.ColumnID, t.ID)

	if err != nil {
		return model.Task{}, err
	}
	rowsCount, err := res.RowsAffected()
	if err != nil {
		return model.Task{}, err
	} else if rowsCount == 0 {
		return model.Task{}, store.ErrNotFound
	}

	return t, nil
}

// DeleteByID deletes the task with specific ID.
func (r *taskRepo) DeleteByID(id int) error {
	res, err := r.db.Exec("DELETE FROM tasks WHERE id = $1;", id)

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
