package year2024

import (
	"testing"

	"github.com/vlone310/advent-of-code/helpers"
)

func TestReindeerMaze(t *testing.T) {
	input := helpers.ReadFile("./input.txt")
	out := reindeerMaze(input)
	want := 11048

	if out != want {
		t.Errorf("reindeerMaze() = %v, want = %v", out, want)
	}
}

func TestReindeerMazeP2(t *testing.T) {
	input := helpers.ReadFile("./input.txt")
	out := reindeerMazeP2(input)
	want := 64

	if out != want {
		t.Errorf("reindeerMazeP2() = %v, want = %v", out, want)
	}
}
