package year2024

import (
	"testing"

	"github.com/vlone310/advent-of-code/helpers"
)

func TestRamRun(t *testing.T) {
	input := helpers.ReadFile("./input.txt")
	out := ramRun(input, 6, 12)
	want := 22

	if out != want {
		t.Errorf("ramRun() = %v, want = %v", out, want)
	}
}

func TestRamRunP2(t *testing.T) {
	input := helpers.ReadFile("./input.txt")
	out := ramRunP2(input, 6)
	want := "6,1"

	if out != want {
		t.Errorf("ramRunP2() = %v, want = %v", out, want)
	}
}
