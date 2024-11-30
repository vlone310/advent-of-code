package year2023

import (
	"math"
	"strconv"
	"strings"
)

func IfYouGiveASeedAFertilizer(input string) int {
	lowestLocationNum := math.MaxInt
	seeds := extractSeeds(input)
	graph := extractGraphFromInput(strings.Split(input, "\n")[1:])

	for _, seed := range seeds {

		for _, dots := range graph {
			for _, conditions := range dots {
				destination := conditions[0]
				source := conditions[1]
				ranges := conditions[2]

				if seed >= source && seed < source+ranges {
					seed = destination + (seed - source)
					break
				}
			}
		}

		lowestLocationNum = int(math.Min(float64(lowestLocationNum), float64(seed)))
	}
	return lowestLocationNum
}

func IfYouGiveASeedAFertilizerP2(input string) int {
	seeds := extractSeeds(input) // nth(2n - 1) index is seed, nth(2n) index is range
	graph := extractGraphFromInput(strings.Split(input, "\n"))

	for i := 0; i < math.MaxInt; i++ {
		seed := i
		for graphIndex := len(graph) - 1; graphIndex >= 0; graphIndex-- {
			currentGraph := graph[graphIndex]
			for _, conditions := range currentGraph {
				destination := conditions[0]
				source := conditions[1]
				ranges := conditions[2]

				if seed >= destination && seed < destination+ranges {
					seed = source + (seed - destination)
					break
				}
			}
		}

		for j := 0; j < len(seeds); j += 2 {
			if seed >= seeds[j] && seed < seeds[j]+seeds[j+1] {
				return i
			}
		}
	}

	return 0
}

func extractSeeds(input string) []int {
	startIndex := strings.Index(input, "seeds: ") + len("seeds: ")
	endIndex := strings.Index(input, "\n")

	seedsSlice := input[startIndex:endIndex]

	seeds := strings.Split(seedsSlice, " ")
	numSeeds := make([]int, len(seeds))

	for i, seed := range seeds {
		num, err := strconv.Atoi(seed)

		if err != nil {
			panic(err)
		}

		numSeeds[i] = num
	}

	return numSeeds
}

func extractGraphFromInput(lines []string) [][][]int {
	var mapping [][][]int
	var currGraph [][]int

	for _, line := range lines {
		lineSplit := strings.Split(line, " ")

		if _, err := strconv.Atoi(lineSplit[0]); err != nil {
			if currGraph != nil {
				mapping = append(mapping, currGraph)
			}
			currGraph = nil
		} else {
			currMapValue := make([]int, len(lineSplit))
			for i, num := range lineSplit {
				numInt, err := strconv.Atoi(num)

				if err != nil {
					panic(err)
				}
				currMapValue[i] = numInt

			}

			currGraph = append(currGraph, currMapValue)
		}
	}

	if currGraph != nil {
		mapping = append(mapping, currGraph)
	}

	return mapping
}
