package year2023

import (
	"advent-of-code/helpers"
	"testing"
)

func TestGearRatios(t *testing.T) {
	input := helpers.ReadFile("./day3/input.txt")

	out := GearRatios(input)
	var want int64 = 4361

	if out != want {
		t.Errorf("GearRatios() = %v, want %v", out, want)
	}
}

func TestGearRatiosP2(t *testing.T) {
	input := helpers.ReadFile("./day3/input.txt")

	out := GearRatiosP2(input)
	var want int64 = 467835

	if out != want {
		t.Errorf("GearRatiosP2() = %v, want %v", out, want)
	}
}
