package main

import (
	"advent-of-code/helpers"
	"advent-of-code/year2022"
	"fmt"
)

func main() {
	input := helpers.ReadFile("./year2022/day2/input.txt")
	out := year2022.RockPaperScissorsP2(input)
	fmt.Printf("result: %v\n", out)
}
