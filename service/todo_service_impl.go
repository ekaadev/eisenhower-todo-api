package service

import (
	"context"
	"database/sql"
	"eisenhower-todo-api/helper"
	"eisenhower-todo-api/model/domain"
	"eisenhower-todo-api/model/web"
	"eisenhower-todo-api/repository"
	"time"

	"github.com/go-playground/validator/v10"
)

type TodoServiceImpl struct {
	Validate       *validator.Validate
	DB             *sql.DB
	TodoRepository repository.TodoRepository
}

func NewTodoService(validate *validator.Validate, db *sql.DB, todoRepository repository.TodoRepository) TodoService {
	return &TodoServiceImpl{
		Validate:       validate,
		DB:             db,
		TodoRepository: todoRepository,
	}
}

func (service *TodoServiceImpl) Create(ctx context.Context, request web.TodoCreateRequest) web.TodoResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	description := sql.NullString{
		String: request.Description,
		Valid:  true,
	}

	todo := domain.Todo{
		Title:       request.Title,
		Description: description,
		Type:        domain.TodoType(request.Type),
	}

	todo = service.TodoRepository.Create(ctx, tx, todo)

	return web.TodoResponse{
		Id:          todo.Id,
		Title:       todo.Title,
		Description: todo.Description.String,
		Type:        string(todo.Type),
		IsDone:      todo.IsDone,
		CreatedAt:   todo.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   todo.CreatedAt.Format(time.RFC3339),
	}
}

func (service *TodoServiceImpl) Patch(ctx context.Context, request web.TodoPatchRequest) web.TodoResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	todo, err := service.TodoRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	if request.Title != "" {
		todo.Title = request.Title
	}

	if request.Description != "" {
		description := sql.NullString{
			String: request.Description,
			Valid:  true,
		}

		todo.Description = description
	}

	if request.Type != "" {
		todo.Type = domain.TodoType(request.Type)
	}

	if request.IsDone {
		todo.IsDone = request.IsDone
	}

	todo = service.TodoRepository.Patch(ctx, tx, todo)

	return web.TodoResponse{
		Id:          todo.Id,
		Title:       todo.Title,
		Description: todo.Description.String,
		Type:        string(todo.Type),
		IsDone:      todo.IsDone,
		CreatedAt:   todo.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   todo.UpdatedAt.Format(time.RFC3339),
	}

}

func (service *TodoServiceImpl) Delete(ctx context.Context, todoId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	todo, err := service.TodoRepository.FindById(ctx, tx, todoId)
	helper.PanicIfError(err)

	service.TodoRepository.Delete(ctx, tx, todo.Id)
}

func (service *TodoServiceImpl) FindById(ctx context.Context, todoId int) web.TodoResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	todo, err := service.TodoRepository.FindById(ctx, tx, todoId)
	helper.PanicIfError(err)

	return web.TodoResponse{
		Id:          todo.Id,
		Title:       todo.Title,
		Description: todo.Description.String,
		Type:        string(todo.Type),
		IsDone:      todo.IsDone,
		CreatedAt:   todo.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   todo.UpdatedAt.Format(time.RFC3339),
	}
}

func (service *TodoServiceImpl) FindAll(ctx context.Context) []web.TodoResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	todos := service.TodoRepository.FindAll(ctx, tx)

	var todoResponses []web.TodoResponse

	for _, todo := range todos {
		todoResponse := web.TodoResponse{
			Id:          todo.Id,
			Title:       todo.Title,
			Description: todo.Description.String,
			Type:        string(todo.Type),
			IsDone:      todo.IsDone,
			CreatedAt:   todo.CreatedAt.Format(time.RFC3339),
			UpdatedAt:   todo.UpdatedAt.Format(time.RFC3339),
		}

		todoResponses = append(todoResponses, todoResponse)
	}

	return todoResponses
}
