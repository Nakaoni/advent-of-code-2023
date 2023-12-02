package part_two

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type Bag struct {
	red   int
	green int
	blue  int
}

func GetResult(input []string) (string, error) {
	if len(input) == 0 {
		return "", errors.New("empty input")
	}

	var powers []int

	for _, line := range input {
		power, err := parseAndTestGame(line)

		if err != nil {
			continue
		}

		powers = append(powers, power)
	}

	sum := 0
	for _, p := range powers {
		sum += p
	}

	return strconv.Itoa(sum), nil
}

func parseAndTestGame(line string) (int, error) {
	bag := getBagForGame(line)


	return bag.red * bag.green * bag.blue, nil
}

func getBagForGame(line string) Bag {
	bag := Bag{
		red: 0,
		green: 0,
		blue: 0,
	}

	setsPattern := regexp.MustCompile(`.*:(.*)`)

	sets := setsPattern.FindStringSubmatch(line)
	if len(sets) <= 1 || len(sets) > 2 {
		panic("error regexp")
	}

	for _, set := range strings.Split(sets[1], ";") {
		for _, s := range strings.Split(set, ",") {
			pattern := regexp.MustCompile(`\s(\d+)\s(.*)`)

			value := pattern.FindStringSubmatch(s)
			n, _ := strconv.Atoi(value[1])
			switch value[2] {
			case "red":
				if n > bag.red {
					bag.red = n
				}
			case "green":
				if n > bag.green {
					bag.green = n
				}
			case "blue":
				if n > bag.blue {
					bag.blue = n
				}
			}
		}
	}

	return bag
}
