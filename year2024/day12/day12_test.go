package year2024

import (
	"testing"

	"github.com/vlone310/advent-of-code/helpers"
)

func TestGardenGroups(t *testing.T) {
	input := helpers.ReadFile("./input.txt")
	out := gardenGroups(input)
	want := 1930

	if out != want {
		t.Errorf("gardenGroups() = %v, want = %v", out, want)
	}
}

func TestGardenGroupsP2(t *testing.T) {
	input := helpers.ReadFile("./input.txt")
	out := gardenGroupsP2(input)
	want := 1206

	if out != want {
		t.Errorf("gardenGroupsP2() = %v, want = %v", out, want)
	}
}
