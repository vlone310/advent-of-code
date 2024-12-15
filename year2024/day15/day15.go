package year2024

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/vlone310/advent-of-code/helpers"
)

type coords struct {
	x, y int
}

type grid struct {
	robotPos coords
	whMap    [][]byte
	moves    []byte
}

var positions = map[byte]coords{
	'^': {-1, 0},
	'v': {1, 0},
	'<': {0, -1},
	'>': {0, 1},
}

func warehouseWoes(input string) int {
	grid := newGrid(input)

	for _, move := range grid.moves {
		grid.moveRobot(move)
	}

	return grid.calculateSum()
}

func warehouseWoesP2(input string) int {
	grid := newGrid(input)
	grid.scale()

	for _, move := range grid.moves {
		grid.moveRobot(move)
	}

	return grid.calculateSum()
}

func newGrid(input string) *grid {
	rawMapAndMoves := strings.Split(input, "\n\n")
	rawMap := strings.Split(rawMapAndMoves[0], "\n")

	moves := []byte(strings.ReplaceAll(rawMapAndMoves[1], "\n", ""))
	var robotPos coords
	whMap := make([][]byte, 0, len(rawMap))

	for x, row := range rawMap {
		whMap = append(whMap, []byte(row))

		for y, col := range []byte(row) {
			if col == '@' {
				robotPos = coords{x, y}
			}
		}
	}

	return &grid{
		robotPos: robotPos,
		whMap:    whMap,
		moves:    moves,
	}
}

func (g grid) calculateSum() int {
	var res int

	for x, row := range g.whMap {
		for y, el := range row {
			if el == 'O' || el == '[' {
				res += 100*x + y
			}
		}
	}

	return res
}

func (g *grid) moveRobot(move byte) {
	nextXY, exists := positions[move]
	if !exists {
		log.Fatalf("Unknown move: %c", move)
	}

	nextRobotPos := coords{g.robotPos.x + nextXY.x, g.robotPos.y + nextXY.y}
	nextEl := g.whMap[nextRobotPos.x][nextRobotPos.y]

	// wall, do nothing
	if nextEl == '#' {
		return
	}

	// simple positions swap with empty square
	if nextEl == '.' {
		g.whMap[nextRobotPos.x][nextRobotPos.y], g.whMap[g.robotPos.x][g.robotPos.y] = g.whMap[g.robotPos.x][g.robotPos.y], g.whMap[nextRobotPos.x][nextRobotPos.y]
		g.robotPos = nextRobotPos
		return
	}

	// any move with box O OR move left(<) and right(>) with box []
	if (nextEl == 'O') || ((nextEl == '[' || nextEl == ']') && (move == '<' || move == '>')) {
		curCoords := nextRobotPos

		for g.whMap[curCoords.x][curCoords.y] != '.' && g.whMap[curCoords.x][curCoords.y] != '#' {
			curCoords = coords{curCoords.x + nextXY.x, curCoords.y + nextXY.y}
		}

		if g.whMap[curCoords.x][curCoords.y] == '#' {
			return
		}

		for curCoords.x != g.robotPos.x || curCoords.y != g.robotPos.y {
			nextCoords := coords{curCoords.x - nextXY.x, curCoords.y - nextXY.y}
			g.whMap[curCoords.x][curCoords.y], g.whMap[nextCoords.x][nextCoords.y] = g.whMap[nextCoords.x][nextCoords.y], g.whMap[curCoords.x][curCoords.y]
			curCoords = nextCoords
		}

		g.robotPos = nextRobotPos
	}

	// moves up (^) and down (v) with box []
	if (nextEl == '[' || nextEl == ']') && (move == '^' || move == 'v') {

		boxCluster := g.findBoxCluster(move, nextRobotPos)
		if boxCluster == nil {
			return
		}

		for _, val := range boxCluster {
			nextPosCoords := positions[move]
			next := coords{val.x + nextPosCoords.x, val.y + nextPosCoords.y}
			g.whMap[val.x][val.y], g.whMap[next.x][next.y] = g.whMap[next.x][next.y], g.whMap[val.x][val.y]
		}
		g.whMap[nextRobotPos.x][nextRobotPos.y], g.whMap[g.robotPos.x][g.robotPos.y] = g.whMap[g.robotPos.x][g.robotPos.y], g.whMap[nextRobotPos.x][nextRobotPos.y]
		g.robotPos = nextRobotPos
	}
}

func (g grid) findBoxCluster(move byte, curCoords coords) []coords {
	boxes := make(map[coords]bool)
	queue := helpers.Queue[coords]{curCoords}
	nextCoords := positions[move]

	for len(queue) > 0 {
		curXY := queue.Pop()
		cur := g.whMap[curXY.x][curXY.y]

		if _, exists := boxes[curXY]; exists {
			continue
		}

		if cur == ']' {
			boxes[curXY] = true
			queue.Push(coords{curXY.x, curXY.y - 1})
			queue.Push(coords{curXY.x + nextCoords.x, curXY.y + nextCoords.y})
		} else if cur == '[' {
			boxes[curXY] = true
			queue.Push(coords{curXY.x, curXY.y + 1})
			queue.Push(coords{curXY.x + nextCoords.x, curXY.y + nextCoords.y})
		} else if cur == '#' {
			boxes = nil
			break
		}
	}

	if boxes == nil {
		return nil
	}

	keys := make([]coords, 0, len(boxes))
	for key := range boxes {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		if keys[i].x != keys[j].x {
			if move == 'v' {
				return keys[i].x > keys[j].x
			}
			return keys[i].x < keys[j].x
		}
		if move == 'v' {
			return keys[i].y > keys[j].y
		}
		return keys[i].y < keys[j].y
	})

	return keys
}

func (g *grid) scale() {
	newWhMap := make([][]byte, 0, len(g.whMap))

	for x, row := range g.whMap {
		newRow := []byte{}
		for _, el := range row {
			switch el {
			case '@':
				g.robotPos = coords{x, len(newRow)}
				newRow = append(newRow, '@', '.')
			case '#':
				newRow = append(newRow, '#', '#')
			case 'O':
				newRow = append(newRow, '[', ']')
			default:
				newRow = append(newRow, '.', '.')
			}
		}
		newWhMap = append(newWhMap, newRow)
	}
	g.whMap = newWhMap
}

func (g grid) print() {
	fmt.Println("robotPos", g.robotPos)
	fmt.Println("moves", string(g.moves))

	for _, row := range g.whMap {
		fmt.Println(string(row))
	}
}
