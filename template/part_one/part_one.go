package part_one

import "errors"

func GetResult(input []string) (string, error) {
	if len(input) == 0 {
		return "", errors.New("empty input")
	}

	return input[0], nil
}
