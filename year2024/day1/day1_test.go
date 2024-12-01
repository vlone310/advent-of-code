package year2024

import (
	"testing"

	"github.com/vlone310/advent-of-code/helpers"
)

func TestHistorianHysteria(t *testing.T) {
	input := helpers.ReadFile("./input.txt")

	out := HistorianHysteria(input)
	want := 11

	if out != want {
		t.Errorf("HistorianHysteria() = out %v, want = %v", out, want)
	}
}

func TestHistorianHysteriaP2(t *testing.T) {
	input := helpers.ReadFile("./input.txt")

	out := HistorianHysteriaP2(input)
	want := 31

	if out != want {
		t.Errorf("HistorianHysteria() = out %v, want = %v", out, want)
	}
}
