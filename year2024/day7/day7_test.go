package year2024

import (
	"testing"

	"github.com/vlone310/advent-of-code/helpers"
)

func TestBridgeRepair(t *testing.T) {
	input := helpers.ReadFile("./input.txt")
	out := bridgeRepair(input)
	want := 3749

	if out != want {
		t.Errorf("bridgeRepair() = %v, want = %v", out, want)
	}
}

func TestBridgeRepairP2(t *testing.T) {
	input := helpers.ReadFile("./input.txt")
	out := bridgeRepairP2(input)
	want := 11387

	if out != want {
		t.Errorf("bridgeRepairP2() = %v, want = %v", out, want)
	}
}
