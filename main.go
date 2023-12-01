package main

import (
	"advent-of-code/helpers"
	"advent-of-code/year2023"
	"fmt"
)

func main() {
	input := helpers.ReadFile("./year2023/day1/input.txt")
	out := year2023.TrebuchetP2(input)
	fmt.Printf("result: %v\n", out)
}
