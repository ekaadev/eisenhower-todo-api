package main

import (
	"eisenhower-todo-api/app"
	"eisenhower-todo-api/controller"
	"eisenhower-todo-api/helper"
	"eisenhower-todo-api/repository"
	"eisenhower-todo-api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func main() {
	// TODO: Implement main function
	db := app.NewDB()
	defer db.Close()

	validate := validator.New()

	todoRepository := repository.NewTodoRepository()
	todoService := service.NewTodoService(validate, db, todoRepository)
	todoController := controller.NewTodoController(todoService)

	router := app.NewRouter(todoController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
