package year2024

import (
	"testing"

	"github.com/vlone310/advent-of-code/helpers"
)

func TestMullItOver(t *testing.T) {
	input := helpers.ReadFile("./input.txt")
	out := MullItOver(input)
	want := 161

	if out != want {
		t.Errorf("MullItOver() = %v, want = %v", out, want)
	}
}

func TestMullItOverP2(t *testing.T) {
	input := helpers.ReadFile("./input2.txt")
	out := MullItOverP2(input)
	want := 48

	if out != want {
		t.Errorf("MullItOverP2() = %v, want = %v", out, want)
	}
}
