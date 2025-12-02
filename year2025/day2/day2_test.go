package year2025

import (
	"testing"

	"github.com/vlone310/advent-of-code/helpers"
)

func TestGiftShop(t *testing.T) {
	input := helpers.ReadFile("./input.txt")

	out := GiftShop(input)
	want := 1227775554

	if out != want {
		t.Errorf("GiftShop() = out %v, want = %v", out, want)
	}
}

func TestGiftShopP2(t *testing.T) {
	input := helpers.ReadFile("./input.txt")

	out := GiftShopP2(input)
	want := 4174379265

	if out != want {
		t.Errorf("GiftShopP2() = out %v, want = %v", out, want)
	}
}
