package year2024

import (
	"strings"

	"github.com/vlone310/advent-of-code/helpers"
)

type coords struct {
	x, y int
}

var guard = '^'
var up = coords{-1, 0}
var down = coords{1, 0}
var right = coords{0, 1}
var left = coords{0, -1}

func guardGallivant(input string) int {
	defer helpers.Timer()()
	arr := strings.Split(input, "\n")
	x, y := findGuard(arr)
	curDir := up
	visited := map[coords]bool{{x, y}: true}

	for {
		nextX := x + curDir.x
		nextY := y + curDir.y
		if helpers.IsOutOfBound(nextX, len(arr)) || helpers.IsOutOfBound(nextY, len(arr[nextX])) {
			break
		}

		// rotate 90deg
		if arr[nextX][nextY] == '#' {
			curDir = getRotation(curDir)
			continue
		}

		visited[coords{nextX, nextY}] = true
		x = nextX
		y = nextY
	}

	return len(visited)
}

func guardGallivantP2(input string) int {
	defer helpers.Timer()()

	arr := strings.Split(input, "\n")
	x, y := findGuard(arr)
	possibleCrates := map[coords]bool{}
	checkedCrates := map[coords]bool{}
	curDir := up

	for {
		nextX := x + curDir.x
		nextY := y + curDir.y
		if helpers.IsOutOfBound(nextX, len(arr)) || helpers.IsOutOfBound(nextY, len(arr[nextX])) {
			break
		}

		// rotate 90deg
		if arr[nextX][nextY] == '#' {
			curDir = getRotation(curDir)
			continue
		}

		// check if possible to place crate and detect cycle
		if _, exists := checkedCrates[coords{nextX, nextY}]; !exists && isValidCycle(x, y, coords{nextX, nextY}, curDir, arr) {
			possibleCrates[coords{nextX, nextY}] = true
		}

		checkedCrates[coords{nextX, nextY}] = true

		x = nextX
		y = nextY
	}

	return len(possibleCrates)
}

func isValidCycle(startX, startY int, possibleCrate, dir coords, arr []string) bool {
	x := startX
	y := startY
	curDir := dir
	visited := map[coords][]coords{}

	for {
		nextX := x + curDir.x
		nextY := y + curDir.y

		if helpers.IsOutOfBound(nextX, len(arr)) || helpers.IsOutOfBound(nextY, len(arr[nextX])) {
			return false
		}

		visitedDirs, exists := visited[coords{x, y}]
		if exists && isVisitedWithDir(curDir, visitedDirs) {
			return true
		}
		if exists {
			visited[coords{x, y}] = append(visited[coords{x, y}], curDir)
		} else {
			visited[coords{x, y}] = []coords{curDir}
		}

		if arr[nextX][nextY] == '#' || (nextX == possibleCrate.x && nextY == possibleCrate.y) {
			curDir = getRotation(curDir)
			continue
		}

		x += curDir.x
		y += curDir.y
	}
}

func isVisitedWithDir(dir coords, visitedDirs []coords) bool {
	for _, coord := range visitedDirs {
		if coord == dir {
			return true
		}
	}
	return false
}

func getRotation(c coords) coords {
	switch c {
	case up:
		return right
	case right:
		return down
	case down:
		return left
	default:
		return up
	}
}

func findGuard(arr []string) (x, y int) {
	for x, row := range arr {
		for y, el := range row {
			if el == guard {
				return x, y
			}
		}
	}

	return -1, -1
}

func copyVisited(original map[coords][]coords) map[coords][]coords {
	// Create a new map
	copyMap := make(map[coords][]coords)

	// Iterate over the original map
	for key, valueSlice := range original {
		// Create a new slice for the copied map
		newSlice := make([]coords, len(valueSlice))

		// Copy each element of the slice
		copy(newSlice, valueSlice)

		// Set the copied slice in the new map
		copyMap[key] = newSlice
	}

	return copyMap
}
