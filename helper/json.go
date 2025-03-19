package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(request *http.Request, target any) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(target)
	PanicIfError(err)
}

func WriteToResponseBody(writer http.ResponseWriter, result any, statusCode int) {

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)

	encoder := json.NewEncoder(writer)
	err := encoder.Encode(result)
	PanicIfError(err)
}
