package repository

import (
	"context"
	"database/sql"
	"eisenhower-todo-api/exception"
	"eisenhower-todo-api/helper"
	"eisenhower-todo-api/model/domain"
)

type TodoRepositoryImpl struct {
}

// Constructor
func NewTodoRepository() TodoRepository {
	return &TodoRepositoryImpl{}
}

// Implement method create from Interface TodoRepository
func (repository *TodoRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, todo domain.Todo) domain.Todo {
	SQL := "INSERT INTO todos (title, description, type, is_done) VALUES($1, $2, $3, $4) RETURNING id"

	var description string

	if todo.Description.Valid {
		description = todo.Description.String
	} else {
		description = ""
	}

	result := tx.QueryRowContext(ctx, SQL, todo.Title, description, todo.Type, todo.IsDone)

	err := result.Scan(&todo.Id)
	helper.PanicIfError(err)

	return todo
}

// Implement method patch from Interface TodoRepository
func (repository *TodoRepositoryImpl) Patch(ctx context.Context, tx *sql.Tx, todo domain.Todo) domain.Todo {
	SQL := "UPDATE todos SET title = $1, description = $2, type = $3, is_done = $4 WHERE id = $5"

	var description string

	if todo.Description.Valid {
		description = todo.Description.String
	} else {
		description = ""
	}

	_, err := tx.ExecContext(ctx, SQL, todo.Title, description, todo.Type, todo.IsDone, todo.Id)
	helper.PanicIfError(err)

	return todo
}

// Implement method delete from Interface TodoRepository
func (repository *TodoRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, todoId int) {
	SQL := "DELETE FROM todos WHERE id = $1"

	_, err := tx.ExecContext(ctx, SQL, todoId)
	helper.PanicIfError(err)
}

// Implement method find by id from Interface TodoRepository
func (repository *TodoRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, todoId int) (domain.Todo, error) {
	SQL := "SELECT id, title, description, type, is_done, created_at, updated_at FROM todos WHERE id = $1"

	rows, err := tx.QueryContext(ctx, SQL, todoId)
	helper.PanicIfError(err)

	defer rows.Close()

	todo := domain.Todo{}

	if rows.Next() {
		rows.Scan(
			&todo.Id,
			&todo.Title,
			&todo.Description,
			&todo.Type,
			&todo.IsDone,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		)

		return todo, nil
	} else {
		return todo, exception.ErrNotFound
	}
}

// Implement method find all from Interface TodoRepository
func (repository *TodoRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Todo {
	SQL := "SELECT id, title, description, type, is_done, created_at, updated_at FROM todos"

	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	defer rows.Close()

	var todos []domain.Todo

	for rows.Next() {
		todo := domain.Todo{}

		rows.Scan(
			&todo.Id,
			&todo.Title,
			&todo.Description,
			&todo.Type,
			&todo.IsDone,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		)

		todos = append(todos, todo)
	}

	return todos
}
