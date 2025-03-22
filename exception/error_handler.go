package exception

import (
	"eisenhower-todo-api/helper"
	"eisenhower-todo-api/model/web"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// List of error message
var (
	ErrNotFound = errors.New("data not found")
	ErrParams   = errors.New("invalid params")
)

// Function ErrorHandler
// Use for handle PanicHandler from httprouter
// @Parameter, writer http.ResponseWriter, request *http.Request, err interface{}
func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {

	e := err.(error)

	if errors.Is(e, ErrNotFound) {
		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   e.Error(),
		}

		helper.WriteToResponseBody(writer, webResponse, http.StatusNotFound)

		return
	}

	if errors.Is(e, validator.ValidationErrors{}) {

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   e.Error(),
		}

		helper.WriteToResponseBody(writer, webResponse, http.StatusBadRequest)

		return
	}

	if errors.Is(e, ErrParams) {

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   e.Error(),
		}

		helper.WriteToResponseBody(writer, webResponse, http.StatusBadRequest)

		return
	}

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   e.Error(),
	}

	helper.WriteToResponseBody(writer, webResponse, http.StatusInternalServerError)
}
