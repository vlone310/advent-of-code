package year2023

import (
	"testing"

	"github.com/vlone310/advent-of-code/helpers"
)

func TestPointofIncidence(t *testing.T) {
	input := helpers.ReadFile("./input.txt")

	out := PointofIncidence(input)
	var want int = 405

	if out != want {
		t.Errorf("PointofIncidence() = %v, want %v", out, want)
	}
}

func TestPointofIncidenceP2(t *testing.T) {
	input := helpers.ReadFile("./input.txt")

	out := PointofIncidenceP2(input)
	var want int = 400

	if out != want {
		t.Errorf("PointofIncidenceP2() = %v, want %v", out, want)
	}
}
