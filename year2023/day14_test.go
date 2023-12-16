package year2023

import (
	"advent-of-code/helpers"
	"testing"
)

func TestParabolicReflectorDish(t *testing.T) {
	input := helpers.ReadFile("./day14/input.txt")

	out := ParabolicReflectorDish(input)
	var want int = 136

	if out != want {
		t.Errorf("ParabolicReflectorDish() = %v, want %v", out, want)
	}
}

func TestParabolicReflectorDishP2(t *testing.T) {
	input := helpers.ReadFile("./day14/input.txt")

	out := ParabolicReflectorDishP2(input)
	var want int = 64

	if out != want {
		t.Errorf("ParabolicReflectorDishP2() = %v, want %v", out, want)
	}
}
