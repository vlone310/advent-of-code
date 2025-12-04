package year2025

import (
	"testing"

	"github.com/vlone310/advent-of-code/helpers"
)

func TestPrintingDepartment(t *testing.T) {
	input := helpers.ReadFile("./input.txt")

	out := PrintingDepartment(input)
	want := 13

	if out != want {
		t.Errorf("PrintingDepartment() = out %v, want = %v", out, want)
	}
}

func TestPrintingDepartmentP2(t *testing.T) {
	input := helpers.ReadFile("./input.txt")

	out := PrintingDepartmentP2(input)
	want := 43

	if out != want {
		t.Errorf("PrintingDepartmentP2() = out %v, want = %v", out, want)
	}
}
