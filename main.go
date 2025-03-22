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

// Main function to run the application
func main() {
	// Initialize the database
	db := app.NewDB()
	defer db.Close()

	// Initialize the validator
	validate := validator.New()

	// Initialize repository, service, and controller
	todoRepository := repository.NewTodoRepository()
	todoService := service.NewTodoService(validate, db, todoRepository)
	todoController := controller.NewTodoController(todoService)

	// Initialize the router
	router := app.NewRouter(todoController)

	// Configure the server, set address and handler
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	// Start the server
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
