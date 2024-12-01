package year2024

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

func HistorianHysteria(input string) (totalDistance int) {
	// 3   7, 2   5
	inputSlice := strings.Split(input, "\n")
	locationIdsFirst := make([]int, 0, len(inputSlice))
	locationIdsSecond := make([]int, 0, len(inputSlice))

	for _, ids := range inputSlice {
		idsSlice := strings.Split(ids, "   ")

		if len(idsSlice) != 2 {
			continue
		}

		id1, _ := strconv.Atoi(idsSlice[0])
		id2, _ := strconv.Atoi(idsSlice[1])
		locationIdsFirst = append(locationIdsFirst, id1)
		locationIdsSecond = append(locationIdsSecond, id2)
	}

	slices.Sort(locationIdsFirst)
	slices.Sort(locationIdsSecond)

	for idx, id := range locationIdsFirst {
		totalDistance += int(math.Abs(float64(id - locationIdsSecond[idx])))
	}
	return
}

func HistorianHysteriaP2(input string) (similarityScore int) {
	// 3   7, 2   5
	inputSlice := strings.Split(input, "\n")
	locationIdsFirst := make([]int, 0, len(inputSlice))
	locationIdsCounter := map[int]int{}

	for _, ids := range inputSlice {
		idsSlice := strings.Split(ids, "   ")

		if len(idsSlice) != 2 {
			continue
		}

		id1, _ := strconv.Atoi(idsSlice[0])
		id2, _ := strconv.Atoi(idsSlice[1])
		locationIdsFirst = append(locationIdsFirst, id1)

		if _, exists := locationIdsCounter[id2]; exists {
			locationIdsCounter[id2] += 1
		} else {
			locationIdsCounter[id2] = 1
		}
	}

	for _, id := range locationIdsFirst {
		similarityScore += id * locationIdsCounter[id]
	}

	return
}
