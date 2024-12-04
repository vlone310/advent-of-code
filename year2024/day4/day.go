package year2024

import (
	"strings"

	"github.com/vlone310/advent-of-code/helpers"
)

var nearestIndicies = [8][2]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
	{-1, -1},
	{-1, 1},
	{1, -1},
	{1, 1},
}

var diagonalIndicies = [2][2]int{
	{-1, -1},
	{-1, 1},
}

func CeresSearch(input string) (xmasCount int) {
	rows := strings.Split(input, "\n")

	for x, row := range rows {
		for y, char := range row {
			if char == 'X' {
				for _, indicies := range nearestIndicies {
					xmasCount += findXMAS(x, y, rows, indicies, 'X')
				}
			}
		}
	}

	return
}

func CeresSearchP2(input string) (xmasCount int) {
	rows := strings.Split(input, "\n")

	for x, row := range rows {
		for y, char := range row {
			if char == 'A' && isXMAS(x, y, rows) {
				xmasCount += 1
			}
		}
	}

	return
}

func findXMAS(x, y int, grid []string, indicies [2]int, cur byte) int {
	if cur == 'S' {
		return 1
	}

	nextX := x + indicies[0]
	nextY := y + indicies[1]

	if helpers.IsOutOfBound(nextX, len(grid)) || helpers.IsOutOfBound(nextY, len(grid[nextX])) {
		return 0
	}

	nextChar := grid[nextX][nextY]

	if (cur == 'X' && nextChar == 'M') || (cur == 'M' && nextChar == 'A') || (cur == 'A' && nextChar == 'S') {
		return findXMAS(nextX, nextY, grid, indicies, nextChar)
	}

	return 0
}

func isXMAS(x, y int, grid []string) bool {
	for _, indicies := range diagonalIndicies {
		nextX := x + indicies[0]
		nextY := y + indicies[1]
		diagX := x - indicies[0]
		diagY := y - indicies[1]

		if helpers.IsOutOfBound(nextX, len(grid)) ||
			helpers.IsOutOfBound(nextY, len(grid[nextX])) ||
			helpers.IsOutOfBound(diagX, len(grid)) ||
			helpers.IsOutOfBound(diagY, len(grid[diagX])) {
			return false
		}

		char := grid[nextX][nextY]
		diagChar := grid[diagX][diagY]

		if (char != 'M' && char != 'S') || (char == 'M' && diagChar != 'S') || (char == 'S' && diagChar != 'M') {
			return false
		}
	}

	return true
}
