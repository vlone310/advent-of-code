package year2023

import (
	"testing"

	"github.com/vlone310/advent-of-code/helpers"
)

func TestCosmicExpansion(t *testing.T) {
	input := helpers.ReadFile("./input.txt")

	out := CosmicExpansion(input)
	var want int = 374

	if out != want {
		t.Errorf("CosmicExpansion() = %v, want %v", out, want)
	}
}

func TestCosmicExpansionP2(t *testing.T) {
	input := helpers.ReadFile("./input.txt")

	out := CosmicExpansionP2(input)
	var want int = 82000210

	if out != want {
		t.Errorf("CosmicExpansionP2() = %v, want %v", out, want)
	}
}
