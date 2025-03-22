package service

import (
	"context"
	"eisenhower-todo-api/model/web"
)

// TodoService interface
// Contract for TodoService
// @Method Create, for handle service to create new todo
// @Method Patch, for handle service to update todo but not all fields
// @Method Delete, for handle service to delete todo
// @Method FindById, for handle service to get todo by id
// @Method FindAll, for handle service to get all todos
type TodoService interface {
	Create(ctx context.Context, request web.TodoCreateRequest) web.TodoResponse
	Patch(ctx context.Context, request web.TodoPatchRequest) web.TodoResponse
	Delete(ctx context.Context, todoId int)
	FindById(ctx context.Context, todoId int) web.TodoResponse
	FindAll(ctx context.Context) []web.TodoResponse
}
