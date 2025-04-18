package handlers

import (
	"github.com/DENFNC/ToDoList/internal/domain/models"
	"github.com/gofiber/fiber/v3"
)

type TodoService interface {
	CreateTask(
		title string,
		description string,
	) (uint32, error)
	GetTask(
		taskID uint32,
	) (models.Task, error)
	UpdateTask(
		title string,
		description string,
		status string,
	) (uint32, error)
	DeleteTask(
		taskID uint32,
	) (uint32, error)
}

type TodoHandler struct {
	svc TodoService
}

func NewTodoHandler(service TodoService) *TodoHandler {
	return &TodoHandler{
		svc: service,
	}
}

func (h *TodoHandler) Register(rg *fiber.Group) {
	todoRoute := rg.Group("/api/v1/tasks")
	{
		todoRoute.Post("/:id", h.CreateTask)
		todoRoute.Get("/:id", h.GetTask)
		todoRoute.Put("/:id", h.UpdateTask)
		todoRoute.Delete("/:id", h.DeleteTask)
	}
}

func (h *TodoHandler) CreateTask(ctx fiber.Ctx) error { panic("implement me") }
func (h *TodoHandler) GetTask(ctx fiber.Ctx) error    { panic("implement me") }
func (h *TodoHandler) UpdateTask(ctx fiber.Ctx) error { panic("implement me") }
func (h *TodoHandler) DeleteTask(ctx fiber.Ctx) error { panic("implement me") }
