package api

import (
	"encoding/json"
	"github.com/ravelo-systematic-solutions/fwork/exceptions"
	"net/http"
)

type Scope interface {
	GetData(key string) (any, error)
	SetData(key string, val any) error
	OverrideData(key string, val any)
	Method() string
	Path() string
	Reply(status int, body interface{})
	QueryValue(key string) string
	ValidateQuery(payload interface{}) error
	ValidateJsonBody(payload interface{}) error
	ValidateHeaders(payload interface{}) error
}

// scope holds Api Handler context
type scope struct {
	w http.ResponseWriter
	r *http.Request
	s int
	b []byte
	d map[string]any
}

//GetData gets available additional
//contextual data to the request.
//An exception will be thrown if
//and when the passed key does not exist
func (s *scope) GetData(key string) (any, error) {
	if val, ok := s.d[key]; !ok {
		exception := exceptions.NewBuilder()
		exception.SetCode(exceptions.ResourceNotFoundCode)
		exception.SetMessage(exceptions.ResourceNotFoundMessage)

		return nil, exception.Build()
	} else {
		return val, nil
	}
}

//SetData sets additional contextual
//data to the request. An exception
//will be thrown if the key already
//exists
func (s *scope) SetData(key string, val any) error {
	if _, ok := s.d[key]; ok {
		exception := exceptions.NewBuilder()
		exception.SetCode(exceptions.ResourceDuplicatedCode)
		exception.SetMessage(exceptions.ResourceDuplicatedMessage)

		return exception.Build()
	}

	s.OverrideData(key, val)

	return nil
}

//OverrideData sets additional contextual
//data to the request regarless if the key
//already exists or not
func (s *scope) OverrideData(key string, val any) {
	s.d[key] = val
}

//Method retrieves the requested Method
func (s *scope) Method() string {
	return s.r.Method
}

//Path retrieves the requested URL
func (s *scope) Path() string {
	return s.r.URL.RequestURI()
}

// Reply replies to client with json format
func (s *scope) Reply(status int, body interface{}) {
	bodyByte, err := json.Marshal(body)
	if err != nil {
		e := exceptions.NewBuilder()
		e.SetCode(exceptions.ResourceNotEncodedCode)
		e.SetMessage(exceptions.ResourceNotEncodedMessage)
		e.Include(exceptions.Data{Value: err})

		bodyByte, _ = json.Marshal(e.Build())
		status = http.StatusInternalServerError
	}

	s.s = status
	s.b = bodyByte
}

// QueryValue extracts a string from Query parameter
// Sets default value if absent (eg. /a?b=c)
func (s *scope) QueryValue(key string) string {
	return s.r.URL.Query().Get(key)
}

//NewScope creates a Handler's scope instance
func NewScope(w http.ResponseWriter, r *http.Request) *scope {
	return &scope{
		r: r,
		w: w,
		d: make(map[string]any),
	}
}
