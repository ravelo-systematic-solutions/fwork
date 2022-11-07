package api

import (
	"net/http"
	"testing"
)

func TestNotFound(t *testing.T) {
	//given
	scope := &scope{}

	//when
	NotFound(scope)

	//then
	if scope.s != http.StatusNotFound {
		t.Errorf(
			"NotFound(), got %v but want %v",
			scope.s,
			http.StatusNotFound,
		)
	}
}
