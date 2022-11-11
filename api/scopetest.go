package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

type ScopeTest interface {
	Scope
	IsStatus(status int) error
	ReplyWas(body interface{}) error
}

type scopeTest struct {
	scope
	c Controller
}

func (s *scopeTest) IsStatus(status int) error {
	if s.s != status {
		return errors.New(fmt.Sprintf(
			"got %v but want %v",
			s.s,
			status,
		))
	}
	return nil
}

func (s *scopeTest) ReplyWas(body interface{}) error {
	if bodyByte, err := json.Marshal(body); err != nil {
		return errors.New("failed to encode body")
	} else if bytes.Compare(bodyByte, s.b) != 0 {
		return errors.New(fmt.Sprintf(
			"got %v but want %v",
			string(s.b),
			string(bodyByte),
		))
	}
	return nil
}

func (s *scopeTest) Execute() {
	s.c.GetHandler(http.MethodGet, s.c.Url())(s)
}

//NewTestScope creates a Handler's scope instance
//for testing purposes
func NewTestScope(method string, body io.Reader, query interface{}, c Controller) *scopeTest {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest(method, c.Url(), body)
	s := scopeTest{
		scope: scope{
			r: r,
			w: w,
			d: make(map[string]any),
		},
		c: c,
	}

	return &s
}
