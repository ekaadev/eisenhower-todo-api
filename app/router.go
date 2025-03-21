package app

import (
	"eisenhower-todo-api/controller"
	"eisenhower-todo-api/exception"
	"eisenhower-todo-api/middleware"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(controller controller.TodoController) *httprouter.Router {

	router := httprouter.New()

	router.GET("/api/todos", middleware.AuthMiddleware(controller.FindAll))
	router.GET("/api/todos/:id", middleware.AuthMiddleware(controller.FindById))

	router.POST("/api/todos", middleware.AuthMiddleware(controller.Create))

	router.PATCH("/api/todos/:id", middleware.AuthMiddleware(controller.Patch))

	router.DELETE("/api/todos/:id", middleware.AuthMiddleware(controller.Delete))

	router.PanicHandler = exception.ErrorHandler

	return router
}
