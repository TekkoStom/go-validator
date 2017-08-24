package str

import "github.com/pborman/uuid"

func ValidUUID(field string, value string, errors map[string][]string, err string) map[string][]string {
	if value != "" && uuid.Parse(value) == nil {
		errors[field] = append(errors[field], err)
	}

	return errors
}
