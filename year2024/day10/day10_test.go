package year2024

import (
	"testing"

	"github.com/vlone310/advent-of-code/helpers"
)

func TestHoofIt(t *testing.T) {
	input := helpers.ReadFile("./input.txt")
	out := hoofIt(input)
	want := 36

	if out != want {
		t.Errorf("hoofIt() = %v, want = %v", out, want)
	}
}

func TestHoofItP2(t *testing.T) {
	input := helpers.ReadFile("./input.txt")
	out := hoofItP2(input)
	want := 81

	if out != want {
		t.Errorf("hoofItP2() = %v, want = %v", out, want)
	}
}
