package year2023

import (
	"advent-of-code/helpers"
	"testing"
)

func TestLensLibrary(t *testing.T) {
	input := helpers.ReadFile("./day15/input.txt")

	out := LensLibrary(input)
	var want int = 1320

	if out != want {
		t.Errorf("LensLibrary() = %v, want %v", out, want)
	}
}

func TestLensLibraryP2(t *testing.T) {
	input := helpers.ReadFile("./day15/input.txt")

	out := LensLibraryP2(input)
	var want int = 145

	if out != want {
		t.Errorf("LensLibraryP2() = %v, want %v", out, want)
	}
}
