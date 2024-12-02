package year2024

import (
	"testing"

	"github.com/vlone310/advent-of-code/helpers"
)

func TestRedNosedReports(t *testing.T) {
	input := helpers.ReadFile("./input.txt")
	out := RedNosedReports(input)
	want := 2

	if out != want {
		t.Errorf("RedNosedReports() = %v, want = %v", out, want)
	}
}

func TestRedNosedReportsP2(t *testing.T) {
	input := helpers.ReadFile("./input.txt")
	out := RedNosedReportsP2(input)
	want := 4

	if out != want {
		t.Errorf("RedNosedReportsP2() = %v, want = %v", out, want)
	}
}
