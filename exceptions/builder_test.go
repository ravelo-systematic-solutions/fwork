package exceptions_test

import (
	"fwork/exceptions"
	"fwork/utils"
	"reflect"
	"testing"
)

func TestNewBuilder(t *testing.T) {
	//given
	expected := "*builder"

	//when
	b := exceptions.NewBuilder()

	//then
	actual := utils.GetType(b)
	if actual != expected {
		t.Errorf("got %s but want %s", actual, expected)
	}
}

func TestBuilder_Exception(t *testing.T) {
	//given
	expected := &exceptions.Exception{
		Code:    "C1",
		Message: "M1",
		Data: []exceptions.Data{
			{Name: "n1", Tag: "t1", Value: "v1"},
			{Name: "n2", Tag: "t2", Value: 2},
			{Name: "n3", Tag: "t3", Value: true},
		},
	}
	b := exceptions.NewBuilder()
	b.SetCode(expected.Code)
	b.SetMessage(expected.Message)
	b.Include(expected.Data[0])
	b.Include(expected.Data[1])
	b.Include(expected.Data[2])

	//when
	actual := b.Exception()

	//then
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %s but want %s", actual, expected)
	}
}
