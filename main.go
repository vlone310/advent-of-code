package main

import (
	"advent-of-code/helpers"
	"advent-of-code/year2023"
	"fmt"
)

func main() {
	input := helpers.ReadFile("./year2023/day4/input.txt")
	out := year2023.ScratchcardsP2(input)
	fmt.Printf("result: %v\n", out)
}
