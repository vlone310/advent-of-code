package main

import (
	"advent-of-code/helpers"
	"advent-of-code/year2023"
	"fmt"
)

func main() {
	input := helpers.ReadFile("./year2023/day8/input2.txt")
	out := year2023.HauntedWastelandP2(input)
	fmt.Printf("result: %v\n", out)
}
