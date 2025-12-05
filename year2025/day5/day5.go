package year2025

import (
	"sort"
	"strconv"
	"strings"
)

func Cafeteria(input string) (res int) {
	ing := parseIngredients(input)

	for _, id := range ing.ids {
		for _, rng := range ing.ranges {
			if id >= rng.from && id <= rng.to {
				res += 1
				break
			}
		}
	}

	return res
}

func CafeteriaP2(input string) (res int) {
	ing := parseIngredients(input)
	overlapping := calculateOverlapingRanges(ing.ranges)

	for _, rng := range overlapping {
		res += rng.to - rng.from + 1
	}

	return res
}

type Range struct {
	from int
	to   int
}

type Ingredients struct {
	ids    []int
	ranges []Range
}

func parseIngredients(input string) Ingredients {
	parts := strings.Split(input, "\n\n")
	rangesStr := strings.Split(parts[0], "\n")
	idsStr := strings.Split(parts[1], "\n")
	ranges := make([]Range, 0, len(rangesStr))
	ids := make([]int, 0, len(idsStr))

	for _, rangeStr := range rangesStr {
		sliced := strings.Split(rangeStr, "-")
		from, _ := strconv.Atoi(sliced[0])
		to, _ := strconv.Atoi(sliced[1])
		ranges = append(ranges, Range{from, to})
	}

	for _, idStr := range idsStr {
		i, _ := strconv.Atoi(idStr)
		ids = append(ids, i)
	}

	return Ingredients{
		ids,
		ranges,
	}
}

func calculateOverlapingRanges(ranges []Range) []Range {
	if len(ranges) <= 1 {
		return ranges
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].from < ranges[j].from
	})

	merged := make([]Range, 0)
	merged = append(merged, ranges[0])

	for i := 1; i < len(ranges); i++ {
		lastMerged := &merged[len(merged)-1]
		current := ranges[i]

		if current.from <= lastMerged.to {
			if current.to > lastMerged.to {
				lastMerged.to = current.to
			}
		} else {
			merged = append(merged, current)
		}
	}

	return merged
}
