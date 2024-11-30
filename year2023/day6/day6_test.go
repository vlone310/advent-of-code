package year2023

import (
	"testing"

	"github.com/vlone310/advent-of-code/helpers"
)

func TestWaitForIt(t *testing.T) {
	input := helpers.ReadFile("./input.txt")

	out := WaitForIt(input)
	var want int = 288

	if out != want {
		t.Errorf("WaitForIt() = %v, want %v", out, want)
	}
}

func TestWaitForItP2(t *testing.T) {
	input := helpers.ReadFile("./input.txt")

	out := WaitForItP2(input)
	var want int = 71503

	if out != want {
		t.Errorf("WaitForItP2() = %v, want %v", out, want)
	}
}
