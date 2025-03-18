package service

import (
	"context"
	"eisenhower-todo-api/model/web"
)

type TodoService interface {
	Create(ctx context.Context, request web.TodoCreateRequest) web.TodoResponse
	Patch(ctx context.Context, request web.TodoPatchRequest) web.TodoResponse
	Delete(ctx context.Context, todoId int)
	FindById(ctx context.Context, todoId int) web.TodoResponse
	FindAll(ctx context.Context) []web.TodoResponse
}
