package year2022

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func CalorieCounting(input string) float64 {
	var maxSum float64
	var curSum float64

	stringsList := strings.Split(input, "\n")
	for _, str := range stringsList {
		if str == "" {
			curSum = 0
		} else {
			num, err := strconv.ParseFloat(str, 64)
			if err != nil {
				fmt.Printf("Error converting string to int: %v\n", err)
				continue
			}

			curSum += num

			maxSum = math.Max(maxSum, curSum)
		}
	}

	return maxSum
}

func CalorieCountingP2(input string) float64 {
	var calorieCountsPerElf []float64
	var curSum float64

	stringsList := strings.Split(input, "\n")
	for _, str := range stringsList {
		if str == "" {
			calorieCountsPerElf = append(calorieCountsPerElf, curSum)
			curSum = 0
		} else {
			num, err := strconv.ParseFloat(str, 64)
			if err != nil {
				fmt.Printf("Error converting string to int: %v\n", err)
				continue
			}

			curSum += num
		}
	}

	sort.SliceStable(calorieCountsPerElf, func(i, j int) bool {
		return calorieCountsPerElf[i] > calorieCountsPerElf[j]
	})

	// Get the sum of the three highest numbers
	var sum float64
	for i := 0; i < 3 && i < len(calorieCountsPerElf); i++ {
		sum += calorieCountsPerElf[i]
	}

	return sum
}