package year2023

import (
	"testing"

	"github.com/vlone310/advent-of-code/helpers"
)

func TestTheFloorWillBeLava(t *testing.T) {
	input := helpers.ReadFile("./input.txt")

	out := TheFloorWillBeLava(input)
	var want int = 46

	if out != want {
		t.Errorf("TheFloorWillBeLava() = %v, want %v", out, want)
	}
}

func TestTheFloorWillBeLavaP2(t *testing.T) {
	input := helpers.ReadFile("./input.txt")

	out := TheFloorWillBeLavaP2(input)
	var want int = 51

	if out != want {
		t.Errorf("TheFloorWillBeLavaP2() = %v, want %v", out, want)
	}
}
