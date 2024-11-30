package year2023

import (
	"testing"

	"github.com/vlone310/advent-of-code/helpers"
)

func TestLensLibrary(t *testing.T) {
	input := helpers.ReadFile("./input.txt")

	out := LensLibrary(input)
	var want int = 1320

	if out != want {
		t.Errorf("LensLibrary() = %v, want %v", out, want)
	}
}

func TestLensLibraryP2(t *testing.T) {
	input := helpers.ReadFile("./input.txt")

	out := LensLibraryP2(input)
	var want int = 145

	if out != want {
		t.Errorf("LensLibraryP2() = %v, want %v", out, want)
	}
}
