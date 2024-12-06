package year2024

import (
	"testing"

	"github.com/vlone310/advent-of-code/helpers"
)

func TestGuardGallivant(t *testing.T) {
	input := helpers.ReadFile("./input.txt")
	out := guardGallivant(input)
	want := 41

	if out != want {
		t.Errorf("guardGallivant() = %v, want = %v", out, want)
	}
}

func TestGuardGallivantP2(t *testing.T) {
	input := helpers.ReadFile("./input.txt")
	out := guardGallivantP2(input)
	want := 6

	if out != want {
		t.Errorf("guardGallivantP2() = %v, want = %v", out, want)
	}
}
