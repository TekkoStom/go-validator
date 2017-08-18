package validator

import "testing"

func TestValidator_Required(t *testing.T) {
	test := TestStruct{}

	validator := Validator{Entity: test}
	validator.Required("Name")

	if len(validator.ValidationErrors["Name"]) < 1 {
		t.Error("Required")
	}
}

func TestValidator_HasErrors(t *testing.T) {
	test := TestStruct{}

	validator := Validator{Entity: test}
	validator.Required("Name")

	if !validator.HasErrors() {
		t.Error("Should not return false")
	}
}

func TestValidator_HasNotErrors(t *testing.T) {
	test := TestStruct{}

	validator := Validator{Entity: test}

	if validator.HasErrors() {
		t.Error("Should not return true")
	}
}

type TestStruct struct {
	ID int
	Name string
	Title string
}
