package str

func Length(field string, value string, from int, to int, errors map[string][]string, err string) map[string][]string {
	if len(value) < from || len(value) > to {
		errors[field] = append(errors[field], err)
	}

	return errors
}

func Max(field string, value string, max int, errors map[string][]string, err string) map[string][]string {
	if len(value) > max {
		errors[field] = append(errors[field], err)
	}

	return errors
}

func Min(field string, value string, min int, errors map[string][]string, err string) map[string][]string {
	if len(value) < min {
		errors[field] = append(errors[field], err)
	}

	return errors
}
