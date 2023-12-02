package part_one

import (
	"regexp"
	"testing"
)

func TestGetResult(t *testing.T) {
	input := "Part 1"
	want := regexp.MustCompile(`\b` + input + `\b`)

	result, err := GetResult("Part 1")

	if !want.MatchString(result) || err != nil {
		t.Fatalf(`GetResult("Part 1") = %q, %v, want match for %#q, nil`, result, err, want)
	}
}

func TestGetResultEmpty(t *testing.T) {
	result, err := GetResult("")

	if result != "" || err == nil {
		t.Fatalf(`GetResult("") = %q, %v, want ""`, result, err)
	}
}
