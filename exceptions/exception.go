package exceptions

import "fmt"

//ExceptionBlueprint declares what it is
//needed to create custom exceptions
//for an applications
type ExceptionBlueprint interface {
	Error() string
}

//Data stores information about an exception
type Data struct {
	Name  string `json:"name" bson:"name" xml:"name" yaml:"name" asn1:"utf8"`
	Tag   string `json:"tag" bson:"tag" xml:"tag" yaml:"tag" asn1:"utf8"`
	Value any    `json:"value" bson:"value" xml:"value" yaml:"value" asn1:"utf8"`
}

//Exception holds information about expected
//errors in an application
type Exception struct {
	Code    string `json:"code" bson:"code" xml:"code" yaml:"code" asn1:"utf8"`
	Message string `json:"message" bson:"message" xml:"message" yaml:"message" asn1:"utf8"`
	Data    []Data `json:"data,omitempty" bson:"data,omitempty" xml:"data,omitempty" yaml:"data,omitempty" asn1:"utf8"`
}

//Error ensures that the struct
//implements the Error interface
func (e *Exception) Error() string {
	return fmt.Sprintf(
		"E{%v}:M{%v}:P{%v}",
		e.Code,
		e.Message,
		e.Data,
	)
}
