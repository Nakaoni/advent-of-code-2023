package part_one

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

var target = Bag{
	red:   12,
	green: 13,
	blue:  14,
}

func GetResult(input []string) (string, error) {
	if len(input) == 0 {
		return "", errors.New("empty input")
	}

	var eligibleGameIDs []int

	for _, line := range input {
		id, err := parseAndTestGame(line)

		if err != nil {
			continue
		}

		eligibleGameIDs = append(eligibleGameIDs, id)
	}

	sum := 0
	for _, id := range eligibleGameIDs {
		sum += id
	}

	return strconv.Itoa(sum), nil
}

func parseAndTestGame(line string) (int, error) {
	id, err := getGameId(line)
	if err != nil {
		return 0, errors.New("error parsing")
	}

	if isEligible(line) {
		return id, nil
	}

	return 0, nil
}

func getGameId(line string) (int, error) {
	gameId := regexp.MustCompile(`^Game\s(\d+):`)

	match := gameId.FindStringSubmatch(line)
	if len(match) <= 1 || len(match) > 2 {
		return 0, errors.New("error parsing")
	}

	id, err := strconv.Atoi(match[1])

	if err != nil {
		return 0, errors.New("error parsing")
	}

	return id, nil
}

func isEligible(line string) bool {
	setsPattern := regexp.MustCompile(`.*:(.*)`)

	sets := setsPattern.FindStringSubmatch(line)
	if len(sets) <= 1 || len(sets) > 2 {
		return false
	}

	for _, set := range strings.Split(sets[1], ";") {
		for _, s := range strings.Split(set, ",") {
			pattern := regexp.MustCompile(`\s(\d+)\s(.*)`)

			value := pattern.FindStringSubmatch(s)
			n, _ := strconv.Atoi(value[1])
			switch value[2] {
			case "red":
				if n > target.red {
					return false
				}
			case "green":
				if n > target.green {
					return false
				}
			case "blue":
				if n > target.blue {
					return false
				}
			}
		}
	}

	return true
}
