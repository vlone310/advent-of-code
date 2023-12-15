package year2023

import (
	"advent-of-code/helpers"
	"testing"
)

func TestPointofIncidence(t *testing.T) {
	input := helpers.ReadFile("./day13/input.txt")

	out := PointofIncidence(input)
	var want int = 405

	if out != want {
		t.Errorf("PointofIncidence() = %v, want %v", out, want)
	}
}

func TestPointofIncidenceP2(t *testing.T) {
	input := helpers.ReadFile("./day13/input.txt")

	out := PointofIncidenceP2(input)
	var want int = 400

	if out != want {
		t.Errorf("PointofIncidenceP2() = %v, want %v", out, want)
	}
}
