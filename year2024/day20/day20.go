package year2024

import (
	// "fmt"
	"math"
	"strings"
)

type coords struct {
	x, y int
}

type track struct {
	grid       [][]byte
	start, end coords
	path       []coords
}

var directions = []coords{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

type opts struct {
	cheatTreshold, cheatPenalty, maxCheat int
}

func raceCondition(input string, options opts) int {
	t := parseTrack(input)

	return t.cheatCount(options.cheatTreshold, options.maxCheat)

}

func (t track) cheatCount(cheatTreshold, maxCheat int) int {
	res := 0

	for i, c1 := range t.path[:len(t.path)-cheatTreshold] {
		for j, c2 := range t.path[i+cheatTreshold:] {
			if dist := heuristic(c2, c1); dist <= maxCheat && dist <= j {
				res += 1
			}
		}
	}

	return res
}

func heuristic(a, b coords) int {
	return int(math.Abs(float64(a.x-b.x)) + math.Abs(float64(a.y-b.y)))
}

func parseTrack(input string) (t track) {
	split := strings.Split(input, "\n")
	t.grid = make([][]byte, 0, len(split))

	for x, row := range split {
		for y, el := range row {
			if byte(el) == 'S' {
				t.start = coords{x, y}
			}
			if byte(el) == 'E' {
				t.end = coords{x, y}
			}
		}
		t.grid = append(t.grid, []byte(row))
	}

	findFullPath(&t)

	return
}

func findFullPath(t *track) {
	visited := make(map[coords]bool)
	cur := t.start
	path := []coords{t.start}

	for cur != t.end {
		for _, dir := range directions {
			next := coords{cur.x + dir.x, cur.y + dir.y}
			if visited[next] {
				continue
			}

			nextEl := t.grid[next.x][next.y]

			if nextEl == '.' || nextEl == 'E' {
				visited[next] = true
				path = append(path, next)
				cur = next
				break
			}
		}
	}

	t.path = path

}
