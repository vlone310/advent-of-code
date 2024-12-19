package year2024

import (
	"testing"

	"github.com/vlone310/advent-of-code/helpers"
)

func TestLinenLayout(t *testing.T) {
	input := helpers.ReadFile("./input.txt")
	out := linenLayout(input)
	want := 6

	if out != want {
		t.Errorf("linenLayout() = %v, want = %v", out, want)
	}
}

func TestLinenLayoutP2(t *testing.T) {
	input := helpers.ReadFile("./input.txt")
	out := linenLayoutP2(input)
	want := 16

	if out != want {
		t.Errorf("linenLayoutP2() = %v, want = %v", out, want)
	}
}
