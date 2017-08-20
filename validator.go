package validator

import (
	"reflect"

	"github.com/TekkoStom/go-validator/str"
)

type Validator struct {
	Entity           interface{}
	ValidationErrors map[string][]string
}

var v *Validator
var prop string

func InitValidator(entity interface{}) *Validator {
	v.Entity 			= entity
	v.ValidationErrors 	= make(map[string][]string)

	return v
}

func (v *Validator) Property(property string) *Validator {
	prop = property

	return v
}

func (v *Validator) Required() *Validator {
	fieldType, fieldValue := v.getField()

	switch fieldType.String() {
	case "string":
		v.ValidationErrors = str.Required(prop, fieldValue.String(), v.ValidationErrors)
	}

	return v
}

func (v *Validator) Length(from int, to int) *Validator {
	fieldValue, _ := v.getField()

	v.ValidationErrors = str.Length(prop, fieldValue.String(), from, to, v.ValidationErrors)

	return v
}

func (v *Validator) HasErrors() bool {
	if len(v.ValidationErrors) == 0 {
		return false
	}

	return true
}

func (v *Validator) getField() (reflect.Type, reflect.Value) {
	fieldType 	:= reflect.ValueOf(v.Entity).Elem().FieldByName(prop).Type()
	fieldValue 	:= reflect.ValueOf(v.Entity).Elem().FieldByName(prop)

	return fieldType, fieldValue;
}
