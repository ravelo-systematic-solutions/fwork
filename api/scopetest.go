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
	w http.ResponseWriter
	r *http.Request
	s int
	b []byte
	c Controller
	d map[string]any
}

func (s *scopeTest) GetData(key string) (any, error) {
	return nil, nil
}

func (s *scopeTest) SetData(key string, val any) error {
	return nil
}

func (s *scopeTest) OverrideData(key string, val any) {

}

func (s *scopeTest) Method() string {
	return ""
}

func (s *scopeTest) Path() string {
	return ""
}

func (s *scopeTest) Reply(status int, body interface{}) {

}

func (s *scopeTest) QueryValue(key string) string {
	return ""
}

func (s *scopeTest) ValidateQuery(payload interface{}) error {

	return nil
}

func (s *scopeTest) ValidateJsonBody(payload interface{}) error {
	return nil
}

func (s *scopeTest) ValidateHeaders(payload interface{}) error {
	return nil
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
func NewTestScope(method, url string, body io.Reader, c Controller) *scopeTest {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest(method, url, body)
	s := scopeTest{
		r: r,
		w: w,
		c: c,
		d: make(map[string]any),
	}

	return &s
}
