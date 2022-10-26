package api

import (
	"fwork/response"
	"net/http"
)

//Handler handles HTTP requests with
//a given Scope
type Handler func(scope *Scope)

//Interceptor can be executed before
//and after a request with the given
//Scope
type Interceptor func(scope *Scope) error

//NotFound is the default handler
//used if no handler matched the request
func NotFound(scope *Scope) {
	scope.JsonRes(http.StatusNotFound, response.Void{})
}
