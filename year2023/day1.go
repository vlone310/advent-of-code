package year2023

import (
	"advent-of-code/helpers"
	"strconv"
	"strings"
)

func Trebuchet(input string) float64 {
	var sum float64

	inputSlice := strings.Split(input, "\n")

	for _, v := range inputSlice {
		left := 0
		right := len(v) - 1
		var leftNum byte
		var rightNum byte

		for left <= right {
			if leftNum != 0 && rightNum != 0 {
				break
			}

			if leftNum == 0 {
				if helpers.IsDigit(v[left]) {
					leftNum = v[left]
				}
			}

			if rightNum == 0 {
				if helpers.IsDigit(v[right]) {
					rightNum = v[right]
				}
			}
			if leftNum == 0 {
				left++
			}
			if rightNum == 0 {
				right--
			}
		}

		s := string(leftNum) + string(rightNum)
		val, err := strconv.ParseFloat(s, 64)
		if err != nil {
			continue
		}

		sum += val
	}

	return sum
}

func replaceNumWordsWithNums(input string) string {
	stringToNumMap := map[string]string{
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4r",
		"five":  "f5e",
		"six":   "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine":  "n9e",
	}

	str := input

	for k, v := range stringToNumMap {
		str = strings.ReplaceAll(str, k, v)
	}

	return str
}

func TrebuchetP2(input string) float64 {

	var sum float64

	inputSlice := strings.Split(input, "\n")

	for _, v := range inputSlice {
		var leftNum byte
		var rightNum byte

		str := replaceNumWordsWithNums(v)

		left := 0
		right := len(str) - 1

		for left <= right {
			if leftNum != 0 && rightNum != 0 {
				break
			}

			if leftNum == 0 {
				if helpers.IsDigit(str[left]) {
					leftNum = str[left]
				}
			}

			if rightNum == 0 {
				if helpers.IsDigit(str[right]) {
					rightNum = str[right]
				}
			}

			if leftNum == 0 {
				left++
			}

			if rightNum == 0 {
				right--
			}
		}

		s := string(leftNum) + string(rightNum)
		val, err := strconv.ParseFloat(s, 64)
		if err != nil {
			continue
		}

		sum += val
	}

	return sum
}

func Main() {
	Trebuchet("123")
}
