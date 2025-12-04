package year2025

import (
	"strings"

	"github.com/vlone310/advent-of-code/helpers"
)

var adjacent = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
	{-1, -1},
	{1, 1},
	{1, -1},
	{-1, 1},
}

func PrintingDepartment(input string) (res int) {
	grid := prepareInput(input)

	for x, row := range grid {

		for y, ch := range row {
			if ch == '@' && accessibleRoll(grid, x, y) {
				res++
			}
		}
	}
	return res
}

func PrintingDepartmentP2(input string) (res int) {
	grid := prepareInput(input)

	for {
		modified := false
		for x, row := range grid {
			for y, ch := range row {
				if ch == '@' && accessibleRoll(grid, x, y) {
					res++
					grid[x][y] = '.'
					modified = true
				}
			}
		}
		if !modified {
			break
		}
	}

	return res
}

func prepareInput(input string) [][]rune {

	rows := strings.Split(input, "\n")
	out := [][]rune{}

	for _, row := range rows {
		if row == "" {
			continue
		}
		out = append(out, []rune(row))
	}

	return out
}

func accessibleRoll(input [][]rune, x, y int) bool {
	adjacentRollsCount := 0

	for _, pos := range adjacent {
		newX := x + pos[0]
		newY := y + pos[1]

		if !helpers.IsOutOfBound(newX, len(input)) && !helpers.IsOutOfBound(newY, len(input[newX])) && input[newX][newY] == '@' {
			adjacentRollsCount++
		}
	}

	return adjacentRollsCount < 4
}
