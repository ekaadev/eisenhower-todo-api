package controller

import (
	"eisenhower-todo-api/exception"
	"eisenhower-todo-api/helper"
	"eisenhower-todo-api/model/web"
	"eisenhower-todo-api/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Struct TodoControllerImpl
// Implement TodoController interface
// Use for handle the request from client
// @Attribute, TodoService: service.TodoService
type TodoControllerImpl struct {
	TodoService service.TodoService
}

// Function NewTodoConctroller
// Use for create new instance of TodoControllerImpl (Constructor)
// @Parameter, TodoService: service.TodoService
// @Return, TodoController
func NewTodoController(todoService service.TodoService) TodoController {
	return &TodoControllerImpl{
		TodoService: todoService,
	}
}

// Function Create
// Implement Create method from TodoController interface
// @Parameter, writer: http.ResponseWriter, request: *http.Request, params: httprouter.Params
// Description, Use for handle request to create new todo
func (controller *TodoControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	todoCreateRequest := web.TodoCreateRequest{}

	helper.ReadFromRequestBody(request, &todoCreateRequest)

	todoResponse := controller.TodoService.Create(request.Context(), todoCreateRequest)

	webResponse := web.WebResponse{
		Code:   http.StatusCreated,
		Status: "CREATED",
		Data:   todoResponse,
	}

	helper.WriteToResponseBody(writer, webResponse, http.StatusCreated)
}

// Function Patch
// Implement Patch method from TodoController interface
// @Parameter, writer: http.ResponseWriter, request: *http.Request, params: httprouter.Params
// Description, Use for handle request to update todo but not all fields
func (controller *TodoControllerImpl) Patch(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	todoPatchRequest := web.TodoPatchRequest{}

	// Get todo id from request params
	// convert string to int
	// if err its mean cannot convert, send panic
	todoId := params.ByName("id")
	id, err := strconv.Atoi(todoId)
	if err != nil {
		panic(exception.ErrParams)
	}

	todoPatchRequest.Id = id

	helper.ReadFromRequestBody(request, &todoPatchRequest)

	todoResponse := controller.TodoService.Patch(request.Context(), todoPatchRequest)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   todoResponse,
	}

	helper.WriteToResponseBody(writer, webResponse, http.StatusOK)
}

// Function Delete
// Implement Delete method from TodoController interface
// @Parameter, writer: http.ResponseWriter, request: *http.Request, params: httprouter.Params
// Description, Use for handle request to delete exist todo
func (controller *TodoControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	todoId := params.ByName("id")

	// Get todo id from request params
	// convert string to int
	// if err its mean cannot convert, send panic
	id, err := strconv.Atoi(todoId)
	if err != nil {
		panic(exception.ErrParams)
	}

	controller.TodoService.Delete(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse, http.StatusOK)
}

// Function FindById
// Implement FindById method from TodoController interface
// @Parameter, writer: http.ResponseWriter, request: *http.Request, params: httprouter.Params
// Description, Use for handle request to find todo by id
func (controller *TodoControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	todoId := params.ByName("id")

	// Get todo id from request params
	// convert string to int
	// if err its mean cannot convert, send panic
	id, err := strconv.Atoi(todoId)
	if err != nil {
		panic(exception.ErrParams)
	}

	todoResponse := controller.TodoService.FindById(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   todoResponse,
	}

	helper.WriteToResponseBody(writer, webResponse, http.StatusOK)
}

// Function FindAll
// Implement FindAll method from TodoController interface
// @Parameter, writer: http.ResponseWriter, request: *http.Request, params: httprouter.Params
// Description, Use for handle request to find all todos
func (controller *TodoControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	todoResponses := controller.TodoService.FindAll(request.Context())

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   todoResponses,
	}

	helper.WriteToResponseBody(writer, webResponse, http.StatusOK)
}
