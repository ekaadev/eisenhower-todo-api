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

type TodoControllerImpl struct {
	TodoService service.TodoService
}

func NewTodoController(todoService service.TodoService) TodoController {
	return &TodoControllerImpl{
		TodoService: todoService,
	}
}

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

func (controller *TodoControllerImpl) Patch(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	todoPatchRequest := web.TodoPatchRequest{}

	todoId := params.ByName("id")
	id, err := strconv.Atoi(todoId)
	if err != nil {
		panic(exception.ErrParams.Error)
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

func (controller *TodoControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	todoId := params.ByName("id")

	id, err := strconv.Atoi(todoId)
	if err != nil {
		panic(exception.ErrParams.Error)
	}

	controller.TodoService.Delete(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse, http.StatusOK)
}

func (controller *TodoControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	todoId := params.ByName("id")

	id, err := strconv.Atoi(todoId)
	if err != nil {
		panic(exception.ErrParams.Error)
	}

	todoResponse := controller.TodoService.FindById(request.Context(), id)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   todoResponse,
	}

	helper.WriteToResponseBody(writer, webResponse, http.StatusOK)
}

func (controller *TodoControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	todoResponses := controller.TodoService.FindAll(request.Context())

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   todoResponses,
	}

	helper.WriteToResponseBody(writer, webResponse, http.StatusOK)
}
