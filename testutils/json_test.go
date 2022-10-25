package testutils

import (
	"strings"
	"testing"
)

func TestJsonToVar(t *testing.T) {
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
