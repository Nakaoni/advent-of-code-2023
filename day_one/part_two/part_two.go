package part_two

import (
	"errors"
	"fmt"
	"strconv"
)

func GetResult(input []string) (string, error) {
	if len(input) == 0 {
		return "", errors.New("empty input")
	}

	sum := 0
	for _, line := range input {
		var lhs int
		var rhs int

		for i := 0; i < len(line); i++ {
			digit, err := strconv.Atoi(string(line[i]))

			if err != nil {
				continue
			}

			lhs = digit
			break
		}

		for i := len(line) - 1; i >= 0; i-- {
			digit, err := strconv.Atoi(string(line[i]))

			if err != nil {
				continue
			}

			rhs = digit
			break
		}

		number, _ := strconv.Atoi(fmt.Sprintf("%v%v", lhs, rhs))
		sum = sum + number
	}

	return fmt.Sprintf("%v", sum), nil
}
