package year2024

import (
	"testing"

	"github.com/vlone310/advent-of-code/helpers"
)

func TestClawContraption(t *testing.T) {
	input := helpers.ReadFile("./input.txt")
	out := clawContraption(input)
	want := 480

	if out != want {
		t.Errorf("clawContraption() = %v, want = %v", out, want)
	}
}

func TestGardenGroupsP2(t *testing.T) {
	input := helpers.ReadFile("./input.txt")
	out := clawContraptionP2(input)
	want := 875318608908

	if out != want {
		t.Errorf("clawContraptionP2() = %v, want = %v", out, want)
	}
}
