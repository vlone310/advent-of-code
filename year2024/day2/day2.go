package year2024

import (
	"math"
	"strconv"
	"strings"
)

const (
	none = iota
	increase
	decrease
)

func RedNosedReports(input string) (safeReports int) {
	reports := strings.Split(input, "\n")

	for _, row := range reports {
		reportRow := prepareRow(row)

		if len(reportRow) < 2 {
			continue
		}

		state := none
		safeReports += 1

		for idx, level := range reportRow[1:] {
			prevLevel := reportRow[idx]

			if state == none {
				if level > prevLevel {
					state = increase
				} else {
					state = decrease
				}
			}

			if (state == increase && level < prevLevel) || (state == decrease && level > prevLevel) {
				safeReports -= 1
				break
			}

			diff := math.Abs(float64(level - prevLevel))

			if diff > 3 || diff == 0 {
				safeReports -= 1
				break
			}

		}

	}

	return safeReports
}

func RedNosedReportsP2(input string) (safeReports int) {
	reports := strings.Split(input, "\n")

	var solveRow func(row []int, removedIdx int) int

	solveRow = func(row []int, removedIdx int) int {
		reportRow := make([]int, len(row))
		copy(reportRow, row)

		if removedIdx > len(row)-1 {
			return 0
		}

		if removedIdx != -1 {
			reportRow = reportRow[:removedIdx+copy(reportRow[removedIdx:], reportRow[removedIdx+1:])]
		}

		state := none

		for idx, level := range reportRow[1:] {
			prevLevel := reportRow[idx]

			if state == none {
				if level > prevLevel {
					state = increase
				} else {
					state = decrease
				}
			}

			if (state == increase && level < prevLevel) || (state == decrease && level > prevLevel) {
				return solveRow(row, removedIdx+1)
			}

			diff := math.Abs(float64(level - prevLevel))

			if diff > 3 || diff == 0 {
				return solveRow(row, removedIdx+1)
			}
		}

		return 1
	}

	for _, row := range reports {
		reportRow := prepareRow(row)

		if len(reportRow) < 2 {
			continue
		}

		safeReports += solveRow(reportRow, -1)
	}

	return safeReports
}

func prepareRow(row string) (intSlice []int) {
	slice := strings.Split(row, " ")

	for _, el := range slice {
		elInt, _ := strconv.Atoi(el)
		intSlice = append(intSlice, elInt)
	}

	return intSlice
}
