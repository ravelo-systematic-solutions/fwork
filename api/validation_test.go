package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Sample struct {
	String string  `json:"s" header:"s" query:"s" validate:"required"`
	Int    int     `json:"i" header:"i" query:"i" validate:"required"`
	Bool   bool    `json:"b" header:"b" query:"b" validate:"required"`
	Float  float32 `json:"f" header:"f" query:"f" validate:"required"`
}

func TestScope_JsonBody(t *testing.T) {
	//given
	body := []byte("{\"s\":\"str\",\"i\":123,\"b\":true,\"f\":123.123}")
	req := httptest.NewRequest(http.MethodPost, "/some-url", bytes.NewReader(body))
	scope := Scope{
		r: req,
	}

	var actual Sample

	//when
	err := scope.JsonBody(&actual)

	//then
	if err != nil {
		t.Errorf("JsonBody() unexpected error %v", err)
	}

	if actual.String != "str" {
		t.Errorf("JsonBody() got %s but want %s", actual.String, "str")
	}

	if actual.Int != 123 {
		t.Errorf("JsonBody() got %v but want %v", actual.Int, 123)
	}

	if actual.Bool != true {
		t.Errorf("JsonBody() got %v but want %v", actual.Bool, true)
	}

	if actual.Float != 123.123 {
		t.Errorf("JsonBody() got %v but want %v", actual.Float, 123.123)
	}
}

func TestScope_Headers(t *testing.T) {
	//given
	var actual Sample
	req, _ := http.NewRequest(http.MethodGet, "/some-url", nil)
	req.Header.Set("s", "str")
	req.Header.Set("i", "123")
	req.Header.Set("b", "true")
	req.Header.Set("f", "123.123")
	scope := Scope{
		r: req,
	}

	//when
	scope.Headers(&actual)

	//then
	if actual.String != "str" {
		t.Errorf("Headers() got %s but want %s", actual.String, "str")
	}

	if actual.Int != 123 {
		t.Errorf("Headers() got %v but want %v", actual.Int, 123)
	}

	if actual.Bool != true {
		t.Errorf("Headers() got %v but want %v", actual.Bool, true)
	}

	if actual.Float != 123.123 {
		t.Errorf("Headers() got %v but want %v", actual.Float, 123.123)
	}

}

func TestScope_Query(t *testing.T) {
	//given
	var actual Sample
	url := "/some-url?s=str&i=123&b=true&f=123.123"
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	scope := Scope{
		r: req,
	}

	//when
	scope.Query(&actual)

	//then
	if actual.String != "str" {
		t.Errorf("Query() got %s but want %s", actual.String, "str")
	}

	if actual.Int != 123 {
		t.Errorf("Query() got %v but want %v", actual.Int, 123)
	}

	if actual.Bool != true {
		t.Errorf("Query() got %v but want %v", actual.Bool, true)
	}

	if actual.Float != 123.123 {
		t.Errorf("Query() got %v but want %v", actual.Float, 123.123)
	}

}

func Test_required(t *testing.T) {
	type args struct {
		v any
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"check nil success", args{nil}, true},
		{"string check succeeds", args{"str"}, false},
		{"empty string check fails", args{""}, true},
		{"bool as true and succeeds", args{true}, false},
		{"bool as false and succeeds", args{false}, false},
		{"int check succeeds", args{123}, false},
		{"int as 0 check succeeds", args{0}, false},
		{"float check succeeds", args{123.123}, false},
		{"float as 0 check succeeds", args{0.0}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := required(tt.args.v); got != tt.want {
				t.Errorf("required() = %v, want %v", got, tt.want)
			}
		})
	}
}
