package year2023

import (
	"advent-of-code/helpers"
	"testing"
)

func TestPipeMaze(t *testing.T) {
	input := helpers.ReadFile("./day10/input.txt")

	out := PipeMaze(input)
	var want int = 8

	if out != want {
		t.Errorf("PipeMaze() = %v, want %v", out, want)
	}
}

func TestPipeMazeP2(t *testing.T) {
	input := helpers.ReadFile("./day10/input2.txt")

	out := PipeMazeP2(input)
	var want int = 0

	if out != want {
		t.Errorf("PipeMazeP2() = %v, want %v", out, want)
	}
}
