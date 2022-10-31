package api

import (
	"time"
)

type InterceptorI interface {
	Before(s *Scope) error
	After(s *Scope) error
}

//Interceptor can be executed before
//and after a request with the given
//Scope
type Interceptor func(s *Scope) error

//Measurement logs information about the
//the api and its performance
type Measurement struct {
	start      time.Time
	end        time.Time
	duration   time.Duration
	method     string
	resource   string
	statusCode int
}

//Before gets called before the endpoint
//gets called
func (m *Measurement) Before(s *Scope) error {
	m.start = time.Now()
	m.resource = s.Path()
	m.method = s.Method()
	return nil
}

//After gets called after the endpoint
//gets called
func (m *Measurement) After(s *Scope) error {
	m.end = time.Now()
	m.duration = m.end.Sub(m.start)
	m.statusCode = s.s
	return nil
}

//AddInterceptor includes interceptor handlers
//to request
func (e *engine) AddInterceptor(i InterceptorI) {
	e.i = append(e.i, i)
}
