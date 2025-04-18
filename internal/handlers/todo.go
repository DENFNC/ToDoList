package handlers

import (
	"github.com/DENFNC/ToDoList/internal/domain/models"
	"github.com/gin-gonic/gin"
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

func (h *TodoHandler) Register(rg *gin.RouterGroup) {
	todoRoute := rg.Group("/api/v1/tasks")
	{
		todoRoute.POST("/:")
		todoRoute.GET("/:id")
		todoRoute.PUT("/:id")
		todoRoute.DELETE("/:id")
	}
}

func (h *TodoHandler) CreateTask(ctx *gin.Context) { panic("implement me") }
func (h *TodoHandler) GetTask(ctx *gin.Context)    { panic("implement me") }
func (h *TodoHandler) UpdateTask(ctx *gin.Context) { panic("implement me") }
func (h *TodoHandler) DeleteTask(ctx *gin.Context) { panic("implement me") }
