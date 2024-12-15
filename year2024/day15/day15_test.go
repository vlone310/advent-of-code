package year2024

import (
	"testing"

	"github.com/vlone310/advent-of-code/helpers"
)

func TestWarehouseWoes(t *testing.T) {
	input := helpers.ReadFile("./input.txt")
	out := warehouseWoes(input)
	want := 10092

	if out != want {
		t.Errorf("warehouseWoes() = %v, want = %v", out, want)
	}
}

func TestWarehouseWoesP2(t *testing.T) {
	input := helpers.ReadFile("./input.txt")
	out := warehouseWoesP2(input)
	want := 9021

	if out != want {
		t.Errorf("warehouseWoesP2() = %v, want = %v", out, want)
	}
}
