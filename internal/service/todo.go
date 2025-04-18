package service

import (
	"github.com/DENFNC/ToDoList/internal/domain/repositories"
)

type TaskService struct {
	repo repositories.TaskRepository
}

func NewTaskService(
	repo repositories.TaskRepository,
) *TaskService {
	return &TaskService{
		repo: repo,
	}
}
