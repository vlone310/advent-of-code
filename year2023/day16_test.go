package year2023

import (
	"advent-of-code/helpers"
	"testing"
)

func TestTheFloorWillBeLava(t *testing.T) {
	input := helpers.ReadFile("./day16/input.txt")

	out := TheFloorWillBeLava(input)
	var want int = 46

	if out != want {
		t.Errorf("TheFloorWillBeLava() = %v, want %v", out, want)
	}
}

func TestTheFloorWillBeLavaP2(t *testing.T) {
	input := helpers.ReadFile("./day16/input.txt")

	out := TheFloorWillBeLavaP2(input)
	var want int = 51

	if out != want {
		t.Errorf("TheFloorWillBeLavaP2() = %v, want %v", out, want)
	}
}
