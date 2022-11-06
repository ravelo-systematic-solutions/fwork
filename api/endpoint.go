package api

import (
	"fwork/response"
	"net/http"
)

//Handler handles HTTP requests with
//a given Scope
type Handler func(scope ScopeInterface)

//NotFound is the default handler
//used if no handler matched the request
func NotFound(scope ScopeInterface) {
	scope.JsonRes(http.StatusNotFound, response.Void{})
}
