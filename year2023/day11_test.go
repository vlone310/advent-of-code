package year2023

import (
	"advent-of-code/helpers"
	"testing"
)

func TestCosmicExpansion(t *testing.T) {
	input := helpers.ReadFile("./day11/input.txt")

	out := CosmicExpansion(input)
	var want int = 374

	if out != want {
		t.Errorf("CosmicExpansion() = %v, want %v", out, want)
	}
}

func TestCosmicExpansionP2(t *testing.T) {
	input := helpers.ReadFile("./day11/input.txt")

	out := CosmicExpansionP2(input)
	var want int = 82000210

	if out != want {
		t.Errorf("CosmicExpansionP2() = %v, want %v", out, want)
	}
}
