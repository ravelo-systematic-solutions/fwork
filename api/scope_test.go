package api

import (
	"encoding/json"
	"fwork/exceptions"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestScope_GetData_success(t *testing.T) {
	tests := []struct {
		name string
		val  any
	}{
		{"store string value", "hello world"},
		{"store bool value", true},
		{"store byte value", byte(2)},
		{"store int value", int(3)},
		{"store int8 value", int8(123)},
		{"store int16 value", int8(123)},
		{"store int32 value", int8(123)},
		{"store int64 value", int8(123)},
		{"store float32 value", float32(20.22)},
		{"store float64 value", float64(20.22)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//given
			scope := Scope{
				d: make(map[string]any, 0),
			}
			scope.SetData("key", test.val)
			//when
			actual, err := scope.GetData("key")

			//then
			if err != nil {
				t.Errorf(
					"GetData(), got unexpected error %v",
					err,
				)
			}

			if actual != test.val {
				t.Errorf(
					"SetData(), got %v but want %v",
					actual,
					test.val,
				)
			}
		})
	}
}

func TestScope_GetData_KeyNotFound(t *testing.T) {
	//given
	scope := Scope{
		d: make(map[string]any, 0),
	}

	exception := exceptions.NewBuilder()
	exception.SetCode(exceptions.ResourceNotFoundCode)
	exception.SetMessage(exceptions.ResourceNotFoundMessage)
	expected := exception.Exception()

	//when
	_, actualErr := scope.GetData("key")
	actual := actualErr.(*exceptions.Exception)

	//then
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(
			"SetData(), got %v but want %v",
			actual,
			expected,
		)
	}
}

func TestScope_SetData_success(t *testing.T) {
	tests := []struct {
		name string
		val  any
	}{
		{"store string value", "hello world"},
		{"store bool value", true},
		{"store byte value", byte(2)},
		{"store int value", int(3)},
		{"store int8 value", int8(123)},
		{"store int16 value", int8(123)},
		{"store int32 value", int8(123)},
		{"store int64 value", int8(123)},
		{"store float32 value", float32(20.22)},
		{"store float64 value", float64(20.22)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//given
			scope := Scope{
				d: make(map[string]any, 0),
			}

			actual := scope.SetData("key", test.val)

			if actual != nil {
				t.Errorf(
					"SetData(), got unexpected error %v",
					actual,
				)
			}
		})
	}
}

func TestScope_SetData_DuplicatedKey(t *testing.T) {
	tests := []struct {
		name string
		val  any
	}{
		{"store string value", "hello world"},
		{"store bool value", true},
		{"store byte value", byte(2)},
		{"store int value", int(3)},
		{"store int8 value", int8(123)},
		{"store int16 value", int8(123)},
		{"store int32 value", int8(123)},
		{"store int64 value", int8(123)},
		{"store float32 value", float32(20.22)},
		{"store float64 value", float64(20.22)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//given
			scope := Scope{
				d: make(map[string]any, 0),
			}
			scope.SetData("key", test.val)

			exception := exceptions.NewBuilder()
			exception.SetCode(exceptions.ResourceDuplicatedCode)
			exception.SetMessage(exceptions.ResourceDuplicatedMessage)
			expected := exception.Exception()

			//when
			actualErr := scope.SetData("key", test.val)
			actual := actualErr.(*exceptions.Exception)

			//then
			if !reflect.DeepEqual(actual, expected) {
				t.Errorf(
					"SetData(), got %v but want %v",
					actual,
					expected,
				)
			}
		})
	}
}

func TestScope_OverrideData_success(t *testing.T) {
	tests := []struct {
		name string
		val  any
		err  error
	}{
		{"store string value", "hello world", nil},
		{"store bool value", true, nil},
		{"store byte value", byte(2), nil},
		{"store int value", int(3), nil},
		{"store int8 value", int8(123), nil},
		{"store int16 value", int8(123), nil},
		{"store int32 value", int8(123), nil},
		{"store int64 value", int8(123), nil},
		{"store float32 value", float32(20.22), nil},
		{"store float64 value", float64(20.22), nil},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//given
			scope := Scope{
				d: make(map[string]any, 0),
			}
			unexpectedErr := test.err

			actual := scope.SetData("key", test.val)

			if actual != nil {
				t.Errorf(
					"SetData(), unexpected error %v",
					unexpectedErr,
				)
			}
		})
	}
}

func TestScope_Method(t *testing.T) {
	//given
	expected := http.MethodGet
	req, _ := http.NewRequest(expected, "/some-url", nil)
	scope := Scope{
		r: req,
	}

	//when
	actual := scope.Method()

	//then
	if actual != expected {
		t.Errorf(
			"Method(), got %v but want %v",
			actual,
			expected,
		)
	}
}

func TestScope_Path(t *testing.T) {
	//given
	expected := "/some-url"
	req, _ := http.NewRequest(http.MethodGet, expected, nil)
	scope := Scope{
		r: req,
	}

	//when
	actual := scope.Path()

	//then
	if actual != expected {
		t.Errorf(
			"Path(), got %v but want %v",
			actual,
			expected,
		)
	}
}

func TestScope_JsonRes_success(t *testing.T) {
	//given
	type person struct {
		Name string `json:"name"`
	}
	expected, _ := json.Marshal(person{
		Name: "Jhonny",
	})
	scope := Scope{}

	//when
	scope.JsonRes(http.StatusAccepted, person{
		Name: "Jhonny",
	})

	//then
	if string(scope.b) != string(expected) {
		t.Errorf(
			"JsonRes(), got %v but want %v",
			string(scope.b),
			string(expected),
		)
	}
}

func TestScope_QueryValue(t *testing.T) {
	//given
	expected := "world"
	req := httptest.NewRequest(http.MethodGet, "/some-url?hello="+expected, nil)
	scope := Scope{
		r: req,
	}

	//when
	actual := scope.QueryValue("hello")

	//then
	if actual != expected {
		t.Errorf(
			"QueryValue(), got %v but want %v",
			actual,
			expected,
		)
	}
}

func TestNewRequest(t *testing.T) {
	//given
	w := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodGet, "/some-url", nil)

	//when
	s := NewScope(w, r)

	//then
	if !reflect.DeepEqual(s.w, w) {
		t.Errorf(
			"NewScope(), go %v but want %v",
			s.w,
			w,
		)
	}

	if !reflect.DeepEqual(s.r, r) {
		t.Errorf(
			"NewScope(), go %v but want %v",
			s.r,
			r,
		)
	}
}
