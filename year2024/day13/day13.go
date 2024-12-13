package year2024

import (
	"regexp"
	"strings"

	"github.com/vlone310/advent-of-code/helpers"
)

type coords struct {
	x, y int
}

type rule struct {
	buttonA coords
	buttonB coords
	target  coords
}

func clawContraption(input string) (res int) {
	rules := parseInput(input, 0)

	for _, rule := range rules {
		x, y := findXY(rule)

		if x > 100 || y > 100 || float64(int(x)) != x || float64(int(y)) != y || x < 0 || y < 0 {
			continue
		}

		res += int(x*3 + y*1)
	}
	return res
}

func clawContraptionP2(input string) (res int) {
	rules := parseInput(input, 10000000000000)

	for _, rule := range rules {
		x, y := findXY(rule)

		if float64(int(x)) != x || float64(int(y)) != y || x < 0 || y < 0 {
			continue
		}

		res += int(x*3 + y*1)
	}

	return res
}

func findXY(rule rule) (float64, float64) {
	a1, b1, c1 := float64(rule.buttonA.x), float64(rule.buttonB.x), float64(rule.target.x)
	a2, b2, c2 := float64(rule.buttonA.y), float64(rule.buttonB.y), float64(rule.target.y)

	new_b1 := a2 * b1
	new_c1 := a2 * c1

	new_b2 := a1 * b2
	new_c2 := a1 * c2

	coeff_y := new_b2 - new_b1
	constant := new_c2 - new_c1

	Y := constant / coeff_y

	X := (c1 - b1*Y) / a1

	return X, Y
}

func parseInput(input string, targetShift int) []rule {
	rawRules := strings.Split(input, "\n\n")
	rules := make([]rule, len(rawRules))

	for i, rawRule := range rawRules {
		re := regexp.MustCompile(`Button A: X[+=](\d+), Y[+=](\d+)\nButton B: X[+=](\d+), Y[+=](\d+)\nPrize: X[=](\d+), Y[=](\d+)`)

		matches := re.FindStringSubmatch(rawRule)

		if len(matches) != 7 {
			continue
		}

		rules[i] = rule{
			buttonA: coords{helpers.ParseInt(matches[1]), helpers.ParseInt(matches[2])},
			buttonB: coords{helpers.ParseInt(matches[3]), helpers.ParseInt(matches[4])},
			target:  coords{helpers.ParseInt(matches[5]) + targetShift, helpers.ParseInt(matches[6]) + targetShift},
		}

	}

	return rules
}
