package api

import (
	"fmt"
	"reflect"
	"strings"
)

// GenerateEndpointKey generates a key used to identify urls
// using a request method and url
func GenerateEndpointKey(method, url string) string {
	return strings.ToLower(fmt.Sprintf("%s-%s", method, url))
}

func extractValue(value reflect.Value) any {
	switch value.Type().String() {
	case "string":
		return value.String()
	case "int":
		return value.Int()
	case "float32":
		return value.Float()
	case "float64":
		return value.Float()
	case "bool":
		return value.Bool()
	}
	return nil
}
