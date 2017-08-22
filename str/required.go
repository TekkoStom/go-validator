package str

func Required(field string, value string, errors map[string][]string, err string) map[string][]string {
	if value == "" {
		errors[field] = append(errors[field], err)
	}

	return errors
}
