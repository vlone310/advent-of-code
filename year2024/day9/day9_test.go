package year2024

import (
	"testing"

	"github.com/vlone310/advent-of-code/helpers"
)

func TestDiskFragmenter(t *testing.T) {
	input := helpers.ReadFile("./input.txt")
	out := diskFragmenter(input)
	want := 1928

	if out != want {
		t.Errorf("diskFragmenterP2() = %v, want = %v", out, want)
	}
}

func TestDiskFragmenterP2(t *testing.T) {
	input := helpers.ReadFile("./input.txt")
	out := diskFragmenterP2(input)
	want := 2858

	if out != want {
		t.Errorf("diskFragmenterP2() = %v, want = %v", out, want)
	}
}
