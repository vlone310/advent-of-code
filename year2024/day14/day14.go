package year2024

import (
	"regexp"
	"strings"

	"github.com/vlone310/advent-of-code/helpers"
)

type coords struct {
	x, y int
}

type instruction struct {
	position coords
	velocity coords
}

var christmasTreePattern = []string{
	"   X   ",
	"  XXX  ",
	" XXXXX ",
	"XXXXXXX",
}

func restroomRedoubt(input string, width, height, seconds int) int {
	instructions := parseInputToInstructions(input)
	robotsPos := []coords{}

	for _, instruction := range instructions {
		destX := (instruction.velocity.x * seconds) + instruction.position.x
		destY := (instruction.velocity.y * seconds) + instruction.position.y
		resX := destX % width
		resY := destY % height

		if resX < 0 {
			resX = width + resX
		}
		if resY < 0 {
			resY = height + resY
		}
		robotsPos = append(robotsPos, coords{resX, resY})
	}

	middleX := width / 2
	middleY := height / 2

	positionsCount := map[string]int{}
	res := 1

	for _, pos := range robotsPos {
		if pos.x < middleX && pos.y < middleY {
			positionsCount["tl"] += 1
		} else if pos.x > middleX && pos.y < middleY {
			positionsCount["tr"] += 1
		} else if pos.x < middleX && pos.y > middleY {
			positionsCount["bl"] += 1
		} else if pos.x > middleX && pos.y > middleY {
			positionsCount["br"] += 1
		}
	}

	for _, v := range positionsCount {
		res *= v
	}

	return res
}

func restroomRedoubtP2(input string, width, height int) int {
	defer helpers.Timer()()
	var seconds int
	var res int
	instructions := parseInputToInstructions(input)

	for res == 0 {
		robotsPos := []coords{}

		for _, instruction := range instructions {
			destX := (instruction.velocity.x * seconds) + instruction.position.x
			destY := (instruction.velocity.y * seconds) + instruction.position.y
			resX := destX % width
			resY := destY % height

			if resX < 0 {
				resX = width + resX
			}
			if resY < 0 {
				resY = height + resY
			}
			robotsPos = append(robotsPos, coords{resX, resY})
		}

		grid := make([][]string, height)
		for i := range grid {
			grid[i] = make([]string, width)
			for j := range grid[i] {
				grid[i][j] = " "
			}
		}

		for _, coord := range robotsPos {
			grid[coord.y][coord.x] = "X"
		}

		var currPatternEntryI int

		for _, row := range grid {
			if currPatternEntryI == len(christmasTreePattern) {
				res = seconds
				break
			}
			if strings.Contains(strings.Join(row, ""), christmasTreePattern[currPatternEntryI]) {
				currPatternEntryI++
			} else {
				currPatternEntryI = 0
			}
		}
		seconds++
	}

	return res
}

func parseInputToInstructions(input string) []instruction {
	rows := strings.Split(input, "\n")
	res := make([]instruction, 0, len(rows))

	for _, row := range rows {
		re := regexp.MustCompile(`[-]?\d+`)

		matches := re.FindAllString(row, -1)

		if len(matches) < 4 {
			continue
		}

		res = append(res, instruction{
			position: coords{helpers.ParseInt(matches[0]), helpers.ParseInt(matches[1])},
			velocity: coords{helpers.ParseInt(matches[2]), helpers.ParseInt(matches[3])},
		})

	}

	return res
}
