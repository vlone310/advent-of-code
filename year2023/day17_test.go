package year2023

import (
	"advent-of-code/helpers"
	"testing"
)

func TestClumsyCrucible(t *testing.T) {
	input := helpers.ReadFile("./day16/input.txt")

	out := ClumsyCrucible(input)
	var want int = 102

	if out != want {
		t.Errorf("ClumsyCrucible() = %v, want %v", out, want)
	}
}

func TestClumsyCrucibleP2(t *testing.T) {
	input := helpers.ReadFile("./day16/input.txt")

	out := ClumsyCrucible(input)
	var want int = 102

	if out != want {
		t.Errorf("ClumsyCrucibleP2() = %v, want %v", out, want)
	}
}
