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

	fieldType, fieldValue := v.getField(field)

	switch fieldType.String() {
	case "string":
		v.ValidationErrors = str.Required(field, fieldValue.String(), v.ValidationErrors)
	}

	return v
}

func (v *Validator) Length(field string, from int, to int) *Validator {
	initValidator(v)

	fieldValue, _ := v.getField(field)

	v.ValidationErrors = str.Length(field, fieldValue.String(), from, to, v.ValidationErrors)

	return v
}

func (v *Validator) HasErrors() bool {
	initValidator(v)

	if len(v.ValidationErrors) == 0 {
		return false
	}

	return true
}

func (v *Validator) getField(name string) (reflect.Type, reflect.Value) {
	fieldType 	:= reflect.ValueOf(v.Entity).FieldByName(name).Type()
	fieldValue 	:= reflect.ValueOf(v.Entity).FieldByName(name)

	return fieldType, fieldValue;
}
