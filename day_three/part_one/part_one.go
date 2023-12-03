package part_one

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

func GetResult(input []string) (string, error) {
	if len(input) == 0 {
		return "", errors.New("empty input")
	}

	matrix := buildMatrix(input)

	for _, v := range matrix {
		fmt.Println(v)
	}

	getPartNumbers(matrix)

	return input[0], nil
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
	for i := 0; i < len(matrix); i++ {
		currentLine := matrix[i]
		//var lineAbove []string
		//var lineBelow []string

		//if i == 0 {
		//	lineAbove = nil
		//	lineBelow = matrix[i+1]
		//} else if i == len(matrix)-1 {
		//	lineAbove = matrix[i-1]
		//	lineBelow = nil
		//} else {
		//	lineAbove = matrix[i-1]
		//	lineBelow = matrix[i+1]
		//}

		numbers := findNumbersInLine(currentLine)
		fmt.Println(numbers)
	}

	return []int{}
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

	for i := 0; i < len(line)-1; i++ {
		char := line[i]
		_, err := strconv.Atoi(char)

		if err != nil {
			symbol := regexp.MustCompile(`[*/%=\-+$@#&]`)

			if (char == "." || symbol.MatchString(char)) && number != "" {
				n, _ := strconv.Atoi(number)
				list = append(list, NumberInLIne{
					number:     n,
					startIndex: startIndex,
					length:     length,
				})
				number = ""
				length = 0
				continue
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
