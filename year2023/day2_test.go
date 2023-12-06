package year2023

import (
	"advent-of-code/helpers"
	"testing"
)

func TestCubeConundrum(t *testing.T) {
	input := helpers.ReadFile("./day2/input.txt")

	out := CubeConundrum(input)
	var want int64 = 8

	if out != want {
		t.Errorf("CubeConundrum() = %v, want %v", out, want)
	}
}

func TestCubeConundrumP2(t *testing.T) {
	input := helpers.ReadFile("./day2/input.txt")

	out := CubeConundrumP2(input)
	var want float64 = 2286

	if out != want {
		t.Errorf("CubeConundrumP2() = %v, want %v", out, want)
	}
}
