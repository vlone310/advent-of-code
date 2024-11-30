package year2023

import (
	"math"
	"strconv"
	"strings"

	"github.com/vlone310/advent-of-code/helpers"
)

var dirToNewIndicies = map[string][]int{
	"R": {0, 1},
	"L": {0, -1},
	"U": {-1, 0},
	"D": {1, 0},
}

type Coordinates struct {
	y, x int
}

var (
	dx = []int{1, -1, 0, 0}
	dy = []int{0, 0, 1, -1}
)

func LavaductLagoon(input string) int {
	trench := make(map[Coordinates]bool, 0)
	cubicMetersDigged := 0
	lines := strings.Split(input, "\n")
	y, x := 0, 0
	maxY, maxX := -math.MaxInt, -math.MaxInt
	minY, minX := math.MaxInt, math.MaxInt

	for _, line := range lines {
		lineSplit := strings.Split(line, " ")
		dir := lineSplit[0]
		steps, _ := strconv.Atoi(lineSplit[1])
		newIndicies := dirToNewIndicies[dir]
		newY, newX := newIndicies[0], newIndicies[1]

		for i := 0; i < steps; i++ {
			x += newX
			y += newY
			cubicMetersDigged++
			trench[Coordinates{y, x}] = true
			maxY = helpers.Max(maxY, y)
			maxX = helpers.Max(maxX, x)
			minY = helpers.Min(minY, y)
			minX = helpers.Min(minX, x)
		}
	}

	trenchesInside := countTrenches(trench, minY, maxY, minX, maxX)
	return cubicMetersDigged + trenchesInside
}

func LavaductLagoonP2(input string) int {
	trench := make(map[Coordinates]bool, 0)
	cubicMetersDigged := 0
	lines := strings.Split(input, "\n")
	y, x := 0, 0
	maxY, maxX := -math.MaxInt, -math.MaxInt
	minY, minX := math.MaxInt, math.MaxInt

	for _, line := range lines {
		lineSplit := strings.Split(line, " ")
		hexStr := lineSplit[2]
		newIndicies := dirToNewIndicies[byteToDir(hexStr[len(hexStr)-2])]
		steps, _ := strconv.ParseInt(hexStr[2:len(hexStr)-2], 16, 64)
		newY, newX := newIndicies[0], newIndicies[1]

		for i := 0; i < int(steps); i++ {
			x += newX
			y += newY
			cubicMetersDigged++
			trench[Coordinates{y, x}] = true
			maxY = helpers.Max(maxY, y)
			maxX = helpers.Max(maxX, x)
			minY = helpers.Min(minY, y)
			minX = helpers.Min(minX, x)
		}
	}

	trenchesInside := countTrenches(trench, minY, maxY, minX, maxX)
	return cubicMetersDigged + trenchesInside
}

func countTrenches(trench map[Coordinates]bool, minY, maxY, minX, maxX int) int {
	visited := make(map[Coordinates]bool, 0)
	count := 0
	countOutOfBoundaryDots(minY, maxY, minX, maxX, visited, trench)

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			_, isTrench := trench[Coordinates{y, x}]
			_, isVisited := visited[Coordinates{y, x}]
			if !isVisited && !isTrench {
				count += 1
			}
		}
	}
	return count
}

func countOutOfBoundaryDots(minY, maxY, minX, maxX int, visited, trench map[Coordinates]bool) {
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			_, isVisited := visited[Coordinates{y, x}]
			_, isTrench := trench[Coordinates{y, x}]
			if (y == minY || y == maxY || x == minX || x == maxX) && !isVisited && !isTrench {
				floodFill(y, x, minY, maxY, minX, maxX, visited, trench)
			}
		}
	}
}

func floodFill(y, x, minY, maxY, minX, maxX int, visited, trench map[Coordinates]bool) {
	visited[Coordinates{y, x}] = true
	for i := 0; i < 4; i++ {
		newY := y + dy[i]
		newX := x + dx[i]

		if isValid(newY, newX, minY, maxY, minX, maxX, visited, trench) {
			floodFill(newY, newX, minY, maxY, minX, maxX, visited, trench)
		}
	}
}

func isValid(y, x, minY, maxY, minX, maxX int, visited, trench map[Coordinates]bool) bool {
	_, isPipe := trench[Coordinates{y, x}]
	_, isVisited := visited[Coordinates{y, x}]
	return x >= minX && x <= maxX && y >= minY && y < maxY && !isPipe && !isVisited
}

func byteToDir(r byte) string {
	switch r {
	case '0':
		return "R"
	case '1':
		return "D"
	case '2':
		return "L"
	case '3':
		return "U"
	default:
		return ""
	}
}
