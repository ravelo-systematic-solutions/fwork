package api

import (
	"net/http"
	"net/url"
)

type Request interface {
	QueryValue(k, v string)
	EncodedQuery() string
	HeaderValue(k, v string)
}

type request struct {
	query   *url.Values
	headers *http.Header
	body    interface{}
}

func (r *request) QueryValue(k, v string) {
	r.query.Set(k, v)
}

func (r *request) EncodedQuery() string {
	return r.query.Encode()
}

func (r *request) HeaderValue(k, v string) {
	r.headers.Set(k, v)
}

func NewTestRequest() *request {
	return &request{
		query:   &url.Values{},
		headers: &http.Header{},
	}
}
