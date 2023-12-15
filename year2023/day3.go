package year2023

import (
	"advent-of-code/helpers"
	"strconv"
	"strings"
)

func GearRatios(input string) int64 {
	rows := strings.Split(input, "\n")

	var sum int64

	for rowIndex, row := range rows {
		var curNum string
		isAdjacent := false

		for colIndex, element := range []byte(row) {
			if helpers.IsDigit(element) {
				curNum += string(element)

				if !isAdjacent {
					isAdjacent = isAdjacentTo(rowIndex, colIndex, rows)
				}

				if colIndex == len(row)-1 || !helpers.IsDigit(byte(row[colIndex+1])) {
					if isAdjacent {
						num, _ := strconv.ParseInt(curNum, 10, 64)
						sum += num
					}
					curNum = ""
					isAdjacent = false
				}
			}
		}
	}

	return sum
}

func GearRatiosP2(input string) int64 {
	rows := strings.Split(input, "\n")

	var sum int64

	for rowIndex, row := range rows {
		for colIndex, element := range []byte(row) {
			if string(element) == "*" {
				sum += getAdjacentNumsSum(rowIndex, colIndex, rows)
			}
		}
	}

	return sum
}

func isAdjacentTo(rowIndex int, colIndex int, rows []string) bool {
	indicies := [8][2]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
		{1, 1},
		{-1, -1},
		{1, -1},
		{-1, 1},
	}

	for _, indiciesPair := range indicies {
		currRowIndex := indiciesPair[0] + rowIndex
		currColIndex := indiciesPair[1] + colIndex

		if helpers.IsOutOfBound(currRowIndex, len(rows)) || helpers.IsOutOfBound(currColIndex, len(rows[0])) {
			continue
		}

		if rows[currRowIndex][currColIndex] != '.' && !helpers.IsDigit(rows[currRowIndex][currColIndex]) {
			return true
		}
	}

	return false
}

func getAdjacentNumsSum(rowIndex int, colIndex int, rows []string) int64 {
	indicies := [8][2]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
		{1, 1},
		{-1, -1},
		{1, -1},
		{-1, 1},
	}
	numbersFound := []int64{}
	visited := map[string]bool{}

	for _, indiciesPair := range indicies {
		currRowIndex := indiciesPair[0] + rowIndex
		currColIndex := indiciesPair[1] + colIndex

		if helpers.IsOutOfBound(currRowIndex, len(rows)) || helpers.IsOutOfBound(currColIndex, len(rows[0])) {
			continue
		}

		if _, ok := visited[getKey(currRowIndex, currColIndex)]; !ok && helpers.IsDigit(rows[currRowIndex][currColIndex]) {
			visited[getKey(currRowIndex, currColIndex)] = true
			visited[getKey(currRowIndex, currColIndex-1)] = true
			visited[getKey(currRowIndex, currColIndex+1)] = true
			adjacentNumberString := findNumToLeft(currRowIndex, currColIndex-1, rows, "") + string(rows[currRowIndex][currColIndex]) + findNumToRight(currRowIndex, currColIndex+1, rows, "")

			adjacentNumber, err := strconv.ParseInt(adjacentNumberString, 10, 64)

			if err != nil {
				continue
			}

			numbersFound = append(numbersFound, adjacentNumber)
		}
	}

	if len(numbersFound) == 2 {
		var sum int64 = 1
		for _, number := range numbersFound {
			sum *= number
		}
		return sum
	}

	return 0
}

func findNumToRight(rowIndex int, colIndex int, rows []string, num string) string {
	if helpers.IsOutOfBound(rowIndex, len(rows)) || helpers.IsOutOfBound(colIndex, len(rows[0])) || !helpers.IsDigit(rows[rowIndex][colIndex]) {
		return num
	}

	return findNumToRight(rowIndex, colIndex+1, rows, num+string(rows[rowIndex][colIndex]))
}

func findNumToLeft(rowIndex int, colIndex int, rows []string, num string) string {
	if helpers.IsOutOfBound(rowIndex, len(rows)) || helpers.IsOutOfBound(colIndex, len(rows[0])) || !helpers.IsDigit(rows[rowIndex][colIndex]) {
		return num
	}

	return findNumToLeft(rowIndex, colIndex-1, rows, string(rows[rowIndex][colIndex])+num)
}

func getKey(num1 int, num2 int) string {
	return strconv.Itoa(num1) + strconv.Itoa(num2)
}
