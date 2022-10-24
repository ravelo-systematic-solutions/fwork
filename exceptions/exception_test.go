package exceptions

import (
	"testing"
)

func TestException_Error(t *testing.T) {
	//given
	expected := "E{0123}:M{Hello World}:P{[{n1 t1 v1} {n2 t2 v2} {n3 t3 v3}]}"
	e := Exception{
		Code:    "0123",
		Message: "Hello World",
		Data: []Data{
			{Name: "n1", Tag: "t1", Value: "v1"},
			{Name: "n2", Tag: "t2", Value: "v2"},
			{Name: "n3", Tag: "t3", Value: "v3"},
		},
	}

	//when
	actual := e.Error()

	//then
	if expected != actual {
		t.Errorf("got %s but want %s", actual, expected)
	}
}
