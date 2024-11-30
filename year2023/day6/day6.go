package year2023

import (
	"strconv"
	"strings"
)

func WaitForIt(input string) (sum int) {
	lines := strings.Split(input, "\n")
	times := parseLine(lines[0])
	distances := parseLine(lines[1])

	for i, time := range times {
		distance := distances[i]
		leftDone := false
		rightDone := false
		l := time / 2
		r := l + 1
		winsCount := 0

		for l >= 0 && r <= time {
			if leftDone && rightDone {
				break
			}

			if !leftDone && l*(time-l) > distance {
				winsCount += 1
				l--
			} else {
				leftDone = true
			}

			if !rightDone && r*(time-r) > distance {
				winsCount += 1
				r++
			} else {
				rightDone = true
			}
		}

		if winsCount != 0 {
			if sum != 0 {
				sum *= winsCount
			} else {
				sum = winsCount
			}
		}

	}
	return
}

func WaitForItP2(input string) (sum int) {
	lines := strings.Split(input, "\n")
	time := parseLineP2(lines[0])
	distance := parseLineP2(lines[1])

	leftDone := false
	rightDone := false
	l := time / 2
	r := l + 1

	for l >= 0 && r <= time {
		if leftDone && rightDone {
			break
		}

		if !leftDone && l*(time-l) > distance {
			sum += 1
			l--
		} else {
			leftDone = true
		}

		if !rightDone && r*(time-r) > distance {
			sum += 1
			r++
		} else {
			rightDone = true
		}
	}

	return
}

func parseLineP2(input string) (output int) {
	inputSlice := strings.Split(input, " ")
	res := ""

	for _, elem := range inputSlice {
		if elem != "" {
			_, err := strconv.Atoi(elem)

			if err != nil {
				continue
			}

			res += elem
		}
	}

	numVal, err := strconv.Atoi(res)

	if err != nil {
		panic(err)
	}

	output = numVal

	return
}

func parseLine(input string) (output []int) {
	inputSlice := strings.Split(input, " ")

	for _, elem := range inputSlice {
		if elem != "" {
			numVal, err := strconv.Atoi(elem)

			if err != nil {
				continue
			}

			output = append(output, numVal)
		}
	}
	return
}
