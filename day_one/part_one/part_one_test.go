package part_one

import (
	"github.com/Nakaoni/advent-of-code-2023/template/utils"
	"regexp"
	"testing"
)

func TestGetResult(t *testing.T) {
	expected := "142"
	want := regexp.MustCompile(`\b` + expected + `\b`)

	input, err := utils.GetFileContent("../assets/example.txt")
	if err != nil {
		t.Fatalf("error: %v", err)
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
