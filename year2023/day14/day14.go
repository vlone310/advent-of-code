package year2023

import (
	"strings"
)

func ParabolicReflectorDish(input string) (sum int) {
	matrix := strings.Split(input, "\n")

	rotateNorth(&matrix)

	for i, row := range matrix {
		for _, el := range row {
			if el == 'O' {
				sum += len(matrix) - i
			}
		}
	}

	return sum
}

func ParabolicReflectorDishP2(input string) (sum int) {
	matrix := strings.Split(input, "\n")
	results := make([]int, 0)
	numCycles := 1000000000

	for iterCount := 0; iterCount < numCycles; iterCount++ {
		rotateNorth(&matrix)
		rotateWest(&matrix)
		rotateSouth(&matrix)
		rotateEast(&matrix)

		curSum := 0

		for i, row := range matrix {
			for _, el := range row {
				if el == 'O' {
					curSum += len(matrix) - i
				}
			}
		}

		results = append(results, curSum)

		if frequency, offset, ok := getSequence(results); ok {
			mod := (numCycles - offset) % frequency
			sum = results[offset+mod-1]
			break
		}
	}

	return
}

func getSequence(values []int) (frequency, offset int, ok bool) {
	if len(values) > 10 {
		var reps []int
		for i := 0; i < len(values)-1; i++ {
			if values[i] == values[len(values)-1] {
				reps = append(reps, i)
				if len(reps) == 3 {
					if reps[2]-reps[1] == reps[1]-reps[0] && reps[2]-reps[1] > 1 {
						ok = true
						for i := 0; i < reps[1]-reps[0]; i++ {
							if values[reps[0]+i] != values[reps[1]+i] {
								ok = false
							}
						}
						if ok {
							frequency = reps[2] - reps[1]
							offset = reps[0]
						}
					}
				}
			}
		}
	}
	return frequency, offset, ok
}

var northCache = make(map[string]string)
var westCache = make(map[string]string)
var southCache = make(map[string]string)
var eastCache = make(map[string]string)

func rotateNorth(matrix *[]string) {
	if val, ok := northCache[strings.Join(*matrix, "\n")]; ok {
		*matrix = strings.Split(val, "\n")
		return
	}
	prevMatrix := strings.Join(*matrix, "\n")

	for i, row := range *matrix {
		for j, el := range row {
			if el == 'O' {
				curRowIndex := i

				for curRowIndex > 0 {
					if (*matrix)[curRowIndex-1][j] == '#' || (*matrix)[curRowIndex-1][j] == 'O' {
						break
					} else {
						(*matrix)[curRowIndex] = replaceByIndex((*matrix)[curRowIndex], j, ".")
						(*matrix)[curRowIndex-1] = replaceByIndex((*matrix)[curRowIndex-1], j, "O")
						curRowIndex--
					}
				}
			}
		}
	}

	northCache[prevMatrix] = strings.Join(*matrix, "\n")
}

func rotateWest(matrix *[]string) {
	if val, ok := westCache[strings.Join(*matrix, "\n")]; ok {
		*matrix = strings.Split(val, "\n")
		return
	}
	prevMatrix := strings.Join(*matrix, "\n")

	for i, row := range *matrix {
		for j, el := range row {
			if el == 'O' {
				curColIndex := j

				for curColIndex > 0 {
					if (*matrix)[i][curColIndex-1] == '#' || (*matrix)[i][curColIndex-1] == 'O' {
						break
					} else {
						(*matrix)[i] = replaceByIndex((*matrix)[i], curColIndex, ".")
						(*matrix)[i] = replaceByIndex((*matrix)[i], curColIndex-1, "O")
						curColIndex--
					}
				}
			}
		}
	}

	westCache[prevMatrix] = strings.Join(*matrix, "\n")
}

func rotateSouth(matrix *[]string) {
	if val, ok := southCache[strings.Join(*matrix, "\n")]; ok {
		*matrix = strings.Split(val, "\n")
		return
	}
	prevMatrix := strings.Join(*matrix, "\n")

	for i := len(*matrix) - 1; i >= 0; i-- {
		for j, el := range (*matrix)[i] {
			if el == 'O' {
				curRowIndex := i

				for curRowIndex < len(*matrix)-1 {
					if (*matrix)[curRowIndex+1][j] == '#' || (*matrix)[curRowIndex+1][j] == 'O' {
						break
					} else {
						(*matrix)[curRowIndex] = replaceByIndex((*matrix)[curRowIndex], j, ".")
						(*matrix)[curRowIndex+1] = replaceByIndex((*matrix)[curRowIndex+1], j, "O")
						curRowIndex++
					}
				}
			}
		}
	}

	southCache[prevMatrix] = strings.Join(*matrix, "\n")
}

func rotateEast(matrix *[]string) {
	if val, ok := eastCache[strings.Join(*matrix, "\n")]; ok {
		*matrix = strings.Split(val, "\n")
		return
	}
	prevMatrix := strings.Join(*matrix, "\n")

	for i, row := range *matrix {
		for j := len(row) - 1; j >= 0; j-- {
			if row[j] == 'O' {
				curColIndex := j

				for curColIndex < len(row)-1 {
					if (*matrix)[i][curColIndex+1] == '#' || (*matrix)[i][curColIndex+1] == 'O' {
						break
					} else {
						(*matrix)[i] = replaceByIndex((*matrix)[i], curColIndex, ".")
						(*matrix)[i] = replaceByIndex((*matrix)[i], curColIndex+1, "O")
						curColIndex++
					}
				}
			}
		}
	}

	eastCache[prevMatrix] = strings.Join(*matrix, "\n")
}

func replaceByIndex(str string, i int, replacement string) string {
	return str[:i] + replacement + str[i+1:]
}
