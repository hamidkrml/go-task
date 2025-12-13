package domain

import "time"

type Task struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"` // "pending", "in_progress", "done"
	CreatedAt   time.Time `json:"created_at"`
}

type TaskRepository interface {
	Create(task *Task) error
	GetByID(id int) (*Task, error)
	GetByUserID(userID int) ([]Task, error)
	Update(task *Task) error
	Delete(id int) error
}

type TaskUseCase interface {
	CreateTask(userID int, title, description string) error
	GetUserTasks(userID int) ([]Task, error)
	UpdateTaskStatus(id, userID int, status string) error
	DeleteTask(id, userID int) error
}
