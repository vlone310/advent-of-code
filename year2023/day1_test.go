package year2023

import (
	"advent-of-code/helpers"
	"testing"
)

func TestTrebuchet(t *testing.T) {
	input := helpers.ReadFile("./day1/input.txt")

	out := Trebuchet(input)
	var want float64 = 142

	if out != want {
		t.Errorf("Trebuchet() = %v, want %v", out, want)
	}
}

func TestTrebuchetP2(t *testing.T) {
	input := helpers.ReadFile("./day1/input2.txt")

	out := TrebuchetP2(input)
	var want float64 = 281

	if out != want {
		t.Errorf("TrebuchetP2() = %v, want %v", out, want)
	}
}
