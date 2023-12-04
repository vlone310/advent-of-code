package year2023

import (
	"advent-of-code/helpers"
	"testing"
)

func TestScratchcards(t *testing.T) {
	input := helpers.ReadFile("./day4/input.txt")

	out := Scratchcards(input)
	var want int64 = 27059

	if out != want {
		t.Errorf("Scratchcards() = %v, want %v", out, want)
	}
}

func TestScratchcardsP2(t *testing.T) {
	input := helpers.ReadFile("./day4/input.txt")

	out := ScratchcardsP2(input)
	var want int = 5744979

	if out != want {
		t.Errorf("ScratchcardsP2() = %v, want %v", out, want)
	}
}
