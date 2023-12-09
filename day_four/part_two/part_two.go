package part_two

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type Table struct {
	lhs []int
	rhs []int
}

func GetResult(input []string) (string, error) {
	if len(input) == 0 {
		return "", errors.New("empty input")
	}

	cards := make(map[int]int, len(input))
	for idx, line := range input {

		part1, part2 := parse(line)

		qsort(&part1, 0, len(part1) - 1)
		qsort(&part2, 0, len(part2) - 1)

		duplicates := getDuplicates(part1, part2)

		processCards(&cards, idx, len(duplicates))
	}

	sum := 0
	for _, n := range cards {
		sum += n
	}

	return strconv.Itoa(sum), nil
}

func processCards(cards *map[int]int, idx int, len int) {
	c := *cards

	c[idx] += 1

	for i := 0; i < c[idx]; i++ {
		for j := 0; j < len; j++ {
			c[idx + j + 1] += 1
		}
	}
}

func parse(line string) ([]int, []int) {
	card := regexp.MustCompile(`Card\s+\d+:\s`)
	whitespace := regexp.MustCompile(`\s+`)

	cardLess := card.ReplaceAllString(line, "")
	parts := strings.Split(cardLess, "|")

	part1 := strings.Split(whitespace.ReplaceAllString(strings.TrimSpace(parts[0]), ","), ",")
	part2 := strings.Split(whitespace.ReplaceAllString(strings.TrimSpace(parts[1]), ","), ",")


	return convertToInt(part1), convertToInt(part2)
}

func convertToInt(stringTable []string) []int {
	var intTable []int
	for _, s := range stringTable {
		integer, _ := strconv.Atoi(s)
		intTable = append(intTable, integer)
	}

	return intTable
}

func qsort(arr *[]int, lo int, hi int) {
	if lo >= hi {
		return
	}

	pivot := partition(arr, lo, hi)

	qsort(arr, lo, pivot-1)
	qsort(arr, pivot+1, hi)
}

func partition(arr *[]int, lo int, hi int) int {
	a := *arr
	pivot := a[hi]

	idx := lo - 1

	for i := lo; i < hi; i++ {
		if a[i] <= pivot {
			idx++
			a[i], a[idx] = a[idx], a[i]
		}
	}

	idx++;
	a[hi], a[idx] = a[idx], pivot

	return idx
}

func getDuplicates(part1 []int, part2 []int) []int {
	var duplicates []int

	for _, number := range part1 {
		if binarySearch(number, part2) {
			duplicates = append(duplicates, number)
		}
	}

	return duplicates
}

func binarySearch(number int, arr []int) bool {
	if len(arr) == 0 {
		return false
	}

	mid := int(len(arr) / 2)

	if arr[mid] == number {
		return true
	}

	if arr[mid] > number {
		return binarySearch(number, arr[0: mid])
	}

	if arr[mid] < number {
		return binarySearch(number, arr[mid + 1:])
	}

	return false
}
