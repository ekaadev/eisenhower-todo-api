package repository

import (
	"context"
	"database/sql"
	"eisenhower-todo-api/model/domain"
)

// TodoRepository interface
// Contract for TodoRepository
// @Method Create, for create new todo
// @Method Patch, for update todo but not all fields
// @Method Delete, for delete todo
// @Method FindById, for get todo by id
// @Method FindAll, for get all todos
type TodoRepository interface {
	Create(ctx context.Context, tx *sql.Tx, todo domain.Todo) domain.Todo
	Patch(ctx context.Context, tx *sql.Tx, todo domain.Todo) domain.Todo
	Delete(ctx context.Context, tx *sql.Tx, todoId int)
	FindById(ctx context.Context, tx *sql.Tx, todoId int) (domain.Todo, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Todo
}
