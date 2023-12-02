package part_two

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

var digitFromString = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

var searchDigit = map[string][]string{
	"o": {"one"},
	"t": {"two", "three"},
	"f": {"four", "five"},
	"s": {"six", "seven"},
	"e": {"eight"},
	"n": {"nine"},
}

func tryDigit(char string, word string) int {
	digit, err := strconv.Atoi(string(char))

	if err == nil {
		return digit
	}

	value := searchDigit[char]

	if value == nil {
		return -1
	}

	for _, v := range value {
		want := regexp.MustCompile(`^` + v)

		found := want.FindString(word)

		if found == "" {
			digit = -1
			continue
		}

		digit = digitFromString[v]
		break
	}

	return digit
}

func GetResult(input []string) (string, error) {
	if len(input) == 0 {
		return "", errors.New("empty input")
	}

	sum := 0
	for _, line := range input {
		var lhs int
		var rhs int
		var digits []int

		for i := 0; i < len(line); i++ {
			digit := tryDigit(string(line[i]), line[i:])

			if digit == -1 {
				continue
			}

			digits = append(digits, digit)
		}

		if len(digits) > 1 {
			lhs = digits[0]
			rhs = digits[len(digits)-1]
		} else {
			lhs = digits[0]
			rhs = digits[0]
		}

		number, _ := strconv.Atoi(fmt.Sprintf("%v%v", lhs, rhs))
		sum = sum + number
	}

	return fmt.Sprintf("%v", sum), nil
}
