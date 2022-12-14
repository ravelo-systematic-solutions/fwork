package testutils_test

import (
	"github.com/ravelo-systematic-solutions/fwork/testutils"
	"testing"
)

type SampleStruct struct{}

func SampleFunc() {}

func TestGetStructType_Pointer(t *testing.T) {
	//given
	s := &SampleStruct{}
	want := "*SampleStruct"

	//when & then
	if got := testutils.GetType(s); got != want {
		t.Errorf("GetType() = %v, want %v", got, want)
	}
}

func TestGetStructType_var(t *testing.T) {
	//given
	s := SampleStruct{}
	want := "SampleStruct"

	//when & then
	if got := testutils.GetType(s); got != want {
		t.Errorf("GetType() = %v, want %v", got, want)
	}
}

func TestFuncName_func(t *testing.T) {
	//given
	want := "github.com/ravelo-systematic-solutions/fwork/testutils_test.SampleFunc"

	//when & then
	if got := testutils.GetType(SampleFunc); got != want {
		t.Errorf("FuncName() = %v, want %v", got, want)
	}
}
