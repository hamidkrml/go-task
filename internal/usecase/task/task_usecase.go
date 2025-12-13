package task

import (
	"errors"
	"go-task-manager/internal/domain"
)

type taskUseCase struct {
	repo domain.TaskRepository
}

func NewTaskUseCase(repo domain.TaskRepository) domain.TaskUseCase {
	return &taskUseCase{repo: repo}
}

func (u *taskUseCase) CreateTask(userID int, title, description string) error {
	task := &domain.Task{
		UserID:      userID,
		Title:       title,
		Description: description,
		Status:      "pending", // Varsayılan durum
	}
	return u.repo.Create(task)
}

func (u *taskUseCase) GetUserTasks(userID int) ([]domain.Task, error) {
	return u.repo.GetByUserID(userID)
}

func (u *taskUseCase) UpdateTaskStatus(id, userID int, status string) error {
	task, err := u.repo.GetByID(id)
	if err != nil {
		return err
	}

	// Başkasının görevini güncellemeye çalışıyorsa hata ver
	if task.UserID != userID {
		return errors.New("unauthorized")
	}

	// Basit validasyon
	if status != "pending" && status != "in_progress" && status != "done" {
		return errors.New("invalid status")
	}

	task.Status = status
	return u.repo.Update(task)
}

func (u *taskUseCase) DeleteTask(id, userID int) error {
	task, err := u.repo.GetByID(id)
	if err != nil {
		return err
	}

	// Başkasının görevini silmeye çalışıyorsa hata ver
	if task.UserID != userID {
		return errors.New("unauthorized")
	}

	return u.repo.Delete(id)
}
