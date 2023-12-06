package year2023

import (
	"advent-of-code/helpers"
	"testing"
)

func TestIfYouGiveASeedAFertilizer(t *testing.T) {
	input := helpers.ReadFile("./day5/input.txt")

	out := IfYouGiveASeedAFertilizer(input)
	var want int = 35

	if out != want {
		t.Errorf("IfYouGiveASeedAFertilizer() = %v, want %v", out, want)
	}
}

func TestIfYouGiveASeedAFertilizerP2(t *testing.T) {
	input := helpers.ReadFile("./day5/input.txt")

	out := IfYouGiveASeedAFertilizerP2(input)
	var want int = 46

	if out != want {
		t.Errorf("IfYouGiveASeedAFertilizerP2() = %v, want %v", out, want)
	}
}
