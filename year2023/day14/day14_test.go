package year2023

import (
	"testing"

	"github.com/vlone310/advent-of-code/helpers"
)

func TestParabolicReflectorDish(t *testing.T) {
	input := helpers.ReadFile("./input.txt")

	out := ParabolicReflectorDish(input)
	var want int = 136

	if out != want {
		t.Errorf("ParabolicReflectorDish() = %v, want %v", out, want)
	}
}

func TestParabolicReflectorDishP2(t *testing.T) {
	input := helpers.ReadFile("./input.txt")

	out := ParabolicReflectorDishP2(input)
	var want int = 64

	if out != want {
		t.Errorf("ParabolicReflectorDishP2() = %v, want %v", out, want)
	}
}
