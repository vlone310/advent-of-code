package year2023

import (
	"testing"

	"github.com/vlone310/advent-of-code/helpers"
)

func TestScratchcards(t *testing.T) {
	input := helpers.ReadFile("./input.txt")

	out := Scratchcards(input)
	var want int64 = 13

	if out != want {
		t.Errorf("Scratchcards() = %v, want %v", out, want)
	}
}

func TestScratchcardsP2(t *testing.T) {
	input := helpers.ReadFile("./input.txt")

	out := ScratchcardsP2(input)
	var want int = 30

	if out != want {
		t.Errorf("ScratchcardsP2() = %v, want %v", out, want)
	}
}
