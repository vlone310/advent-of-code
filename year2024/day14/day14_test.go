package year2024

import (
	"testing"

	"github.com/vlone310/advent-of-code/helpers"
)

func TestRestrooRedoubt(t *testing.T) {
	input := helpers.ReadFile("./input.txt")
	out := restroomRedoubt(input, 101, 103, 100)
	want := 21

	if out != want {
		t.Errorf("restroomRedoubt() = %v, want = %v", out, want)
	}
}

func TestRestrooRedoubtP2(t *testing.T) {
	input := helpers.ReadFile("./input2.txt")
	out := restroomRedoubtP2(input, 101, 103)
	want := 6752

	if out != want {
		t.Errorf("restroomRedoubtP2() = %v, want = %v", out, want)
	}
}
