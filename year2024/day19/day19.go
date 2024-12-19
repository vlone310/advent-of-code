package year2024

import (
	"strings"
)

func linenLayout(input string) (res int) {
	patterns, designs, maxPatternLen := parseToLayout(input)

	for _, design := range designs {
		if traverse(patterns, map[string]int{}, design, "", maxPatternLen) > 0 {
			res += 1
		}
	}

	return
}

func linenLayoutP2(input string) (res int) {
	patterns, designs, maxPatternLen := parseToLayout(input)

	for _, design := range designs {
		res += traverse(patterns, map[string]int{}, design, "", maxPatternLen)
	}

	return
}

func traverse(patterns map[string]bool, cache map[string]int, design, slice string, maxPatternLen int) int {
	if slice == design {
		return 1
	}

	if val, exists := cache[slice]; exists {
		return val
	}

	res := 0

	for i := maxPatternLen; i > 0; i-- {
		if len(slice)+i > len(design) {
			continue
		}

		nextPattern := design[len(slice) : len(slice)+i]

		if patterns[nextPattern] {
			variants := traverse(patterns, cache, design, slice+nextPattern, maxPatternLen)
			res += variants
			cache[slice] = res
		}
	}

	return res
}

func parseToLayout(input string) (patterns map[string]bool, designs []string, maxPatternLen int) {
	split := strings.Split(input, "\n\n")
	patternsList := strings.Split(split[0], ", ")
	designs = strings.Split(split[1], "\n")
	patterns = make(map[string]bool)

	for _, pattern := range patternsList {
		patterns[pattern] = true

		if len(pattern) > maxPatternLen {
			maxPatternLen = len(pattern)
		}
	}

	return
}
