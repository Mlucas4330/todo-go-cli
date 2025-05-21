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
		return nil, fmt.Errorf("failed to query tasks: %w", err)
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
			return nil, fmt.Errorf("failed to scan task row: %w", err)
		}
		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %w", err)
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
			return model.Task{}, fmt.Errorf("task with ID %s not found: %w", id, err)
		}
		return model.Task{}, fmt.Errorf("failed to scan task for ID %s: %w", id, err)
	}

	return task, nil
}

func (r *TaskRepository) Create(task *model.Task) error {
	_, err := r.db.Exec(`
		INSERT INTO tasks (id, title, category, description, amount, start_date, end_date, notification_date)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`,
		task.ID,
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
	idPlaceholderNum := len(values) + 1
	query := fmt.Sprintf("UPDATE tasks SET %s WHERE id = $%d", strings.Join(setClauses, ", "), idPlaceholderNum)
	values = append(values, id)

	_, err := r.db.Exec(query, values...)

	return err
}

func (r *TaskRepository) Delete(id string) error {
	_, err := r.db.Exec("DELETE FROM tasks WHERE ID = $1", id)

	return err
}
