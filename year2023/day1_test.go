package year2023

import (
	"advent-of-code/helpers"
	"testing"
)

func TestTrebuchet(t *testing.T) {
	input := helpers.ReadFile("./day1/input.txt")

	out := Trebuchet(input)
	var want float64 = 53080

	if out != want {
		t.Errorf("Trebuchet() = %v, want %v", out, want)
	}
}

func TestTrebuchetP2(t *testing.T) {
	input := helpers.ReadFile("./day1/input.txt")

	out := TrebuchetP2(input)
	var want float64 = 53268

	if out != want {
		t.Errorf("TrebuchetP2() = %v, want %v", out, want)
	}
}
