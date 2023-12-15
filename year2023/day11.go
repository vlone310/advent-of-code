package year2023

import (
	"math"
	"strings"
)

func CosmicExpansion(input string) (sum int) {
	rows := strings.Split(input, "\n")
	galaxies := [][2]int{}
	notEmptyCols := map[int]bool{}
	notEmptyRows := map[int]bool{}

	// Searching for galaxies and empty rows/cols
	for rowIndex, row := range rows {
		for colIndex, el := range row {
			if el == '#' {
				galaxies = append(galaxies, [2]int{rowIndex, colIndex})
				notEmptyRows[rowIndex] = true
				notEmptyCols[colIndex] = true
			}
		}
	}

	// Calculating sum
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			galaxy1 := galaxies[i]
			galaxy2 := galaxies[j]
			curCounter := 0

			startRow := int(math.Min(float64(galaxy1[0]), float64(galaxy2[0])))
			endRow := int(math.Max(float64(galaxy1[0]), float64(galaxy2[0])))

			for i := startRow + 1; i < endRow; i++ {
				if _, ok := notEmptyRows[i]; !ok {
					curCounter++
				}
			}

			startCol := int(math.Min(float64(galaxy1[1]), float64(galaxy2[1])))
			endCol := int(math.Max(float64(galaxy1[1]), float64(galaxy2[1])))

			for i := startCol + 1; i < endCol; i++ {
				if _, ok := notEmptyCols[i]; !ok {
					curCounter++
				}
			}

			sum += int(math.Abs(float64(galaxy1[0]-galaxy2[0]))) + int(math.Abs(float64(galaxy1[1]-galaxy2[1]))) + curCounter
		}
	}

	return
}

func CosmicExpansionP2(input string) (sum int) {
	rows := strings.Split(input, "\n")
	galaxies := [][2]int{}
	notEmptyCols := map[int]bool{}
	notEmptyRows := map[int]bool{}

	// Searching for galaxies and empty rows/cols
	for rowIndex, row := range rows {
		for colIndex, el := range row {
			if el == '#' {
				galaxies = append(galaxies, [2]int{rowIndex, colIndex})
				notEmptyRows[rowIndex] = true
				notEmptyCols[colIndex] = true
			}
		}
	}

	// Calculating sum
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			galaxy1 := galaxies[i]
			galaxy2 := galaxies[j]
			curCounter := 0

			startRow := int(math.Min(float64(galaxy1[0]), float64(galaxy2[0])))
			endRow := int(math.Max(float64(galaxy1[0]), float64(galaxy2[0])))

			for i := startRow + 1; i < endRow; i++ {
				if _, ok := notEmptyRows[i]; !ok {
					curCounter += 999999
				}
			}

			startCol := int(math.Min(float64(galaxy1[1]), float64(galaxy2[1])))
			endCol := int(math.Max(float64(galaxy1[1]), float64(galaxy2[1])))

			for i := startCol + 1; i < endCol; i++ {
				if _, ok := notEmptyCols[i]; !ok {
					curCounter += 999999
				}
			}

			sum += int(math.Abs(float64(galaxy1[0]-galaxy2[0]))) + int(math.Abs(float64(galaxy1[1]-galaxy2[1]))) + curCounter
		}
	}

	return
}
