package part_two

import (
	"github.com/Nakaoni/advent-of-code-2023/utils"
	"regexp"
	"testing"
)

func TestGetResult(t *testing.T) {
	expected := "467835"
	want := regexp.MustCompile(`\b` + expected + `\b`)

	input, err := utils.GetFileContent("../assets/example.txt")
	if err != nil {
	}
	result, err := GetResult(input)

	if !want.MatchString(result) || err != nil {
		t.Fatalf(`GetResult(input) = %q, %v, want match for %#q, nil`, result, err, want)
	}
}

func TestGetResultEmpty(t *testing.T) {
	result, err := GetResult([]string{})

	if result != "" || err == nil {
		t.Fatalf(`GetResult("") = %q, %v, want ""`, result, err)
	}
}
