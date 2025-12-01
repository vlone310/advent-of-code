package year2025

import (
	"testing"

	"github.com/vlone310/advent-of-code/helpers"
)

func TestSecretEntrance(t *testing.T) {
	input := helpers.ReadFile("./input.txt")

	out := SecretEntrance(input)
	want := 3

	if out != want {
		t.Errorf("SecretEntrance() = out %v, want = %v", out, want)
	}
}

func TestSecretEntranceP2(t *testing.T) {
	input := helpers.ReadFile("./input.txt")

	out := SecretEntranceP2(input)
	want := 6

	if out != want {
		t.Errorf("SecretEntranceP2() = out %v, want = %v", out, want)
	}
}
