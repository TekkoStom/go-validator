package str

func In(field string, value string, values []string, errors map[string][]string, err string) map[string][]string {
	if value == "" {
		return errors
	}

	is := false
	for _, val := range values {
		if value == val {
			is = true
		}
	}

	if !is {
		errors[field] = append(errors[field], err)
	}

	return errors
}
