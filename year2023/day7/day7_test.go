package year2023

import (
	"testing"

	"github.com/vlone310/advent-of-code/helpers"
)

func TestCamelCards(t *testing.T) {
	input := helpers.ReadFile("./input.txt")

	out := CamelCards(input)
	var want int = 6440

	if out != want {
		t.Errorf("CamelCards() = %v, want %v", out, want)
	}
}

func TestCamelCardsP2(t *testing.T) {
	input := helpers.ReadFile("./input.txt")

	out := CamelCardsP2(input)
	var want int = 5905

	if out != want {
		t.Errorf("CamelCardsP2() = %v, want %v", out, want)
	}
}
