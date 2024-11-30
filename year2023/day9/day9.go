package year2023

import (
	"strconv"
	"strings"
)

// TODO: Can be optimizied
func MirageMaintenance(input string) (res int) {
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		numsLine := getNums(line)
		sequence := [][]int{numsLine}
		i := 0

		for len(sequence) != len(numsLine) {
			nextRow := []int{}

			for j := 1; j < len(sequence[i]); j++ {
				cur := sequence[i][j]
				prev := sequence[i][j-1]

				nextRow = append(nextRow, cur-prev)
			}
			i++
			sequence = append(sequence, nextRow)
		}

		curV := 0

		for i := len(sequence) - 1; i >= 0; i-- {
			curV += sequence[i][len(sequence[i])-1]
		}

		res += curV

	}
	return
}

func MirageMaintenanceP2(input string) (res int) {
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		numsLine := getNums(line)
		sequence := [][]int{numsLine}
		i := 0

		for len(sequence) != len(numsLine) {
			nextRow := []int{}

			for j := 1; j < len(sequence[i]); j++ {
				cur := sequence[i][j]
				prev := sequence[i][j-1]

				nextRow = append(nextRow, cur-prev)
			}
			i++
			sequence = append(sequence, nextRow)
		}

		curV := 0

		for i := len(sequence) - 1; i >= 0; i-- {
			curV = sequence[i][0] - curV
		}

		res += curV

	}
	return
}

func getNums(input string) (res []int) {
	for _, numStr := range strings.Split(input, " ") {
		num, err := strconv.Atoi(numStr)

		if err != nil {
			panic(err)
		}

		res = append(res, num)
	}
	return
}
