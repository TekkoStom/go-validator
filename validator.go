package validator

import (
	"reflect"

	"github.com/TekkoStom/go-validator/str"
)

type Validator struct {
	Entity           interface{}
	ValidationErrors map[string][]string
}

func initValidator(v *Validator) {
	if v.ValidationErrors == nil {
		v.ValidationErrors = make(map[string][]string)
	}
}

func (v *Validator) Required(field string) *Validator {
	initValidator(v)

	fieldType 	:= reflect.ValueOf(v.Entity).FieldByName(field).Type().String()
	fieldValue 	:= reflect.ValueOf(v.Entity).FieldByName(field)

	switch fieldType {
	case "string":
		v.ValidationErrors = str.Required(field, fieldValue.String(), v.ValidationErrors)
	}

	return v
}

func (v *Validator) Length(field string, from int, to int) *Validator {
	initValidator(v)

	fieldValue := reflect.ValueOf(v.Entity).FieldByName(field)

	v.ValidationErrors = str.Length(field, fieldValue.String(), from, to, v.ValidationErrors)

	return v
}
