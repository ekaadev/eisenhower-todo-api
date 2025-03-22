package app

import (
	"eisenhower-todo-api/controller"
	"eisenhower-todo-api/exception"
	"eisenhower-todo-api/middleware"

	"github.com/julienschmidt/httprouter"
)

// Function NewRouter
// Use for initialize the router in the application and register the router
// Return *httprouter.Router
func NewRouter(controller controller.TodoController) *httprouter.Router {

	// Initialize the router from library httprouter, https://github.com/julienschmidt/httprouter
	router := httprouter.New()

	// Register the routes
	// GET /api/todos, get all todos
	router.GET("/api/todos", middleware.AuthMiddleware(controller.FindAll))
	// GET /api/todos/:id, get todo by id
	router.GET("/api/todos/:id", middleware.AuthMiddleware(controller.FindById))

	// POST /api/todos, create new todo
	router.POST("/api/todos", middleware.AuthMiddleware(controller.Create))

	// PATCH /api/todos/:id, update todo by id
	router.PATCH("/api/todos/:id", middleware.AuthMiddleware(controller.Patch))

	// DELETE /api/todos/:id, delete todo by id
	router.DELETE("/api/todos/:id", middleware.AuthMiddleware(controller.Delete))

	// Assign the panic handler
	router.PanicHandler = exception.ErrorHandler

	return router
}
