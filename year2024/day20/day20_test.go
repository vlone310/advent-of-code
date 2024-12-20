package year2024

import (
	"testing"

	"github.com/vlone310/advent-of-code/helpers"
)

func TestRaceCondition(t *testing.T) {
	input := helpers.ReadFile("./input.txt")
	out := raceCondition(input, opts{cheatPenalty: 2, cheatTreshold: 10, maxCheat: 2})
	want := 10

	if out != want {
		t.Errorf("raceCondition() = %v, want = %v", out, want)
	}
}

func TestRaceConditionP2(t *testing.T) {
	input := helpers.ReadFile("./input.txt")
	out := raceCondition(input, opts{cheatPenalty: 1, cheatTreshold: 50, maxCheat: 20})
	want := 285

	if out != want {
		t.Errorf("raceConditionP2() = %v, want = %v", out, want)
	}
}
