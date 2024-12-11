package year2024

import (
	"testing"

	"github.com/vlone310/advent-of-code/helpers"
)

func TestPlutonianPebbles(t *testing.T) {
	input := helpers.ReadFile("./input.txt")
	out := plutonianPebbles(input, 25)
	want := 55312

	if out != want {
		t.Errorf("plutonianPebbles() = %v, want = %v", out, want)
	}
}

func TestPlutonianPebblesP2(t *testing.T) {
	input := helpers.ReadFile("./input.txt")
	out := plutonianPebbles(input, 75)
	want := 65601038650482

	if out != want {
		t.Errorf("plutonianPebbles() = %v, want = %v", out, want)
	}
}
