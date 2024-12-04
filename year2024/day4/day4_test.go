package year2024

import (
	"testing"

	"github.com/vlone310/advent-of-code/helpers"
)

func TestCeresSearch(t *testing.T) {
	input := helpers.ReadFile("./input.txt")
	out := CeresSearch(input)
	want := 18

	if out != want {
		t.Errorf("CeresSearch() = %v, want = %v", out, want)
	}
}

func TestCeresSearchP2(t *testing.T) {
	input := helpers.ReadFile("./input.txt")
	out := CeresSearchP2(input)
	want := 9

	if out != want {
		t.Errorf("CeresSearchP2() = %v, want = %v", out, want)
	}
}
