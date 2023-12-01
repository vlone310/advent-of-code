package year2023

import (
	"strconv"
	"strings"
)

func Trebuchet(input string) float64 {
	var sum float64
	var leftNum string
	var rightNum string

	inputSlice := strings.Split(input, "\n")

	for _, v := range inputSlice {
		left := 0
		right := len(v) - 1
		for left <= right {
			if leftNum != "" && rightNum != "" {
				break
			}

			if (leftNum == "") {
				_, err := strconv.ParseFloat(string (v[left]), 64)
				if err == nil {
					leftNum = string(v[left])
				}
			}

			if (rightNum == "") {
				_, err := strconv.ParseFloat(string (v[right]), 64)
				if err == nil {
					rightNum = string(v[right])
				}
			}
			if (leftNum == "") {
				left++
			}
			if (rightNum == "") {
				right--
			}
		}

		s := leftNum + rightNum
		val, err := strconv.ParseFloat(s, 64)
		if err != nil {
			continue
		}
		sum += val
		leftNum = ""
		rightNum = ""
	}

	return sum
}

func replaceNumWordsWithNums (input string) string {
	stringToNumMap := map[string]string{
		"one": "o1e",
		"two": "t2o",
		"three": "t3e",
		"four": "f4r",
		"five": "f5e",
		"six": "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine": "n9e",
	}

	str := input;
	for k, v := range stringToNumMap {
		str = strings.ReplaceAll(str, k, v)
	}

	return str;
}

func TrebuchetP2(input string) float64 {

	var leftNum string
	var rightNum string
	var sum float64

	inputSlice := strings.Split(input, "\n")

	for _, v := range inputSlice {

		str := replaceNumWordsWithNums(v)

		left := 0
		right := len(str) - 1

		for left <= right {
			if leftNum != "" && rightNum != "" {
				break
			}

			if leftNum == "" {
				_, err := strconv.ParseFloat(string(str[left]), 64)
				if err == nil {
					leftNum = string(str[left])
				}
			}

			if rightNum == "" {
				_, err := strconv.ParseFloat(string(str[right]), 64)
				if err == nil {
					rightNum = string(str[right])
				}
			}

			if (leftNum == "") {
				left++
			}

			if (rightNum == "") {
				right--
			}
		}

		s := leftNum + rightNum
		val, err := strconv.ParseFloat(s, 64)
		if err != nil {
			continue
		}

		sum += val
		leftNum = ""
		rightNum = ""
	}

	return sum
}

func Main() {
	Trebuchet("123")
}
