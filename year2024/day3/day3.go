package year2024

import (
	"regexp"
	"strconv"
	"strings"
)

var doFuncName = "do"
var dontFuncName = "don't"

var mulPattern = `mul\(\d+,\s*\d+\)`
var numPattern = `\d+`
var mulDoDontPattern = `do\(\)|don\'t\(\)|mul\(\d+,\s*\d+\)`

func MullItOver(input string) (res int) {

	re := regexp.MustCompile(mulPattern)
	reNums := regexp.MustCompile(numPattern)

	matches := re.FindAllString(input, -1)

	for _, match := range matches {
		nums := reNums.FindAllString(match, -1)
		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])
		res += num1 * num2
	}

	return res
}

func MullItOverP2(input string) (res int) {

	re := regexp.MustCompile(mulDoDontPattern)
	reNums := regexp.MustCompile(numPattern)
	shouldEvaluate := true

	matches := re.FindAllString(input, -1)

	for _, match := range matches {
		if strings.HasPrefix(match, dontFuncName) {
			shouldEvaluate = false
			continue
		}

		if strings.HasPrefix(match, doFuncName) {
			shouldEvaluate = true
			continue
		}

		if shouldEvaluate {
			nums := reNums.FindAllString(match, -1)
			num1, _ := strconv.Atoi(nums[0])
			num2, _ := strconv.Atoi(nums[1])
			res += num1 * num2
		}
	}

	return res
}
