package str

import "fmt"

func Length(field string, value string, from int, to int, error map[string][]string) map[string][]string {
	if len(value) < from {
		error[field] = append(error[field], fmt.Sprintf("length.from[%v]", from))
	}

	if len(value) > to {
		error[field] = append(error[field], fmt.Sprintf("length.to[%v]", to))
	}

	return error
}
