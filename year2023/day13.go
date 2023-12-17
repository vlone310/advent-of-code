package year2023

import (
	"advent-of-code/helpers"
	"regexp"
	"strings"
)

func PointofIncidence(input string) (sum int) {

	matrixes := splitByEmptyNewline(input)

	for _, matrix := range matrixes {
		split := strings.Split(matrix, "\n")
		rotated := rotate(split)

		rR := 0
		cR := 0

		// perform horizontal comparison
		for i := 0; i < len(split)-1; i++ {
			if split[i] == split[i+1] {
				l := i - 1
				r := i + 2

				for !helpers.IsOutOfBound(l, len(split)) && !helpers.IsOutOfBound(r, len(split)) && split[l] == split[r] {
					l--
					r++
				}

				if helpers.IsOutOfBound(l, len(split)) || helpers.IsOutOfBound(r, len(split)) {
					rR = (i + 1) * 100
					break
				}
			}
		}

		// Perform vertical comparison
		for i := 0; i < len(rotated)-1; i++ {
			if rotated[i] == rotated[i+1] {
				l := i - 1
				r := i + 2

				for !helpers.IsOutOfBound(l, len(rotated)) && !helpers.IsOutOfBound(r, len(rotated)) && rotated[l] == rotated[r] {
					l--
					r++
				}

				if helpers.IsOutOfBound(l, len(rotated)) || helpers.IsOutOfBound(r, len(rotated)) {
					cR = i + 1
					break
				}

			}
		}

		sum += rR
		sum += cR
	}

	return
}

func PointofIncidenceP2(input string) (sum int) {

	matrixes := splitByEmptyNewline(input)

	for _, matrix := range matrixes {
		split := strings.Split(matrix, "\n")
		rotated := rotate(split)

		rR := 0
		cR := 0

		// perform horizontal comparison
		for i := 0; i < len(split)-1; i++ {
			if split[i] == split[i+1] || canFixSmudge(split[i], split[i+1]) {
				l := i - 1
				r := i + 2
				isSmudgeFixed := canFixSmudge(split[i], split[i+1])

				for !helpers.IsOutOfBound(l, len(split)) && !helpers.IsOutOfBound(r, len(split)) && (split[l] == split[r] || (!isSmudgeFixed && canFixSmudge(split[l], split[r]))) {
					if canFixSmudge(split[l], split[r]) {
						isSmudgeFixed = true
					}
					l--
					r++
				}

				if (helpers.IsOutOfBound(l, len(split)) || helpers.IsOutOfBound(r, len(split))) && isSmudgeFixed {
					rR = (i + 1) * 100
					break
				}
			}
		}

		// Perform vertical comparison
		for i := 0; i < len(rotated)-1; i++ {
			if rotated[i] == rotated[i+1] || canFixSmudge(rotated[i], rotated[i+1]) {
				l := i - 1
				r := i + 2
				isSmudgeFixed := canFixSmudge(rotated[i], rotated[i+1])

				for !helpers.IsOutOfBound(l, len(rotated)) && !helpers.IsOutOfBound(r, len(rotated)) && (rotated[l] == rotated[r] || (!isSmudgeFixed && canFixSmudge(rotated[l], rotated[r]))) {
					if canFixSmudge(rotated[l], rotated[r]) {
						isSmudgeFixed = true
					}
					l--
					r++
				}

				if (helpers.IsOutOfBound(l, len(rotated)) || helpers.IsOutOfBound(r, len(rotated))) && isSmudgeFixed {
					cR = i + 1
					break
				}

			}
		}

		sum += rR
		sum += cR
	}

	return
}

func canFixSmudge(line1 string, line2 string) bool {
	differences := 0

	for i, char := range line1 {
		if char != rune(line2[i]) {
			differences++
		}
	}

	return differences == 1
}

func splitByEmptyNewline(str string) []string {
	strNormalized := regexp.
		MustCompile("\r\n").
		ReplaceAllString(str, "\n")

	return regexp.
		MustCompile(`\n\s*\n`).
		Split(strNormalized, -1)

}

func rotate(input []string) []string {
	res := make([]string, len(input[0]))

	for i := len(input) - 1; i >= 0; i-- {
		for j := 0; j < len(input[i]); j++ {
			res[j] += string(input[i][j])
		}
	}

	return res
}
