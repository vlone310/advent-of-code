package year2024

import (
	"math"
	"strings"

	"github.com/vlone310/advent-of-code/pqueue"
)

type coords struct {
	x, y int
}

type direction int

type grid struct {
	maze  [][]byte
	start coords
	end   coords
	dir   direction
}

type Node struct {
	point    coords
	g, h, f  int // g: cost from start, h: heuristic (Manhattan distance), f: total cost (g + h)
	dir      direction
	previous *Node
}

const (
	up direction = iota
	right
	down
	left
)

var directionsMap = map[direction]coords{
	up:    {-1, 0},
	right: {0, 1},
	down:  {1, 0},
	left:  {0, -1},
}

func reindeerMaze(input string) int {
	grid := parseGrid(input)
	return astar(grid)[0].g
}

func reindeerMazeP2(input string) int {
	grid := parseGrid(input)
	return traverseNodes(astar(grid))
}

func heuristic(a, b coords) int {
	return int(math.Abs(float64(a.x-b.x)) + math.Abs(float64(a.y-b.y)))
}

func astar(grid grid) []*Node {
	lessFunc := func(a, b *Node) bool {
		return a.f < b.f
	}

	pq := pqueue.New(lessFunc)

	startNode := &Node{point: grid.start, g: 0, h: heuristic(grid.start, grid.end), f: heuristic(grid.start, grid.end), dir: grid.dir}

	pq.PushItem(startNode)

	visited := make(map[coords]map[direction]int)
	validPaths := []*Node{}

	for pq.Len() > 0 {
		current := pq.PopItem()

		if current.point == grid.end {
			validPaths = append(validPaths, current)
			continue
		}

		for _, dir := range []direction{current.dir, rotateLeft(current.dir), rotateRight(current.dir)} {
			point := coords{current.point.x + directionsMap[dir].x, current.point.y + directionsMap[dir].y}
			g := current.g + 1
			if dir != current.dir {
				g += 1000
			}
			h := heuristic(point, grid.end)
			f := g + h
			if bestG, exists := visited[point][dir]; grid.maze[point.x][point.y] == '#' || (exists && g > bestG) {
				continue
			}

			if visited[point] == nil {
				visited[point] = make(map[direction]int)
			}
			visited[point][dir] = g

			pq.PushItem(&Node{
				point:    point,
				g:        g,
				h:        h,
				f:        f,
				dir:      dir,
				previous: current,
			})

		}
	}

	pathsWithMinScore := getPathsWithMinScore(validPaths)

	return pathsWithMinScore
}

func traverseNodes(nodes []*Node) int {
	unique := map[coords]bool{}

	for _, node := range nodes {
		curNode := node

		for curNode != nil {
			if _, exists := unique[curNode.point]; !exists {
				unique[curNode.point] = true
			}
			curNode = curNode.previous
		}
	}

	return len(unique)
}

func getPathsWithMinScore(nodes []*Node) []*Node {

	minG := int(math.Inf(1))

	for _, path := range nodes {
		if path.g < minG {
			minG = path.g
		}
	}
	pathsWithMinScore := []*Node{}

	for _, path := range nodes {
		if path.g == minG {
			pathsWithMinScore = append(pathsWithMinScore, path)
		}
	}

	return pathsWithMinScore
}

func parseGrid(input string) grid {
	rows := strings.Split(input, "\n")
	maze := make([][]byte, 0, len(rows))
	var start coords
	var end coords

	for x, row := range rows {
		maze = append(maze, []byte(row))
		for y, el := range row {
			if byte(el) == 'S' {
				start = coords{x, y}
			}
			if byte(el) == 'E' {
				end = coords{x, y}
			}
		}
	}

	return grid{
		maze:  maze,
		start: start,
		end:   end,
		dir:   right,
	}
}

func rotateLeft(dir direction) direction {
	switch dir {
	case up:
		return left
	case left:
		return down
	case down:
		return right
	case right:
		return up
	default:
		return -1
	}
}

func rotateRight(dir direction) direction {
	switch dir {
	case up:
		return right
	case left:
		return up
	case down:
		return left
	case right:
		return down
	default:
		return -1
	}
}
