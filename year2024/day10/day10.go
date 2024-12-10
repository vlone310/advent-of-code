package year2024

import (
	"strings"

	"github.com/vlone310/advent-of-code/helpers"
)

var directions = [4][2]int{
	{0, -1},
	{0, 1},
	{-1, 0},
	{1, 0},
}

type coords struct {
	x, y int
}

func hoofIt(input string) (res int) {
	grid := strings.Split(input, "\n")

	for x, row := range grid {
		for y, el := range row {
			if byte(el) == '0' {
				res += searchTops(x, y, grid, map[coords]bool{})
			}
		}
	}

	return
}

func hoofItP2(input string) (res int) {
	grid := strings.Split(input, "\n")

	for x, row := range grid {
		for y, el := range row {
			if byte(el) == '0' {
				res += searchTopsP2(x, y, grid, map[coords]int{})
			}
		}
	}

	return
}

func searchTopsP2(x, y int, grid []string, tops map[coords]int) int {
	if byte(grid[x][y]) == '9' {
		if _, exists := tops[coords{x, y}]; exists {
			tops[coords{x, y}] += 1
		} else {
			tops[coords{x, y}] = 1
		}
	}

	res := 0

	for _, dir := range directions {
		newX := x + dir[0]
		newY := y + dir[1]
		if !helpers.IsOutOfBound(newX, len(grid)) &&
			!helpers.IsOutOfBound(newY, len(grid[newX])) &&
			int(byte(grid[x][y])) == int(byte(grid[newX][newY]))-1 {
			searchTopsP2(newX, newY, grid, tops)
		}
	}

	for _, count := range tops {
		res += count
	}

	return res
}

func searchTops(x, y int, grid []string, tops map[coords]bool) int {
	if byte(grid[x][y]) == '9' {
		if _, exists := tops[coords{x, y}]; exists {
			return 0
		} else {
			tops[coords{x, y}] = true
			return 1
		}
	}

	res := 0

	for _, dir := range directions {
		newX := x + dir[0]
		newY := y + dir[1]
		if !helpers.IsOutOfBound(newX, len(grid)) &&
			!helpers.IsOutOfBound(newY, len(grid[newX])) &&
			int(byte(grid[x][y])) == int(byte(grid[newX][newY]))-1 {
			res += searchTops(newX, newY, grid, tops)
		}
	}

	return res
}
