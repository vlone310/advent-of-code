package year2025

import (
	"testing"

	"github.com/vlone310/advent-of-code/helpers"
)

func TestLobby(t *testing.T) {
	input := helpers.ReadFile("./input.txt")

	out := Lobby(input)
	want := 357

	if out != want {
		t.Errorf("Lobby() = out %v, want = %v", out, want)
	}
}

func TestLobbyP2(t *testing.T) {
	input := helpers.ReadFile("./input.txt")

	out := LobbyP2(input)
	want := 3121910778619

	if out != want {
		t.Errorf("LobbyP2() = out %v, want = %v", out, want)
	}
}
