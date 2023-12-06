package main

import (
	"advent-of-code/helpers"
	"advent-of-code/year2023"
	"fmt"
)

func main() {
	input := helpers.ReadFile("./year2023/day6/input.txt")
	out := year2023.WaitForItP2(input)
	fmt.Printf("result: %v\n", out)
}
