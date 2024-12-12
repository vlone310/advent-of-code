package year2024

import (
	"strings"

	"github.com/vlone310/advent-of-code/helpers"
)

type coords struct {
	x, y int
}

var directions = [4][2]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func gardenGroups(input string) (res int) {
	garden := strings.Split(input, "\n")
	visited := map[coords]bool{}

	for x, row := range garden {
		for y, plant := range row {
			prevVisits := len(visited)
			per := calculatePlantPerimeterAndMarkVisited(x, y, plant, garden, visited)
			res += (len(visited) - prevVisits) * per
		}
	}

	return
}

func gardenGroupsP2(input string) (res int) {
	garden := strings.Split(input, "\n")
	visited := map[coords]bool{}

	for x, row := range garden {
		for y, plant := range row {
			prevVisits := len(visited)
			sides := countSidesAndMarkVisited(x, y, plant, garden, visited)
			res += (len(visited) - prevVisits) * sides
		}
	}

	return
}

func countSidesAndMarkVisited(x, y int, plant rune, garden []string, visited map[coords]bool) int {
	if _, exists := visited[coords{x, y}]; exists {
		return 0
	}
	visited[coords{x, y}] = true
	cornersCount := getCornerCount(x, y, plant, garden)

	for _, dir := range directions {
		nextX := x + dir[0]
		nextY := y + dir[1]

		if helpers.IsOutOfBound(nextX, len(garden)) ||
			helpers.IsOutOfBound(nextY, len(garden[nextX])) ||
			garden[nextX][nextY] != byte(plant) {
			continue
		}

		cornersCount += countSidesAndMarkVisited(nextX, nextY, plant, garden, visited)
	}

	return cornersCount
}

func getCornerCount(x, y int, plant rune, garden []string) int {
	type square struct {
		neighbours []coords
		diagonalEl coords
	}
	possibleSquares := []square{
		{neighbours: []coords{{x - 1, y}, {x, y + 1}}, diagonalEl: coords{x - 1, y + 1}},
		{neighbours: []coords{{x, y + 1}, {x + 1, y}}, diagonalEl: coords{x + 1, y + 1}},
		{neighbours: []coords{{x + 1, y}, {x, y - 1}}, diagonalEl: coords{x + 1, y - 1}},
		{neighbours: []coords{{x, y - 1}, {x - 1, y}}, diagonalEl: coords{x - 1, y - 1}},
	}

	cornerCount := 0

	for _, possibleSquare := range possibleSquares {
		n1 := possibleSquare.neighbours[0]
		n2 := possibleSquare.neighbours[1]
		diag := possibleSquare.diagonalEl
		if (helpers.IsOutOfBound(n1.x, len(garden)) || helpers.IsOutOfBound(n1.y, len(garden[n1.x])) || garden[n1.x][n1.y] != byte(plant)) &&
			(helpers.IsOutOfBound(n2.x, len(garden)) || helpers.IsOutOfBound(n2.y, len(garden[n2.x])) || garden[n2.x][n2.y] != byte(plant)) {
			cornerCount++
		}

		if !helpers.IsOutOfBound(n1.x, len(garden)) && !helpers.IsOutOfBound(n1.y, len(garden[n1.x])) && garden[n1.x][n1.y] == byte(plant) &&
			!helpers.IsOutOfBound(n2.x, len(garden)) && !helpers.IsOutOfBound(n2.y, len(garden[n2.x])) && garden[n2.x][n2.y] == byte(plant) &&
			(helpers.IsOutOfBound(diag.x, len(garden)) || helpers.IsOutOfBound(diag.y, len(garden[diag.x])) || garden[diag.x][diag.y] != byte(plant)) {
			cornerCount++
		}
	}
	return cornerCount
}

func calculatePlantPerimeterAndMarkVisited(x, y int, plant rune, garden []string, visited map[coords]bool) int {
	if _, exists := visited[coords{x, y}]; exists {
		return 0
	}

	visited[coords{x, y}] = true

	curPer := 4

	for _, dir := range directions {
		nextX := x + dir[0]
		nextY := y + dir[1]

		if helpers.IsOutOfBound(nextX, len(garden)) ||
			helpers.IsOutOfBound(nextY, len(garden[nextX])) ||
			garden[nextX][nextY] != byte(plant) {
			continue
		}

		curPer--
		curPer += calculatePlantPerimeterAndMarkVisited(nextX, nextY, plant, garden, visited)
	}

	return curPer
}
