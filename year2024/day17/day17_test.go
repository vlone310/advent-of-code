package year2024

import (
	"testing"

	"github.com/vlone310/advent-of-code/helpers"
)

func TestChronospatialComputer(t *testing.T) {
	input := helpers.ReadFile("./input.txt")
	c := chronospatialComputer(NewComputer(input))
	want := "4,6,3,5,6,3,5,2,1,0"

	if c.toString() != want {
		t.Errorf("chronospatialComputer() = %v, want = %v", c.toString(), want)
	}
}

func TestChronospatialComputerP2(t *testing.T) {
	input := helpers.ReadFile("./input2.txt")
	out := chronospatialComputerP2(input)
	want := int64(117440)

	if out != want {
		t.Errorf("chronospatialComputerP2() = %v, want = %v", out, want)
	}
}
