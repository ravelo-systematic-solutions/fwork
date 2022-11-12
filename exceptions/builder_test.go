package exceptions_test

import (
	"github.com/ravelo-systematic-solutions/fwork/exceptions"
	"github.com/ravelo-systematic-solutions/fwork/testutils"
	"reflect"
	"testing"
)

func TestNewBuilder(t *testing.T) {
	//given
	expected := "*builder"

	//when
	b := exceptions.NewBuilder()

	//then
	actual := testutils.GetType(b)
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
	actual := b.Build()

	//then
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %s but want %s", actual, expected)
	}
}

func TestBuilder_IsEmpty_true(t *testing.T) {
	//given
	b := exceptions.NewBuilder()
	b.SetCode("C1")
	b.SetMessage("M1")
	b.Include(exceptions.Data{Name: "n1", Tag: "t1", Value: "v1"})

	//when
	actual := b.IsEmpty()

	//then
	if actual {
		t.Errorf("IsEmpty(), should be false")
	}
}

func TestBuilder_IsEmpty_false(t *testing.T) {
	//given
	b := exceptions.NewBuilder()
	b.SetCode("C1")
	b.SetMessage("M1")

	//when
	actual := b.IsEmpty()

	//then
	if !actual {
		t.Errorf("IsEmpty(), should be true")
	}
}
