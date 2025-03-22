package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// TodoController interface
// Contract for TodoController
// @Method Create, for handle request to create new todo
// @Method Patch, for handle request to update todo but not all fields
// @Method Delete, for handle request to delete todo
// @Method FindById, for handle request to get todo by id
// @Method FindAll, for handle request to get all todos
type TodoController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Patch(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
