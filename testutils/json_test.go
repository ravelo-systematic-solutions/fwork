package testutils

import (
	"fwork/exceptions"
	"reflect"
	"strings"
	"testing"
)

func TestJsonToVar_success(t *testing.T) {
	//given
	type person struct {
		Name string `json:"name"`
	}
	var actual person
	expected := person{
		Name: "Jhonny",
	}

	//when
	JsonToVar(strings.NewReader(`{"name":"Jhonny"}`), &actual)

	//then
	if actual != expected {
		t.Errorf(
			"JsonToVar(), got %v but want %v",
			actual,
			expected,
		)
	}
}

func TestJsonToVar_invalidJson(t *testing.T) {
	//given
	e := exceptions.NewBuilder()
	e.SetCode(exceptions.InvalidJsonCode)
	e.SetMessage(exceptions.InvalidJsonMessage)

	expected := e.Exception()

	//when
	actual := JsonToVar(strings.NewReader(`{"name}`), nil)

	//then
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(
			"JsonToVar(), got %v but want %v",
			actual,
			expected,
		)
	}
}
