package year2024

import (
	"testing"

	"github.com/vlone310/advent-of-code/helpers"
)

func TestResonantCollinearity(t *testing.T) {
	input := helpers.ReadFile("./input.txt")
	out := resonantCollinearity(input)
	want := 14

	if out != want {
		t.Errorf("resonantCollinearity() = %v, want = %v", out, want)
	}
}

func TestResonantCollinearityP2(t *testing.T) {
	input := helpers.ReadFile("./input.txt")
	out := resonantCollinearityP2(input)
	want := 34

	if out != want {
		t.Errorf("resonantCollinearityP2() = %v, want = %v", out, want)
	}
}
