package repositories

import (
	"context"

	"github.com/DENFNC/ToDoList/internal/domain/models"
)

type TaskRepository interface {
	Create(
		ctx context.Context,
		task *models.Task,
	) (uint32, error)
	GetByID(
		ctx context.Context,
		id uint32,
	) (*models.Task, error)
	Update(
		ctx context.Context,
		task *models.Task,
	) (uint32, error)
	Delete(
		ctx context.Context,
		id uint32,
	) (uint32, error)
}
