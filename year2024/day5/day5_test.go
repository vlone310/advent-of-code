package year2024

import (
	"testing"

	"github.com/vlone310/advent-of-code/helpers"
)

func TestPringQueue(t *testing.T) {
	input := helpers.ReadFile("./input.txt")
	out := printQueue(input)
	want := 143

	if out != want {
		t.Errorf("printQueue() = %v, want = %v", out, want)
	}
}

func TestPringQueueP2(t *testing.T) {
	input := helpers.ReadFile("./input.txt")
	out := printQueueP2(input)
	want := 123

	if out != want {
		t.Errorf("printQueueP2() = %v, want = %v", out, want)
	}
}
