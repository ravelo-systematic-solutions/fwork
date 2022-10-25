package api

import (
	"encoding/json"
	"fwork/exceptions"
	"net/http"
)

// Scope holds Api Handler context
type Scope struct {
	w http.ResponseWriter
	r *http.Request
	d map[string]any
}

//GetData gets available additional
//contextual data to the request.
//An exception will be thrown if
//and when the passed key does not exist
func (scope *Scope) GetData(key string) (any, error) {
	if val, ok := scope.d[key]; !ok {
		exception := exceptions.NewBuilder()
		exception.SetCode(exceptions.KeyNotFoundCode)
		exception.SetMessage(exceptions.KeyNotFoundMessage)

		return nil, exception.Exception()
	} else {
		return val, nil
	}
}

//SetData sets additional contextual
//data to the request. An exception
//will be thrown if the key already
//exists
func (scope *Scope) SetData(key string, val any) error {
	if _, ok := scope.d[key]; ok {
		exception := exceptions.NewBuilder()
		exception.SetCode(exceptions.DuplicatedKeyCode)
		exception.SetMessage(exceptions.DuplicatedKeyMessage)

		return exception.Exception()
	}

	scope.OverrideData(key, val)

	return nil
}

//OverrideData sets additional contextual
//data to the request regarless if the key
//already exists or not
func (scope *Scope) OverrideData(key string, val any) {
	scope.d[key] = val
}

//Method retrieves the requested Method
func (scope *Scope) Method() string {
	return scope.r.Method
}

//Path retrieves the requested URL
func (scope *Scope) Path() string {
	return scope.r.URL.Path
}

// JsonRes replies to client with json format
func (s *Scope) JsonRes(status int, body interface{}) {
	//s.w.Header().Set("Access-Control-Allow-Origin", "*")
	s.w.Header().Set("Content-Type", "application/json")
	s.w.WriteHeader(status)

	json.NewEncoder(s.w).Encode(body)
}

// QueryValue extracts a string from Query parameter
// Sets default value if absent (eg. /a?b=c)
func (s *Scope) QueryValue(key string) string {
	return s.r.URL.Query().Get(key)
}
