package main

import "testing"

func TestValidator_Required(t *testing.T) {
	test := TestStruct{}

	validator := Validator{Entity: test}
	validator.Required("Name")

	if len(validator.ValidationErrors["Name"]) < 1 {
		t.Error("Required")
	}
}

type TestStruct struct {
	ID int
	Name string
	Title string
}
