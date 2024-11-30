package year2023

import (
	"testing"

	"github.com/vlone310/advent-of-code/helpers"
)

func TestMirageMaintenance(t *testing.T) {
	input := helpers.ReadFile("./input.txt")

	out := MirageMaintenance(input)
	var want int = 114

	if out != want {
		t.Errorf("MirageMaintenance() = %v, want %v", out, want)
	}
}

func TestMirageMaintenanceP2(t *testing.T) {
	input := helpers.ReadFile("./input.txt")

	out := MirageMaintenanceP2(input)
	var want int = 2

	if out != want {
		t.Errorf("MirageMaintenanceP2() = %v, want %v", out, want)
	}
}
