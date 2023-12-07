package year2023

import (
	"advent-of-code/helpers"
	"testing"
)

func TestCamelCards(t *testing.T) {
	input := helpers.ReadFile("./day7/input.txt")

	out := CamelCards(input)
	var want int = 6440

	if out != want {
		t.Errorf("CamelCards() = %v, want %v", out, want)
	}
}

func TestCamelCardsP2(t *testing.T) {
	input := helpers.ReadFile("./day7/input.txt")

	out := CamelCardsP2(input)
	var want int = 5905

	if out != want {
		t.Errorf("CamelCardsP2() = %v, want %v", out, want)
	}
}
