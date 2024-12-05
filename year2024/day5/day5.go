package year2024

import (
	"strconv"
	"strings"
)

func printQueue(input string) (res int) {
	rules, updates := parseInput(input)

	for _, updateRowStr := range updates {
		if updateRowStr == "" {
			continue
		}

		updateRow := strings.Split(updateRowStr, ",")
		visited := map[string]int{}
		isViolation := false

		for i, update := range updateRow {
			if exists, _ := isViolationExist(rules, update, visited); exists {
				isViolation = true
				break
			}
			visited[update] = i
		}

		if !isViolation {
			middleChar, _ := strconv.Atoi(updateRow[len(updateRow)/2])
			res += middleChar
		}
	}

	return
}

func printQueueP2(input string) (res int) {
	rules, updates := parseInput(input)

	for _, updateRowStr := range updates {
		if updateRowStr == "" {
			continue
		}

		updateRow := strings.Split(updateRowStr, ",")
		res += resolveViolationRow(rules, updateRow)

	}

	return
}

func parseInput(input string) (rules []string, updates []string) {
	rawInpt := strings.Split(input, "\n\n")
	rules = strings.Split(rawInpt[0], "\n")   // [47|53 79|28 ...]
	updates = strings.Split(rawInpt[1], "\n") // [47,59,16 22,33,44 ...]

	return
}

func resolveViolationRow(rules []string, row []string) (res int) {
	newRow := make([]string, len(row))
	copy(newRow, row)
	visited := map[string]int{}
	isViolation := false
	violationResolved := false

	for !violationResolved {
		for i, char := range newRow {
			if exists, rule := isViolationExist(rules, char, visited); exists {
				prevChar := rule[1]
				prevCharIdx := visited[rule[1]]
				newRow[i] = prevChar
				newRow[prevCharIdx] = char
				isViolation = true
				visited = map[string]int{}
				break
			} else {
				visited[char] = i
			}

			if i == len(newRow)-1 {
				violationResolved = true
			}
		}
	}

	if isViolation {
		middleChar, _ := strconv.Atoi(newRow[len(newRow)/2])
		res += middleChar
	}

	return
}

func isViolationExist(rules []string, target string, visited map[string]int) (bool, []string) {

	for _, ruleStr := range rules {
		rulesSplit := strings.Split(ruleStr, "|")
		_, exists := visited[rulesSplit[1]]
		if rulesSplit[0] == target && exists {
			return true, rulesSplit
		}
	}
	return false, []string{}
}
