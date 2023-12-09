package part_one

import (
	"errors"
	"regexp"
	"strconv"
)

func GetResult(input []string) (string, error) {
	if len(input) == 0 {
		return "", errors.New("empty input")
	}

	getSeeds(input[0])

	return input[0], nil
}

func getSeeds(s string) []int {
	numbers := regexp.MustCompile(`\d+`)

	seeds := numbers.FindAllString(s, -1)

	var seedsAsInt []int
	for i := 0; i < len(seeds); i++ {
		n, _ := strconv.Atoi(string(seeds[0]))
		seedsAsInt[i] = n
	}

	return seedsAsInt
}