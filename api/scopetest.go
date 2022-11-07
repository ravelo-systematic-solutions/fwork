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
	IsStatus(status int) error
	IsJsonRes(body interface{}) error
}

type scopeTest struct {
	scope
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

func (s *scopeTest) IsJsonRes(body interface{}) error {
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

//NewTestScope creates a Handler's scope instance
//for testing purposes
func NewTestScope(method, url string, body io.Reader, c Controller) *scopeTest {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest(method, url, body)
	s := scopeTest{
		scope{
			r: r,
			w: w,
			d: make(map[string]any),
		},
	}

	c.GetHandler(http.MethodGet, c.Url())(&s)

	return &s
}
