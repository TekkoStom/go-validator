package validator

import "testing"

func TestValidator_Required(t *testing.T) {
	test := TestStruct{}

	validator := InitValidator(test)

	validator.Property("Name").
		Required().
		Length(0, 255)
}

type TestStruct struct {
	ID int
	Name string
	Title string
}
