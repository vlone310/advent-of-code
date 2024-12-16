package year2023

import (
	"strconv"
	"strings"

	"github.com/vlone310/advent-of-code/helpers"
)

func HotSprings(input string) (sum int) {
	rows := strings.Split(input, "\n")

	for _, row := range rows {
		rowSplit := strings.Split(row, " ")
		template := rowSplit[0]
		arrangements := stringNumsToInt(strings.Split(rowSplit[1], ","))
		sum += calculateArragements(template, arrangements, 0, 0)

	}
	return
}

func HotSpringsP2(input string) (sum int) {
	rows := strings.Split(input, "\n")

	for _, row := range rows {
		rowSplit := strings.Split(row, " ")
		template := repeatStringWithSeparator(rowSplit[0], 5, "?")
		arrangements := stringNumsToInt(strings.Split(repeatStringWithSeparator(rowSplit[1], 5, ","), ","))
		sum += calculateArragements(template, arrangements, 0, 0)
	}

	return
}

func repeatStringWithSeparator(str string, count int, separator string) (res string) {
	for i := 0; i < count; i++ {
		res += str
		if i < count-1 {
			res += separator
		}
	}
	return
}

func stringNumsToInt(nums []string) (res []int) {
	for _, num := range nums {
		if val, err := strconv.Atoi(num); err == nil {
			res = append(res, val)
		}
	}
	return
}

func calculateArragements(template string, arrangements []int, lastEl byte, arrangementPos int) (sum int) {
	if len(arrangements) == 0 && strings.Contains(template, "#") {
		return 0
	}

	if len(arrangements) == 0 {
		return 1
	}

	if template == "" && len(arrangements) > 0 {
		return 0
	}

	currEl := template[0]

	if currEl == '#' && lastEl == '#' {
		return 0
	}

	if currEl == '?' || currEl == '.' {
		sum += calculateArragements(template[1:], arrangements, '.', arrangementPos)
	}

	if currEl == '#' || currEl == '?' {
		if lastEl != '#' {
			isPossibleToInsert := true
			for i := 0; i < arrangements[0]; i++ {
				if helpers.IsOutOfBound(i, len(template)) || (template[i] != '#' && template[i] != '?') {
					isPossibleToInsert = false
					break
				}
			}

			if isPossibleToInsert {
				sum += calculateArragements(template[arrangements[0]:], arrangements[1:], '#', arrangementPos+1)
			}
		}
	}

	return sum
}
