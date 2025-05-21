package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/mlucas4330/todo-go-cli/internal/model"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) FindAll() ([]model.Task, error) {
	rows, err := r.db.Query(`
		SELECT id, title, description, amount, category, start_date, end_date, notification_date
		FROM tasks
		ORDER BY start_date ASC, id ASC -- Add an ORDER BY for consistent results
	`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task

	for rows.Next() {
		var task model.Task

		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Amount,
			&task.Category,
			&task.StartDate,
			&task.EndDate,
			&task.NotificationDate,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *TaskRepository) FindById(id string) (model.Task, error) {
	var task model.Task

	row := r.db.QueryRow(`
		SELECT id, title, description, amount, category, start_date, end_date, notification_date
		FROM tasks
		WHERE ID = $1
	`, id)

	err := row.Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Amount,
		&task.Category,
		&task.StartDate,
		&task.EndDate,
		&task.NotificationDate,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return model.Task{}, err
		}
		return model.Task{}, err
	}

	return task, nil
}

func (r *TaskRepository) Create(task *model.Task) error {
	_, err := r.db.Exec(`
		INSERT INTO tasks (title, category, description, amount, start_date, end_date, notification_date)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`,
		task.Title,
		task.Category,
		task.Description,
		task.Amount,
		task.StartDate,
		task.EndDate,
		task.NotificationDate,
	)

	return err
}

func (r *TaskRepository) Update(id string, setClauses []string, values []any) error {
	var assignments []string
	for i, col := range setClauses {
		assignments = append(assignments, fmt.Sprintf("%s = $%d", col, i+1))
	}

	idAssigment := len(values) + 1

	query := fmt.Sprintf("UPDATE tasks SET %s WHERE id = $%d", strings.Join(assignments, ", "), idAssigment)

	values = append(values, id)

	_, err := r.db.Exec(query, values...)

	return err
}

func (r *TaskRepository) Delete(id string) error {
	result, err := r.db.Exec("DELETE FROM tasks WHERE ID = $1", id)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no task found with ID %s", id)
	}

	return nil
}
