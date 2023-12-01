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

func TrebuchetP2(input string) float64 {
	stringToNumMap := map[string]string{
		"one": "1",
		"two": "2",
		"three": "3",
		"four": "4",
		"five": "5",
		"six": "6",
		"seven": "7",
		"eight": "8",
		"nine": "9",
	}

	var leftNum string
	var rightNum string
	var sum float64

	inputSlice := strings.Split(input, "\n")

	for _, v := range inputSlice {
		left := 0
		right := len(v) - 1

		for left <= right {
			if leftNum != "" && rightNum != "" {
				break
			}

			if leftNum == "" {
				_, err := strconv.ParseFloat(string(v[left]), 64)
				if err == nil {
					leftNum = string(v[left])
				} else {
					for k, val := range stringToNumMap {
						if strings.Contains(v[:left+1], k) {
							leftNum = val
							break
						}
					}
				}
			}

			if rightNum == "" {
				_, err := strconv.ParseFloat(string(v[right]), 64)
				if err == nil {
					rightNum = string(v[right])
				} else {
					for k, val := range stringToNumMap {
						if strings.Contains(v[right:], k) {
							rightNum = val
							break
						}
					}
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