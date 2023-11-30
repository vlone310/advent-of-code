package year2022

import (
	"advent-of-code/helpers"
	"testing"
)

func TestCalorieCounting(t *testing.T){
	input := helpers.ReadFile("./day1/input.txt")

	out := CalorieCounting(input)
	var want float64 = 75622

	if out != want {
		t.Errorf("CalorieCounting() = %v, want %v", out, want)
	}
}

func TestCalorieCountingP2(t *testing.T){
	input := helpers.ReadFile("./day1/input.txt")

	out := CalorieCountingP2(input)
	var want float64 = 213159

	if out != want {
		t.Errorf("CalorieCountingP2() = %v, want %v", out, want)
	}
}