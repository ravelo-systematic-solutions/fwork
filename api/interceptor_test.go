package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInterceptor(t *testing.T) {
	//given
	method := http.MethodGet
	url := "/some-url?hello=world"
	w := httptest.NewRecorder()
	r, _ := http.NewRequest(method, url, nil)
	s := NewScope(w, r)
	m := Measurement{}
	s.s = http.StatusAccepted

	//when
	m.Before(s)
	m.After(s)

	//then
	if m.statusCode != http.StatusAccepted {
		t.Errorf(
			"Before()|After(), got %v but want %v",
			m.statusCode,
			http.StatusAccepted,
		)
	}

	if m.method != method {
		t.Errorf(
			"Before()|After(), got %v but want %v",
			m.method,
			method,
		)
	}

	if m.resource != url {
		t.Errorf(
			"Before()|After(), got %v but want %v",
			m.resource,
			url,
		)
	}
}
