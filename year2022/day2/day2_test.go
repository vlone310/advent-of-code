package year2022

import (
	"testing"

	"github.com/vlone310/advent-of-code/helpers"
)

func TestRockPaperScissors(t *testing.T) {
	input := helpers.ReadFile("./input.txt")

	out := RockPaperScissors(input)
	var want float64 = 12679

	if out != want {
		t.Errorf("RockPaperScissors() = %v, want %v", out, want)
	}
}

func TestRockPaperScissorsP2(t *testing.T) {
	input := helpers.ReadFile("./input.txt")

	out := RockPaperScissorsP2(input)
	var want float64 = 14470

	if out != want {
		t.Errorf("RockPaperScissorsP2() = %v, want %v", out, want)
	}
}
