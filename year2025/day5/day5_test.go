package year2025

import (
	"testing"

	"github.com/vlone310/advent-of-code/helpers"
)

func TestCafeteria(t *testing.T) {
	input := helpers.ReadFile("./input.txt")

	out := Cafeteria(input)
	want := 3

	if out != want {
		t.Errorf("Cafeteria() = out %v, want = %v", out, want)
	}
}

func TestCafeteriaP2(t *testing.T) {
	input := helpers.ReadFile("./input.txt")

	out := CafeteriaP2(input)
	want := 14

	if out != want {
		t.Errorf("CafeteriaP2() = out %v, want = %v", out, want)
	}
}
