package utils

func String(errors map[string][]string) string {
	var errs string

	for _, fieldErrors := range errors {
		for _, ff := range fieldErrors {
			if errs != "" {
				errs += ","
			}

			errs += ff
		}
	}

	return errs
}
