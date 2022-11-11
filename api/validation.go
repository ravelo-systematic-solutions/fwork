package api

import (
	"encoding/json"
	"github.com/ravelo-systematic-solutions/fwork/exceptions"
	"log"
	"reflect"
	"strconv"
	"strings"
)

const headerTag = "header"
const queryTag = "query"
const validationTag = "validate"

type ValidationRuleError struct {
	Field      string
	FailedTags []string
	Value      string
}

type ValidationRule func(cv any) bool

//required is a rule that ensures that the
//value is set. Returns true if the value
//failed validation, false otherwise.
func required(v any) bool {
	vType := reflect.TypeOf(v)
	switch {
	case vType == nil:
		return true
	case vType.Name() == "string" && v == "":
		return true
	}
	return false
}

//ValidateQuery extract & validates the Get parameters
//from a request using the "query" tag for the name of the
//fields and the "validate" tag for validation rules.
func (s *scope) ValidateQuery(payload interface{}) error {
	dataType := reflect.TypeOf(payload).Elem()
	dataValue := reflect.ValueOf(payload).Elem()
	ex := exceptions.NewBuilder()

	for i := 0; i < dataType.NumField(); i++ {
		field := dataType.Field(i)

		name := field.Name
		fieldValue := dataValue.Field(i)
		value := extractValue(fieldValue)
		tagKey := field.Tag.Get(queryTag)
		val := s.r.URL.Query().Get(tagKey)
		tagValues := strings.Split(field.Tag.Get(validationTag), ",")

		if fieldValue.IsValid() {
			switch fieldValue.Type().String() {
			case "string":
				fieldValue.SetString(val)
			case "int":
				valInt, _ := strconv.ParseInt(val, 10, 64)
				fieldValue.SetInt(valInt)
			case "float32":
				valFloat, _ := strconv.ParseFloat(val, 32)
				fieldValue.SetFloat(valFloat)
			case "float64":
				valFloat, _ := strconv.ParseFloat(val, 64)
				fieldValue.SetFloat(valFloat)
			case "bool":
				valBool, _ := strconv.ParseBool(val)
				fieldValue.SetBool(valBool)
			}
		}

		if len(tagValues) == 0 {
			continue
		}

		for i := 0; i < len(tagValues); i++ {
			tagValue := tagValues[i]
			switch {
			case tagValue == "required" && required(val):
				ex.Include(exceptions.Data{
					Name:  name,
					Tag:   tagValue,
					Value: value,
				})
			}
		}
	}

	if ex.IsEmpty() {
		return nil
	}

	return ex.Exception()
}

//ValidateJsonBody extract & validates the body from a request
//using the "json" tag for the name of the fields and the
//"validate" tag for validation rules.
func (s *scope) ValidateJsonBody(payload interface{}) error {

	err := json.NewDecoder(s.r.Body).Decode(payload)
	if err != nil {
		log.Panicf("fail to decode money")
	}

	dataType := reflect.TypeOf(payload).Elem()
	dataValue := reflect.ValueOf(payload).Elem()
	ex := exceptions.NewBuilder()

	for i := 0; i < dataType.NumField(); i++ {

		field := dataType.Field(i)
		name := dataType.Field(i).Name
		fieldValue := dataValue.Field(i)
		value := extractValue(fieldValue)
		tags := field.Tag.Get(validationTag)
		tagValues := strings.Split(tags, ",")

		if len(tagValues) == 0 {
			continue
		}

		for i := 0; i < len(tagValues); i++ {
			tagValue := tagValues[i]
			switch {
			case tagValue == "required":
				if required(value) {
					ex.Include(exceptions.Data{
						Name:  name,
						Tag:   tagValue,
						Value: value,
					})
				}
			}
		}
	}

	if ex.IsEmpty() {
		return nil
	}

	return ex.Exception()
}

//ValidateHeaders extract & validates a request header
//using the "header" tag for the name of the fields and the
//"validate" tag for validation rules.
func (s *scope) ValidateHeaders(payload interface{}) error {
	dataType := reflect.TypeOf(payload).Elem()
	dataValue := reflect.ValueOf(payload).Elem()
	ex := exceptions.NewBuilder()

	for i := 0; i < dataType.NumField(); i++ {
		field := dataType.Field(i)

		name := field.Name
		fieldValue := dataValue.Field(i)
		tagKey := field.Tag.Get(headerTag)
		value := s.r.Header.Get(tagKey)
		tagValues := strings.Split(field.Tag.Get(validationTag), ",")

		if fieldValue.IsValid() {
			switch fieldValue.Type().String() {
			case "string":
				fieldValue.SetString(value)
			case "int":
				valInt, _ := strconv.ParseInt(value, 10, 64)
				fieldValue.SetInt(valInt)
			case "float32":
				valFloat, _ := strconv.ParseFloat(value, 32)
				fieldValue.SetFloat(valFloat)
			case "float64":
				valFloat, _ := strconv.ParseFloat(value, 64)
				fieldValue.SetFloat(valFloat)
			case "bool":
				valBool, _ := strconv.ParseBool(value)
				fieldValue.SetBool(valBool)
			}
		}

		if len(tagValues) == 0 {
			continue
		}

		for i := 0; i < len(tagValues); i++ {
			tagValue := tagValues[i]
			switch {
			case tagValue == "required" && required(value):
				ex.Include(exceptions.Data{
					Name:  name,
					Tag:   tagValue,
					Value: value,
				})
			}
		}
	}

	if ex.IsEmpty() {
		return nil
	}

	return ex.Exception()
}
