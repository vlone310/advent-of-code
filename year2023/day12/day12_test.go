package year2023

import (
	"testing"

	"github.com/vlone310/advent-of-code/helpers"
)

func TestHotSprings(t *testing.T) {
	input := helpers.ReadFile("./input.txt")

	out := HotSprings(input)
	var want int = 21

	if out != want {
		t.Errorf("HotSprings() = %v, want %v", out, want)
	}
}

func TestHotSpringsP2(t *testing.T) {
	input := helpers.ReadFile("./input.txt")

	out := HotSpringsP2(input)
	var want int = 525152

	if out != want {
		t.Errorf("HotSpringsP2() = %v, want %v", out, want)
	}
}
