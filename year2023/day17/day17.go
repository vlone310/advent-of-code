package year2023

import (
	"strconv"
	"strings"

	"github.com/vlone310/advent-of-code/helpers"
)

const (
	R = "RIGHT"
	L = "LEFT"
	U = "UP"
	D = "DOWN"
)

type Coord struct {
	x, y int
}

type LowestHeatByCoordinates map[Coord]int

type QueueCoords struct {
	x, y, incX, incY, curHeat int
}

func ClumsyCrucible(input string) int {
	matrix := strings.Split(input, "\n")

	return searchPath(matrix)
}

func searchPath(matrix []string) int {
	queue := helpers.Queue[QueueCoords]{{0, 0, 0, 1, 0}, {0, 0, 1, 0, 0}}
	lowestHeatByCoordinatesMap := make(LowestHeatByCoordinates)

	for len(queue) > 0 {
		coords := queue.Pop()

		if lowestHeat, ok := lowestHeatByCoordinatesMap[Coord{coords.x, coords.y}]; ok && lowestHeat < coords.curHeat { // ???
			continue
		}

		if coords.x != 0 || coords.y != 0 {
			lowestHeatByCoordinatesMap[Coord{coords.x, coords.y}] = coords.curHeat
		}

		leftIncX, leftIncY := rotateLeft(coords.incX, coords.incY)
		rightIncX, rightIncY := rotateRight(coords.incX, coords.incY)
		newCurHeat := coords.curHeat

		for i := 1; i <= 3; i++ {
			newX := coords.x + (coords.incX * i)
			newY := coords.y + (coords.incY * i)

			if !helpers.IsOutOfBound(newX, len(matrix)) && !helpers.IsOutOfBound(newY, len(matrix[0])) {
				val, _ := strconv.Atoi(string(matrix[newX][newY]))
				newCurHeat += val
				cachedHeat, ok := lowestHeatByCoordinatesMap[Coord{newX, newY}]
				if (ok && cachedHeat > newCurHeat) || !ok {
					queue.Push(QueueCoords{
						newX,
						newY,
						leftIncX,
						leftIncY,
						newCurHeat,
					})

					queue.Push(QueueCoords{
						newX,
						newY,
						rightIncX,
						rightIncY,
						newCurHeat,
					})
				}
			}
		}
	}

	val := lowestHeatByCoordinatesMap[Coord{len(matrix) - 1, len(matrix[0]) - 1}]

	return val
}
func rotateLeft(x, y int) (int, int) {
	if x == -1 {
		return 0, -1
	} else if y == -1 {
		return 1, 0
	} else if y == 1 {
		return -1, 0
	} else {
		return 0, 1
	}
}

func rotateRight(x, y int) (int, int) {
	if x == -1 {
		return 0, 1
	} else if y == -1 {
		return -1, 0
	} else if y == 1 {
		return 1, 0
	} else {
		return 0, -1
	}
}
