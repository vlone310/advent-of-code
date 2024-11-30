package year2023

import (
	"testing"

	"github.com/vlone310/advent-of-code/helpers"
)

func TestIfYouGiveASeedAFertilizer(t *testing.T) {
	input := helpers.ReadFile("./input.txt")

	out := IfYouGiveASeedAFertilizer(input)
	var want int = 35

	if out != want {
		t.Errorf("IfYouGiveASeedAFertilizer() = %v, want %v", out, want)
	}
}

func TestIfYouGiveASeedAFertilizerP2(t *testing.T) {
	input := helpers.ReadFile("./input.txt")

	out := IfYouGiveASeedAFertilizerP2(input)
	var want int = 46

	if out != want {
		t.Errorf("IfYouGiveASeedAFertilizerP2() = %v, want %v", out, want)
	}
}
