package api

import (
	"net/http"
	"testing"
)

func TestGenerateEndpointKey(t *testing.T) {
	url := "/some-url"
	type args struct {
		method string
		url    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Get request", args{http.MethodGet, url}, "get-/some-url"},
		{"post request", args{http.MethodPost, url}, "post-/some-url"},
		{"put request", args{http.MethodPut, url}, "put-/some-url"},
		{"delete request", args{http.MethodDelete, url}, "delete-/some-url"},
		{"camel case url", args{http.MethodDelete, "/SoMe-UrL"}, "delete-/some-url"},
		{"camel case method", args{"GeT", url}, "get-/some-url"},
		{"capital case", args{http.MethodDelete, "/SOME-URL"}, "delete-/some-url"},
		{"lower case", args{"delete", url}, "delete-/some-url"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateEndpointKey(tt.args.method, tt.args.url); got != tt.want {
				t.Errorf("GenerateEndpointKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
