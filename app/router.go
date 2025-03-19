package app

import (
	"eisenhower-todo-api/controller"
	"eisenhower-todo-api/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(controller controller.TodoController) *httprouter.Router {

	router := httprouter.New()

	router.GET("/api/todos", controller.FindAll)
	router.GET("/api/todos/:id", controller.FindById)

	router.POST("/api/todos", controller.Create)

	router.PATCH("/api/todos/:id", controller.Patch)

	router.DELETE("/api/todos/:id", controller.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
