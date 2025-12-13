package task

import (
	"database/sql"
	"go-task-manager/internal/domain"
)

type taskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) domain.TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) Create(task *domain.Task) error {
	query := `INSERT INTO tasks (user_id, title, description, status) VALUES ($1, $2, $3, $4) RETURNING id, created_at`
	return r.db.QueryRow(query, task.UserID, task.Title, task.Description, task.Status).Scan(&task.ID, &task.CreatedAt)
}

func (r *taskRepository) GetByID(id int) (*domain.Task, error) {
	query := `SELECT id, user_id, title, description, status, created_at FROM tasks WHERE id = $1`
	var task domain.Task
	err := r.db.QueryRow(query, id).Scan(&task.ID, &task.UserID, &task.Title, &task.Description, &task.Status, &task.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *taskRepository) GetByUserID(userID int) ([]domain.Task, error) {
	query := `SELECT id, user_id, title, description, status, created_at FROM tasks WHERE user_id = $1`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []domain.Task
	for rows.Next() {
		var task domain.Task
		if err := rows.Scan(&task.ID, &task.UserID, &task.Title, &task.Description, &task.Status, &task.CreatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *taskRepository) Update(task *domain.Task) error {
	query := `UPDATE tasks SET status = $1 WHERE id = $2`
	_, err := r.db.Exec(query, task.Status, task.ID)
	return err
}

func (r *taskRepository) Delete(id int) error {
	query := `DELETE FROM tasks WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
