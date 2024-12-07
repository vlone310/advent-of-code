package year2024

import (
	"strconv"
	"strings"

	"github.com/vlone310/advent-of-code/helpers"
)

func bridgeRepair(input string) (res int) {
	defer helpers.Timer()()
	equations := strings.Split(input, "\n")

	for _, equation := range equations {
		if len(equation) == 0 {
			continue
		}

		instructions := strings.Split(equation, ": ")
		result, _ := strconv.Atoi(instructions[0])
		nums := parseNumbers(instructions[1])

		if isValidEquation(result, nums[0], nums[1:], false) {
			res += result
		}
	}

	return
}

func bridgeRepairP2(input string) (res int) {
	defer helpers.Timer()()
	equations := strings.Split(input, "\n")

	for _, equation := range equations {
		if len(equation) == 0 {
			continue
		}

		instructions := strings.Split(equation, ": ")
		result, _ := strconv.Atoi(instructions[0])
		nums := parseNumbers(instructions[1])

		if isValidEquation(result, nums[0], nums[1:], true) {
			res += result
		}
	}

	return
}

func isValidEquation(target, cur int, nums []int, withConcat bool) bool {
	if len(nums) == 0 && target == cur {
		return true
	}

	if len(nums) == 0 {
		return false
	}

	return isValidEquation(target, cur*nums[0], nums[1:], withConcat) ||
		isValidEquation(target, cur+nums[0], nums[1:], withConcat) ||
		(withConcat && isValidEquation(target, concatInts(cur, nums[0]), nums[1:], withConcat))
}

func concatInts(a, b int) int {
	bLen := 1
	for tmp := b; tmp > 0; tmp /= 10 {
		bLen *= 10
	}
	return a*bLen + b
}

func parseNumbers(input string) []int {
	nums := strings.Split(input, " ")
	res := make([]int, len(nums))

	for i, el := range nums {
		num, _ := strconv.Atoi(el)
		res[i] = num
	}

	return res
}
