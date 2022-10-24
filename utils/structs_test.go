package utils_test

import (
	"fwork/utils"
	"testing"
)

type SampleStruct struct{}

func SampleFunc() {}

func TestGetStructType_Pointer(t *testing.T) {
	//given
	s := &SampleStruct{}
	want := "*SampleStruct"

	//when & then
	if got := utils.GetType(s); got != want {
		t.Errorf("GetType() = %v, want %v", got, want)
	}
}

func TestGetStructType_var(t *testing.T) {
	//given
	s := SampleStruct{}
	want := "SampleStruct"

	//when & then
	if got := utils.GetType(s); got != want {
		t.Errorf("GetType() = %v, want %v", got, want)
	}
}

func TestFuncName_func(t *testing.T) {
	//given
	want := "fwork/utils_test.SampleFunc"

	//when & then
	if got := utils.GetType(SampleFunc); got != want {
		t.Errorf("FuncName() = %v, want %v", got, want)
	}
}
