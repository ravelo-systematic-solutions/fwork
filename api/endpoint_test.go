package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNotFound(t *testing.T) {
	//given
	res := httptest.NewRecorder()
	scope := &Scope{
		w: res,
	}

	//when
	NotFound(scope)

	//then
	if res.Code != http.StatusNotFound {
		t.Errorf(
			"NotFound(), got %v but want %v",
			res.Code,
			http.StatusNotFound,
		)
	}
}
