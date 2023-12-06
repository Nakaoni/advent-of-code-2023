package part_two

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

var symbol = regexp.MustCompile(`[\*\/%=\-+$@#&]`)

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
	gears := make(map[int][]int)

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

		multSignPositions := getMultSignPositions(matrix)
		numbers := findNumbersInLine(currentLine)

		searchGear(i, numbers, multSignPositions, currentLine, lineAbove, lineBelow, &gears)
	}

	for _, numbers := range gears {
		if len(numbers) < 2 {
			continue
		}

		mult := 1
		for _, n := range numbers {
			mult = mult * n
		}

		partNumbers = append(partNumbers, mult)
	}

	return partNumbers
}

func getMultSignPositions(matrix [][]string) map[int][]int {
	positions := make(map[int][]int, len(matrix))
	multSign := regexp.MustCompile(`[*]`)
	digit := regexp.MustCompile(`\d`)

	for i, row := range matrix {
		for j, col := range row {
			if !multSign.MatchString(col) {
				continue
			}

			// ...
			// .*.
			// ...

			var hasSurroundingDigits []bool
			if i > 0 {
				if
					digit.MatchString(matrix[i - 1][j - 1]) ||
					digit.MatchString(matrix[i - 1][j]) ||
					digit.MatchString(matrix[i - 1][j + 1]) {
					hasSurroundingDigits = append(hasSurroundingDigits, true)
				}
			}

			if i < len(matrix)- 1 {
				if
					digit.MatchString(matrix[i + 1][j - 1]) ||
					digit.MatchString(matrix[i + 1][j]) ||
					digit.MatchString(matrix[i + 1][j + 1]) {
					hasSurroundingDigits = append(hasSurroundingDigits, true)
				}
			}

			hasSurroundingDigits = append(hasSurroundingDigits, digit.MatchString(matrix[i][j - 1]))
			hasSurroundingDigits = append(hasSurroundingDigits, digit.MatchString(matrix[i][j + 1]))

			sumTrue := 0
			for _, b := range hasSurroundingDigits {
				if b {
					sumTrue++
				}
			}

			if sumTrue < 2 {
				continue
			}

			positions[i] = append(positions[i], j)
		}
	}


	return positions
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

	if number != "" {
		n, _ := strconv.Atoi(number)
		list = append(list, NumberInLIne{
			number:     n,
			startIndex: startIndex,
			length:     length,
		})
	}

	return list
}

func searchGear(
	matrixIndex int,
	numbers []NumberInLIne,
	multSigns map[int][]int,
	currentLine []string,
	lineAbove []string,
	lineBelow []string,
	gears *map[int][]int,
) {

	for _, number := range numbers {

		index := matrixIndex
		limit := number.startIndex + number.length
		start := number.startIndex

		if start > 0 {
			start -= 1
		}

		if lineAbove != nil {
			index -= 1
			signs, exists := multSigns[matrixIndex - 1]
			if exists && checkIfNumberIsSurroundedByMult(index, signs, start, limit) {
				g := *gears
				g[index] = append(g[index], number.number)
				continue
			}
			index = matrixIndex
		}

		if lineBelow != nil {
			index += 1
			signs, exists := multSigns[matrixIndex + 1]

			if exists && checkIfNumberIsSurroundedByMult(index, signs, start, limit) {
				g := *gears
				g[index] = append(g[index], number.number)
				continue
			}
			index = matrixIndex
		}

		signs, exists := multSigns[matrixIndex]
		if exists && checkIfNumberIsSurroundedByMult(index, signs, start, limit) {
			g := *gears
			g[index] = append(g[index], number.number)
			continue
		}
	}
}

func checkIfNumberIsSurroundedByMult(lineIndex int, signsInLine []int, start int, end int) bool {
	for _, sign := range signsInLine {

		if sign >= start && sign <= end {
			return true
		}
	}

	return false
}
