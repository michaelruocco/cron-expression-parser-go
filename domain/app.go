package domain

func Run(args []string) (string, error) {
	santizedArgs, err := Sanitize(args)
	if err != nil {
		return "", err
	}

	result, err := Parse(santizedArgs)
	if err != nil {
		return "", err
	}

	return Format(result), nil
}
