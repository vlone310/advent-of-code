package year2024

import (
	"fmt"
	"math"
	"strings"

	"github.com/vlone310/advent-of-code/helpers"
	"github.com/vlone310/advent-of-code/pqueue"
)

type coords struct {
	x, y int
}

type Node struct {
	point    coords
	g, h, f  int
	previous *Node
}

var directions = [4][2]int{
	{0, -1},
	{0, 1},
	{-1, 0},
	{1, 0},
}

func ramRun(input string, gridSize, bytesSize int) int {
	corrupted := getCorruptedCoordinates(input, bytesSize)
	end := astar(gridSize, corrupted)

	if end != nil {
		return end.g
	}

	return 0
}

func ramRunP2(input string, gridZise int) string {
	bytes := strings.Split(input, "\n")
	start := 0
	end := len(bytes) - 2

	for end-start > 1 {
		middle := (end + start) / 2
		corrupted := getCorruptedCoordinates(input, middle)
		path := astar(gridZise, corrupted)
		if path == nil {
			end = middle
		} else {
			start = middle
		}
	}

	return bytes[start]
}

func getCorruptedCoordinates(input string, size int) map[coords]bool {
	res := make(map[coords]bool)
	rows := strings.Split(input, "\n")

	for i := 0; i < size; i++ {
		rawCoords := strings.Split(rows[i], ",")
		if len(rawCoords) < 2 {
			continue
		}

		res[coords{helpers.ParseInt(rawCoords[1]), helpers.ParseInt(rawCoords[0])}] = true
	}

	return res
}

func heuristic(a, b coords) int {
	return int(math.Abs(float64(a.x-b.x)) + math.Abs(float64(a.y-b.y)))
}

func astar(gridSize int, corrupted map[coords]bool) *Node {
	lessFunc := func(a, b *Node) bool {
		return a.f < b.f
	}

	pq := pqueue.New(lessFunc)
	start := coords{0, 0}
	end := coords{gridSize, gridSize}

	startNode := &Node{point: start, g: 0, h: heuristic(start, end), f: heuristic(start, end)}

	pq.PushItem(startNode)

	visited := make(map[coords]int)

	for pq.Len() > 0 {
		current := pq.PopItem()

		if current.point == end {
			return current
		}

		visited[current.point] = current.g

		for _, dir := range directions {
			point := coords{current.point.x + dir[0], current.point.y + dir[1]}
			g := current.g + 1
			h := heuristic(point, end)
			f := g + h

			if point.x < 0 || point.y < 0 || point.x > gridSize || point.y > gridSize {
				continue
			}

			if corrupted[point] {
				continue
			}
			if prevCost, exists := visited[point]; exists && prevCost <= g {
				continue
			}

			pq.PushItem(&Node{
				point:    point,
				g:        g,
				h:        h,
				f:        f,
				previous: current,
			})

		}
	}

	return nil
}

func traversePath(node *Node) map[coords]bool {
	path := map[coords]bool{}
	cur := node

	for cur != nil {
		path[cur.point] = true
		cur = cur.previous
	}

	return path
}

func printGrid(gridSise int, corrupted map[coords]bool, path map[coords]bool) {
	for i := 0; i <= gridSise; i++ {
		row := ""
		for j := 0; j <= gridSise; j++ {
			next := "."
			point := coords{i, j}
			if corrupted[point] {
				next = "#"
			}

			if path[point] {
				next = "O"
			}

			row += next
		}
		fmt.Println(row)
	}
}
