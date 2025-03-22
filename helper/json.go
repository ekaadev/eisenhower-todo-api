package helper

import (
	"encoding/json"
	"net/http"
)

// Function ReadFromRequestBody
// Use for read request body and decode it to target
// @Parameter, request: *http.Request, target any
func ReadFromRequestBody(request *http.Request, target any) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(target)
	PanicIfError(err)
}

// Function WriteToResponseBody
// Use for write response body and encode it to writer
// @Parameter, writer http.ResponseWriter, result any, statusCode int
func WriteToResponseBody(writer http.ResponseWriter, result any, statusCode int) {

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)

	encoder := json.NewEncoder(writer)
	err := encoder.Encode(result)
	PanicIfError(err)
}
