package year2024

import (
	"strconv"
	"strings"

	"github.com/vlone310/advent-of-code/helpers"
)

func plutonianPebbles(input string, blinks int) int {
	defer helpers.Timer()()
	stones := parseInput(input)
	counts := make(map[int]int)

	for _, stone := range stones {
		counts[stone]++
	}

	for blink := 0; blink < blinks; blink++ {
		newCounts := make(map[int]int)

		for stone, count := range counts {
			switch {
			case stone == 0:
				newCounts[1] += count
			case len(strconv.Itoa(stone))%2 == 0:
				split := splitStone(stone)
				for _, s := range split {
					newCounts[s] += count
				}
			default:
				newStone := stone * 2024
				newCounts[newStone] += count
			}
		}

		counts = newCounts
	}

	totalStones := 0
	for _, count := range counts {
		totalStones += count
	}

	return totalStones
}

func splitStone(num int) []int {
	str := strconv.Itoa(num)

	middle := len(str) / 2
	part1, _ := strconv.Atoi(str[:middle])
	part2, _ := strconv.Atoi(str[middle:])

	return []int{part1, part2}

}

func parseInput(input string) (res []int) {
	for _, el := range strings.Split(strings.ReplaceAll(input, "\n", ""), " ") {
		num, _ := strconv.Atoi(el)
		res = append(res, num)
	}

	return
}
