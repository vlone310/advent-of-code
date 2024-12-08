package year2024

import (
	"strings"

	"github.com/vlone310/advent-of-code/helpers"
)

type coords struct {
	x, y int
}

func resonantCollinearity(input string) int {
	grid := getGrid(input)
	antennasToCoords := map[rune][]coords{}
	antinodes := map[coords]bool{}

	for x, row := range grid {
		for y, el := range row {
			// skipping empty element, we want to proccess only antennas
			if el == '.' {
				continue
			}

			// antenna
			similarAntennaCoords, exists := antennasToCoords[el]
			if exists {
				for _, c := range similarAntennaCoords {
					// getting absolute distance between current antenna and similar antenna
					distanceX, distanceY := x-c.x, y-c.y
					// calculate antinodes indicies
					prevAntinodeX, prevAntinodeY := c.x-distanceX, c.y-distanceY
					nextAntinodeX, nextAntinodeY := x+distanceX, y+distanceY

					for _, ac := range []coords{{prevAntinodeX, prevAntinodeY}, {nextAntinodeX, nextAntinodeY}} {
						if !helpers.IsOutOfBound(ac.x, len(grid)) &&
							!helpers.IsOutOfBound(ac.y, len(grid[ac.x])) &&
							grid[ac.x][ac.y] != el {
							antinodes[coords{ac.x, ac.y}] = true
						}
					}

				}
				antennasToCoords[el] = append(antennasToCoords[el], coords{x, y})
			} else {
				antennasToCoords[el] = []coords{{x, y}}
			}

		}
	}

	return len(antinodes)
}

func resonantCollinearityP2(input string) int {
	grid := getGrid(input)
	antennasToCoords := map[rune][]coords{}
	antinodes := map[coords]bool{}

	for x, row := range grid {
		for y, el := range row {
			// skipping empty element, we want to proccess only antennas
			if el == '.' {
				continue
			}

			// antenna
			similarAntennaCoords, exists := antennasToCoords[el]
			if exists {
				for _, c := range similarAntennaCoords {
					// getting absolute distance between current antenna and similar antenna
					distanceX, distanceY := x-c.x, y-c.y
					// calculate antinodes indicies
					prevAntinodeX, prevAntinodeY := c.x-distanceX, c.y-distanceY
					nextAntinodeX, nextAntinodeY := x+distanceX, y+distanceY
					antinodes[coords{c.x, c.y}] = true
					antinodes[coords{x, y}] = true

					for _, ac := range [][]coords{{{prevAntinodeX, prevAntinodeY}, {-distanceX, -distanceY}}, {{nextAntinodeX, nextAntinodeY}, {distanceX, distanceY}}} {
						tempX := ac[0].x
						tempY := ac[0].y
						for !helpers.IsOutOfBound(tempX, len(grid)) && !helpers.IsOutOfBound(tempY, len(grid[tempX])) {
							if grid[tempX][tempY] != el {
								antinodes[coords{tempX, tempY}] = true
							}
							tempX += ac[1].x
							tempY += ac[1].y
						}
					}

				}
				antennasToCoords[el] = append(antennasToCoords[el], coords{x, y})
			} else {
				antennasToCoords[el] = []coords{{x, y}}
			}

		}
	}

	return len(antinodes)
}

func getGrid(input string) [][]rune {
	rows := strings.Split(input, "\n")
	grid := make([][]rune, len(rows)-1)

	for x, row := range rows {
		if row == "" {
			continue
		}

		grid[x] = []rune(row)
	}

	return grid
}
