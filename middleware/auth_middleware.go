package middleware

import (
	"eisenhower-todo-api/helper"
	"eisenhower-todo-api/model/web"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func AuthMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		api_key := request.Header.Get("X-API-Key")

		err := godotenv.Load(".env")
		helper.PanicIfError(err)

		if api_key != os.Getenv("X_API_KEY") {
			webResponse := web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "UNAUTHORIZED",
			}

			helper.WriteToResponseBody(writer, webResponse, http.StatusUnauthorized)
			return
		}

		next(writer, request, params)
	}

}
