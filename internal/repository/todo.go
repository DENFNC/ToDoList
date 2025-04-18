package repository

import (
	"context"
	"time"

	"github.com/DENFNC/ToDoList/internal/domain/models"
	psql "github.com/DENFNC/ToDoList/internal/pkg/storage/postgre"
	"github.com/doug-martin/goqu/v9"
)

type TaskRepository struct {
	*psql.Storage
	goqu *goqu.DialectWrapper
}

func NewTaskRepository(
	db *psql.Storage,
	goqu *goqu.DialectWrapper,
) *TaskRepository {
	return &TaskRepository{
		db,
		goqu,
	}
}

func (r *TaskRepository) Create(
	ctx context.Context,
	task *models.Task,
) (uint32, error) {
	stmt, args, err := r.goqu.Insert("tasks").Returning("id").Rows(
		goqu.Record{
			"title":       task.Title,
			"description": task.Description,
			"status":      task.Status,
		},
	).Prepared(true).ToSQL()
	if err != nil {
		return 0, err
	}

	var taskID uint32
	if err := r.DB.QueryRow(
		ctx,
		stmt,
		args...,
	).Scan(&taskID); err != nil {
		return 0, err
	}

	return taskID, nil
}

func (r *TaskRepository) GetByID(
	ctx context.Context,
	id uint32,
) (*models.Task, error) {
	stmt, args, err := r.goqu.Select(
		"id",
		"title",
		"description",
		"status",
		"created_at",
		"updated_at",
	).From("tasks").Where(
		goqu.C("id").Eq(id),
	).Prepared(true).ToSQL()
	if err != nil {
		return nil, err
	}

	var task *models.Task
	if err := r.DB.QueryRow(
		ctx,
		stmt,
		args...,
	).Scan(&task); err != nil {
		return nil, err
	}

	return task, nil
}

func (r *TaskRepository) Update(
	ctx context.Context,
	task *models.Task,
) (uint32, error) {
	stmt, args, err := r.goqu.Update("tasks").Returning("id").Set(
		goqu.Record{
			"title":       task.Title,
			"description": task.Description,
			"status":      task.Status,
			"updated_at":  time.Now(),
		},
	).Prepared(true).ToSQL()
	if err != nil {
		return 0, err
	}

	var taskID uint32
	if err := r.DB.QueryRow(ctx, stmt, args...).Scan(&taskID); err != nil {
		return 0, err
	}

	return taskID, nil
}

func (r *TaskRepository) Delete(
	ctx context.Context,
	id uint32,
) (uint32, error) {
	stmt, args, err := r.goqu.Delete("tasks").Returning(goqu.C("id")).Where(
		goqu.C("id").Eq(id),
	).Prepared(true).ToSQL()
	if err != nil {
		return 0, err
	}

	var taskID uint32
	if err := r.DB.QueryRow(ctx, stmt, args...).Scan(&taskID); err != nil {
		return 0, err
	}

	return taskID, nil
}
