package year2023

import (
	"advent-of-code/helpers"
	"testing"
)

func TestHauntedWasteland(t *testing.T) {
	input := helpers.ReadFile("./day8/input.txt")

	out := HauntedWasteland(input)
	var want int = 6

	if out != want {
		t.Errorf("HauntedWasteland() = %v, want %v", out, want)
	}
}

func TestHauntedWastelandP2(t *testing.T) {
	input := helpers.ReadFile("./day8/input2.txt")

	out := HauntedWastelandP2(input)
	var want int = 6

	if out != want {
		t.Errorf("HauntedWastelandP2() = %v, want %v", out, want)
	}
}
