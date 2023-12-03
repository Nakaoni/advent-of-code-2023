package part_one

import (
	"errors"
	"fmt"
)

func GetResult(input []string) (string, error) {
	if len(input) == 0 {
		return "", errors.New("empty input")
	}

	matrix := buildMatrix(input)

	for _, v := range matrix {
		fmt.Println(v)
	}

	partNumbers := getPartNumbers()

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
