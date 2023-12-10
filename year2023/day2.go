package year2023

import (
	"advent-of-code/helpers"
	"math"
	"strconv"
	"strings"
)

func GetId(input string) (int64, error) {
	stringId := strings.Split(strings.Split(input, ":")[0], " ")[1]

	return strconv.ParseInt(stringId, 10, 64)
}

func CubeConundrum(input string) int64 {
	var sum int64
	maxCubes := map[string]int64{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	gamesSlice := strings.Split(input, "\n")

	for _, game := range gamesSlice {
		if game == "" {
			continue
		}
		id, err := GetId(game)

		if err != nil {
			continue
		}

		rounds := strings.Split(strings.Split(game, ": ")[1], "; ")

		for _, round := range rounds {
			condition := strings.Split(round, ", ")

			for _, values := range condition {
				value := strings.Split(values, " ")

				cubesCount, _ := strconv.ParseInt(value[0], 10, 64)
				cubeColor := value[1]

				if cubesCount > maxCubes[cubeColor] {
					id = 0
					break
				}
			}
		}

		sum += id

	}

	return sum
}

func CubeConundrumP2(input string) float64 {
	var sum float64

	gamesSlice := strings.Split(input, "\n")

	for _, game := range gamesSlice {
		cubesMaxes := map[string]float64{
			"red":   1,
			"green": 1,
			"blue":  1,
		}

		if game == "" {
			continue
		}

		round := strings.Split(game, ": ")[1]
		values := helpers.SplitByMultipleSeparators(round, `; |, `)

		for _, value := range values {
			valueSlice := strings.Split(value, " ")
			cubesCount, _ := strconv.ParseFloat(valueSlice[0], 64)
			cubesColor := valueSlice[1]

			cubesMaxes[cubesColor] = math.Max(cubesMaxes[cubesColor], cubesCount)
		}

		res := cubesMaxes["red"] * cubesMaxes["green"] * cubesMaxes["blue"]
		sum += res
	}

	return sum
}
