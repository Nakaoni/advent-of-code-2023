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
	fmt.Println(partNumbers)

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
		var lineAbove bool
		var lineBelow bool

		if i == 0 {
			lineAbove = false
			lineBelow = true
		} else if i == len(matrix)-1 {
			lineAbove = true
			lineBelow = false
		} else {
			lineAbove = true
			lineBelow = true
		}

		multSignPositions := getMultSignPositions(matrix)
		numbers := findNumbersInLine(currentLine)

		for _, number := range searchGear(i, numbers, multSignPositions, currentLine, lineAbove, lineBelow) {
			partNumbers = append(partNumbers, number)
		}
	}

	return partNumbers
}

func getMultSignPositions(matrix [][]string) map[int][]int {
	positions := make(map[int][]int, len(matrix))
	multSign := regexp.MustCompile(`[*]`)

	for i, row := range matrix {
		for j, col := range row {
			if !multSign.MatchString(col) {
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
	lineAbove bool,
	lineBelow bool,
) []int {
	var gears []int

	for _, number := range numbers {
		addAsGear := false

		for i:= 0; i <= number.length; i++ {
			limit := number.startIndex + number.length + 1
			if lineAbove {
				signs, exists := multSigns[matrixIndex - 1]
				if exists {
					for _, sign := range signs {
						if sign >= number.startIndex - 1 && sign <= limit {
							addAsGear = true
							break
						}
					}
				}

				if addAsGear == true {
					break
				}
			}

			if lineBelow {
				signs, exists := multSigns[matrixIndex + 1]
				if exists {
					for _, sign := range signs {
						if sign >= number.startIndex - 1 && sign <= limit {
							addAsGear = true
							break
						}
					}
				}

				if addAsGear == true {
					break
				}
			}

			signs, exists := multSigns[matrixIndex]
			if exists {
				for _, sign := range signs {
					if sign >= number.startIndex && sign <= limit {
						addAsGear = true
						break
					}
				}

				if addAsGear == true {
					break
				}
			}
		}

		if addAsGear {
			gears = append(gears, number.number)
		}
	}

	return gears
}
