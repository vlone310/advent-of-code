package year2023

import (
	"advent-of-code/helpers"
	"math"
	"slices"
	"strings"
)

type XY struct {
	x, y int
}

type EnergizedMap map[XY][]string

const (
	RIGHT = "R"
	LEFT  = "L"
	UP    = "U"
	DOWN  = "D"
)

var nearestIndiciesMap = map[string][2]int{
	UP:    {0, -1},
	DOWN:  {0, 1},
	LEFT:  {-1, 0},
	RIGHT: {1, 0},
}

var mirrorsToDirectionMap = map[string]map[string]string{
	"/": {
		UP:    RIGHT,
		DOWN:  LEFT,
		LEFT:  DOWN,
		RIGHT: UP,
	},
	`\`: {
		UP:    LEFT,
		DOWN:  RIGHT,
		LEFT:  UP,
		RIGHT: DOWN,
	},
}

func TheFloorWillBeLava(input string) int {
	energized := make(EnergizedMap)
	matrix := strings.Split(input, "\n")

	findDirectionAndMarkEnergized(&energized, matrix, 0, 0, RIGHT)

	return len(energized)
}

func TheFloorWillBeLavaP2(input string) (max int) {
	matrix := strings.Split(input, "\n")

	// TOP ROW
	for i := 0; i < len(matrix[0]); i++ {
		energized := make(EnergizedMap)
		findDirectionAndMarkEnergized(&energized, matrix, i, 0, DOWN)
		max = int(math.Max(float64(max), float64(len(energized))))
	}

	// BOTTOM ROW
	for i := 0; i < len(matrix[0]); i++ {
		energized := make(EnergizedMap)
		findDirectionAndMarkEnergized(&energized, matrix, i, len(matrix)-1, UP)
		max = int(math.Max(float64(max), float64(len(energized))))
	}

	// LEFT ROW
	for i := 0; i < len(matrix); i++ {
		energized := make(EnergizedMap)
		findDirectionAndMarkEnergized(&energized, matrix, 0, i, RIGHT)
		max = int(math.Max(float64(max), float64(len(energized))))
	}

	// RIGHT ROW
	for i := 0; i < len(matrix); i++ {
		energized := make(EnergizedMap)
		findDirectionAndMarkEnergized(&energized, matrix, len(matrix[0])-1, i, LEFT)
		max = int(math.Max(float64(max), float64(len(energized))))
	}

	return
}

func findDirectionAndMarkEnergized(energized *EnergizedMap, matrix []string, Y int, X int, dir string) int {
	if helpers.IsOutOfBound(X, len(matrix)) || helpers.IsOutOfBound(Y, len(matrix[0])) {
		return 0
	}

	if val, ok := (*energized)[XY{X, Y}]; ok && slices.Contains(val, dir) {
		return 0
	}

	(*energized)[XY{X, Y}] = append((*energized)[XY{X, Y}], dir)

	if matrix[X][Y] == '.' || (matrix[X][Y] == '|' && (dir == UP || dir == DOWN)) || (matrix[X][Y] == '-' && (dir == LEFT || dir == RIGHT)) {
		findDirectionAndMarkEnergized(energized, matrix, Y+nearestIndiciesMap[dir][0], X+nearestIndiciesMap[dir][1], dir)
	} else if matrix[X][Y] == '|' {
		findDirectionAndMarkEnergized(energized, matrix, Y+nearestIndiciesMap[UP][0], X+nearestIndiciesMap[UP][1], UP)
		findDirectionAndMarkEnergized(energized, matrix, Y+nearestIndiciesMap[DOWN][0], X+nearestIndiciesMap[DOWN][1], DOWN)
	} else if matrix[X][Y] == '-' {
		findDirectionAndMarkEnergized(energized, matrix, Y+nearestIndiciesMap[LEFT][0], X+nearestIndiciesMap[LEFT][1], LEFT)
		findDirectionAndMarkEnergized(energized, matrix, Y+nearestIndiciesMap[RIGHT][0], X+nearestIndiciesMap[RIGHT][1], RIGHT)
	} else {
		newDir := mirrorsToDirectionMap[string(matrix[X][Y])][dir]
		findDirectionAndMarkEnergized(energized, matrix, Y+nearestIndiciesMap[newDir][0], X+nearestIndiciesMap[newDir][1], newDir)
	}

	return 0
}
