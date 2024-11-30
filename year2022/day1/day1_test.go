package year2022

import (
	"testing"

	"github.com/vlone310/advent-of-code/helpers"
)

func TestCalorieCounting(t *testing.T) {
	input := helpers.ReadFile("./input.txt")

	out := CalorieCounting(input)
	var want float64 = 75622

	if out != want {
		t.Errorf("CalorieCounting() = %v, want %v", out, want)
	}
}

func TestCalorieCountingP2(t *testing.T) {
	input := helpers.ReadFile("./input.txt")

	out := CalorieCountingP2(input)
	var want float64 = 213159

	if out != want {
		t.Errorf("CalorieCountingP2() = %v, want %v", out, want)
	}
}
