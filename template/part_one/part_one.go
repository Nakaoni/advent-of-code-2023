package part_one

import "errors"

func GetResult(input string) (string, error) {
	if input == "" {
		return "", errors.New("empty input")
	}

	return input, nil
}
