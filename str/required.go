package str

func Required(field string, value string, error map[string][]string) map[string][]string {
	if value == "" {
		error[field] = append(error[field], "required")
	}

	return error
}
