package adapter

import (
	"net/http"

	"github.com/GuidoGdR/go-speed-test/internal/platform/errorBody"
)

type Response struct {
	StatusCode int               `json:"statusCode"`
	Headers    map[string]string `json:"headers"`
	Body       any               `json:"body"`
}

//func (r *ErrorResponse) Generic(status string, message string, fields map[string][]string) *ErrorResponse {

func (r *Response) OK(body any) *Response {
	r.StatusCode = http.StatusOK
	r.Body = body
	return r
}

func (r *Response) Created(body any) *Response {
	r.StatusCode = http.StatusCreated
	r.Body = body
	return r
}

func (r *Response) Accepted(body any) *Response {
	r.StatusCode = http.StatusAccepted
	r.Body = body
	return r
}

func (r *Response) InternalServerError() *Response {
	r.StatusCode = http.StatusInternalServerError
	r.Body = errorBody.InternalServerError()
	return r
}

func (r *Response) BadRequest(message string) *Response {
	r.StatusCode = http.StatusBadRequest
	r.Body = errorBody.BadRequest(message)
	return r
}
func (r *Response) BadRequestFormat() *Response {
	r.StatusCode = http.StatusBadRequest
	r.Body = errorBody.BadRequest("Format error")
	return r
}
func (r *Response) BadRequestFields(message string, fields map[string][]string) *Response {
	r.StatusCode = http.StatusBadRequest
	r.Body = errorBody.BadRequestFields(message, fields)
	return r
}

func (r *Response) MethodNotAllowed(message string) *Response {
	r.StatusCode = http.StatusMethodNotAllowed
	r.Body = errorBody.MethodNotAllowed(message)
	return r
}

func (r *Response) Unauthorized(message string) *Response {
	r.StatusCode = http.StatusUnauthorized
	r.Body = errorBody.MethodNotAllowed(message)
	return r
}

func (r *Response) UnauthorizedInvalidCredentials(message string) *Response {
	r.StatusCode = http.StatusUnauthorized
	r.Body = errorBody.MethodNotAllowed("Invalid credentials")
	return r
}
