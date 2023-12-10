package year2023

import (
	"advent-of-code/helpers"
	"testing"
)

func TestMirageMaintenance(t *testing.T) {
	input := helpers.ReadFile("./day9/input.txt")

	out := MirageMaintenance(input)
	var want int = 114

	if out != want {
		t.Errorf("MirageMaintenance() = %v, want %v", out, want)
	}
}

func TestMirageMaintenanceP2(t *testing.T) {
	input := helpers.ReadFile("./day9/input.txt")

	out := MirageMaintenanceP2(input)
	var want int = 2

	if out != want {
		t.Errorf("MirageMaintenanceP2() = %v, want %v", out, want)
	}
}
