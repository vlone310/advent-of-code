package year2023

import (
	"advent-of-code/helpers"
	"errors"
	"math"
	"slices"
	"strings"
)

var connectorsMap = map[string][]string{
	"J": {"L", "T"},
	"F": {"B", "R"},
	"L": {"T", "R"},
	"7": {"L", "B"},
	"|": {"T", "B"},
	"-": {"L", "R"},
	"S": {"T", "B", "L", "R"},
}

var nearestIndicies = map[string][2]int{
	"T": {0, -1},
	"B": {0, 1},
	"L": {-1, 0},
	"R": {1, 0},
}

var reversedDir = map[string]string{
	"T": "B",
	"B": "T",
	"L": "R",
	"R": "L",
}

func PipeMaze(input string) (distance int) {
	rows := strings.Split(input, "\n")

	for rI, row := range rows {
		for cI, el := range row {

			if el != 'S' {
				continue
			}

			for key, val := range nearestIndicies {
				length, err := findCycle(rI+val[1], cI+val[0], rows, reversedDir[key], 1)

				if err == nil {
					distance = int(math.Max(float64(distance), float64(length))) / 2
					break
				}
			}

		}
	}
	return
}

func PipeMazeP2(input string) (count int) {
	// rows :=

	return
}

func findCycle(rI int, cI int, rows []string, cameFrom string, currLength int) (int, error) {
	if helpers.IsOutOfBound(rI, len(rows)) || helpers.IsOutOfBound(cI, len(rows[0])) {
		return 0, errors.New("")
	}

	currPipe := rows[rI][cI]

	if currPipe == '.' {
		return 0, errors.New("")
	}

	if currPipe == 'S' {
		return currLength, nil
	}

	if !slices.Contains(connectorsMap[string(currPipe)], cameFrom) {
		return 0, errors.New("")
	}

	var nextDir string

	for _, dir := range connectorsMap[string(currPipe)] {
		if dir == cameFrom {
			continue
		}

		nextDir = dir
	}

	return findCycle(rI+nearestIndicies[nextDir][1], cI+nearestIndicies[nextDir][0], rows, reversedDir[nextDir], currLength+1)
}
