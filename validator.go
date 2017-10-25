package validator

import (
	"reflect"
	"github.com/TekkoStom/go-validator/str"
	"github.com/TekkoStom/go-validator/utils"
)

type Validator struct {
	Entity           interface{}
	ValidationErrors map[string][]string
}

var v *Validator
var prop string

func InitValidator(entity interface{}) *Validator {
	v = &Validator{
		Entity: entity,
		ValidationErrors: make(map[string][]string),
	}

	return v
}

func (v *Validator) Property(property string) *Validator {
	prop = property

	return v
}

func (v *Validator) Required(err ...string) *Validator {
	fieldType, fieldValue := v.getField()

	if err == nil {

	}

	switch fieldType.String() {
	case "string":
		v.ValidationErrors = str.Required(prop, fieldValue.String(), v.ValidationErrors, err[0])
	}

	return v
}

func (v *Validator) Length(from int, to int, err ...string) *Validator {
	_, fieldValue := v.getField()

	v.ValidationErrors = str.Length(prop, fieldValue.String(), from, to, v.ValidationErrors, err[0])

	return v
}

func (v *Validator) Max(max int, err ...string) *Validator {
	_, fieldValue := v.getField()

	v.ValidationErrors = str.Max(prop, fieldValue.String(), max, v.ValidationErrors, err[0])

	return v
}

func (v *Validator) ValidUUID(err ...string) *Validator {
	_, fieldValue := v.getField()

	v.ValidationErrors = str.ValidUUID(prop, fieldValue.String(), v.ValidationErrors, err[0])

	return v
}

func (v *Validator) In(values []string, err ...string) *Validator {
	_, fieldValue := v.getField()

	v.ValidationErrors = str.In(prop, fieldValue.String(), values, v.ValidationErrors, err[0])

	return v
}


func (v *Validator) HasErrors() bool {
	if len(v.ValidationErrors) == 0 {
		return false
	}

	return true
}

func (v *Validator) Errors() []string {
	var errs []string

	for _, fieldErrors := range v.ValidationErrors {
		for _, ff := range fieldErrors {
			errs = append(errs, ff)
		}
	}

	return errs
}

func (v *Validator) String() string {
	return utils.String(v.ValidationErrors)
}

func (v *Validator) getField() (reflect.Type, reflect.Value) {
	fieldType 	:= reflect.ValueOf(v.Entity).Elem().FieldByName(prop).Type()
	fieldValue 	:= reflect.ValueOf(v.Entity).Elem().FieldByName(prop)

	return fieldType, fieldValue
}
