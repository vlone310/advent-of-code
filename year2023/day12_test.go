package year2023

import (
	"advent-of-code/helpers"
	"testing"
)

func TestHotSprings(t *testing.T) {
	input := helpers.ReadFile("./day12/input.txt")

	out := HotSprings(input)
	var want int = 21

	if out != want {
		t.Errorf("HotSprings() = %v, want %v", out, want)
	}
}

func TestHotSpringsP2(t *testing.T) {
	input := helpers.ReadFile("./day12/input.txt")

	out := HotSpringsP2(input)
	var want int = 525152

	if out != want {
		t.Errorf("HotSpringsP2() = %v, want %v", out, want)
	}
}
