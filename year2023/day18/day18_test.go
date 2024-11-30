package year2023

import (
	"testing"

	"github.com/vlone310/advent-of-code/helpers"
)

func TestLavaductLagoon(t *testing.T) {
	input := helpers.ReadFile("./input.txt")

	out := LavaductLagoon(input)
	var want int = 62

	if out != want {
		t.Errorf("LavaductLagoon() = %v, want %v", out, want)
	}
}

func TestLavaductLagoonP2(t *testing.T) {
	input := helpers.ReadFile("./input.txt")

	out := LavaductLagoon(input)
	var want int = 62

	if out != want {
		t.Errorf("LavaductLagoonP2() = %v, want %v", out, want)
	}
}
