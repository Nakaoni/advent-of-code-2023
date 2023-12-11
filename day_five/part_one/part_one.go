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

	seeds := getSeeds(input[0])
	mappingRange := getMaps(input[2:])

	return input[0], nil
}

func getSeeds(s string) []int {
	numbers := regexp.MustCompile(`\d+`)

	seeds := numbers.FindAllString(s, -1)

	seedsAsInt := make([]int, len(seeds))
	for i := 0; i < len(seeds); i++ {
		n, _ := strconv.Atoi(string(seeds[0]))
		seedsAsInt[i] = n
	}

	return seedsAsInt
}

type Range struct {
	dest int
	source int
	length int
}

type Map struct {
	to string
	rg []Range
}

func getMaps(input []string) map[string]Map {
	mapping := make(map[string]Map)

	mappingString := regexp.MustCompile(`(\w+)-to-(\w+)\smap:`)
	rg := regexp.MustCompile(`(\d+)\s(\d+)\s(\d+)`)

	var curr string
	var rgs []Range
	for _, line := range input {
		s := mappingString.FindStringSubmatch(line)
		r := rg.FindStringSubmatch(line)

		if line == "" {
			mapping[curr] = Map{
				to: mapping[curr].to,
				rg: rgs,
			}

			curr = ""
			rgs = []Range{}

			continue
		}

		if len(s) != 0 {
			curr = s[1]
			mapping[curr] = Map{
				to: s[2],
				rg: nil,
			}

			continue
		}

		if len(r) != 0 {
			dest, _ := strconv.Atoi(r[1])
			src, _ := strconv.Atoi(r[2])
			lg, _ := strconv.Atoi(r[3])

			rgs = append(rgs, Range{
				dest: dest,
				source: src,
				length: lg,
			})

			continue
		}
	}

	return mapping
}
