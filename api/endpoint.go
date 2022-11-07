package api

import (
	"fwork/response"
	"net/http"
)

//Handler handles HTTP requests with
//a given scope
type Handler func(scope Scope)

//NotFound is the default handler
//used if no handler matched the request
func NotFound(scope Scope) {
	scope.JsonRes(http.StatusNotFound, response.Void{})
}
