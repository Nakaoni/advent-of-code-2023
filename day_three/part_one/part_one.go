package part_one

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

var symbol = regexp.MustCompile(`[*\/%=\-+$@#&]`)

func GetResult(input []string) (string, error) {
	if len(input) == 0 {
		return "", errors.New("empty input")
	}

	matrix := buildMatrix(input)

	partNumbers := getPartNumbers(matrix)

	sum := 0
	for _, number := range partNumbers {
		sum += number
	}

	return strconv.Itoa(sum), nil
}

func buildMatrix(input []string) [][]string {
	matrix := make([][]string, len(input[0]), len(input))

	for i, line := range input {
		for j := 0; j < len(line); j++ {
			matrix[i] = append(matrix[i], string(line[j]))
		}
	}

	return matrix
}

func getPartNumbers(matrix [][]string) []int {
	var partNumbers []int

	for i := 0; i < len(matrix); i++ {
		currentLine := matrix[i]
		var lineAbove []string
		var lineBelow []string

		if i == 0 {
			lineAbove = nil
			lineBelow = matrix[i+1]
		} else if i == len(matrix)-1 {
			lineAbove = matrix[i-1]
			lineBelow = nil
		} else {
			lineAbove = matrix[i-1]
			lineBelow = matrix[i+1]
		}

		numbers := findNumbersInLine(currentLine)

		for _, number := range searchPartNumbersForLine(numbers, currentLine, lineAbove, lineBelow) {
			partNumbers = append(partNumbers, number)
		}
	}

	return partNumbers
}

type NumberInLIne struct {
	number     int
	startIndex int
	length     int
}

func findNumbersInLine(line []string) []NumberInLIne {
	var list []NumberInLIne
	var number string
	var startIndex int
	length := 0

	for i := 0; i < len(line); i++ {
		char := line[i]
		_, err := strconv.Atoi(char)

		if err != nil {
			if (char == "." || symbol.MatchString(char)) && number != "" {
				n, _ := strconv.Atoi(number)
				list = append(list, NumberInLIne{
					number:     n,
					startIndex: startIndex,
					length:     length,
				})
				number = ""
				length = 0
			}

			continue
		}

		// if number is empty then set startIndex
		if number == "" {
			startIndex = i
		}

		// build number as string
		number = fmt.Sprintf("%v%v", number, char)

		// length + 1
		length += 1
	}

	return list
}

func searchPartNumbersForLine(
	numbers []NumberInLIne,
	currentLine []string,
	lineAbove []string,
	lineBelow []string,
) []int {
	var partNumbers []int

	for _, n := range numbers {
		if n.startIndex > 0 {
			char := currentLine[n.startIndex-1]
			if symbol.MatchString(char) {
				partNumbers = append(partNumbers, n.number)
				continue
			}
		}

		if n.startIndex+n.length < len(currentLine)-1 {
			char := currentLine[n.startIndex+n.length]
			if symbol.MatchString(char) {
				partNumbers = append(partNumbers, n.number)
				continue
			}
		}

		i := n.startIndex

		if i > 0 {
			i = i - 1
		}

		length := i + n.length
		if length  >= len(currentLine) - 1 {
			length = length - 1
		}
		for ; i <= length + 1; i++ {

			if lineAbove != nil {
				char := lineAbove[i]
				if symbol.MatchString(char) {
					partNumbers = append(partNumbers, n.number)
					continue
				}
			}
			if lineBelow != nil {
				char := lineBelow[i]
				if symbol.MatchString(char) {
					partNumbers = append(partNumbers, n.number)
					continue

				}
			}
		}
	}

	return partNumbers
}
