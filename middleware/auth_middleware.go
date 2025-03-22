package middleware

import (
	"eisenhower-todo-api/helper"
	"eisenhower-todo-api/model/web"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

// Function AuthMiddleware
// Use for handle the request to check the api key as Middleware
// @Parameter, next httprouter.Handle
// @Return, httprouter.Handle
func AuthMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		api_key := request.Header.Get("X-API-Key")

		err := godotenv.Load(".env")
		helper.PanicIfError(err)

		// Check the api key, if not match then return unauthorized
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
