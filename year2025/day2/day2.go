package year2025

import (
	"strconv"
	"strings"
)

func GiftShop(input string) (res int) {
	prep := strings.Split(input, "\n")[0]
	idsRange := strings.Split(prep, ",")

	for _, idRange := range idsRange {
		ids := strings.Split(idRange, "-")
		start, _ := strconv.Atoi(ids[0])
		end, _ := strconv.Atoi(ids[1])

		for start <= end {

			sId := strconv.Itoa(start)
			idLen := len(sId)

			if sId[0:idLen/2] == sId[idLen/2:idLen] {
				res += start
			}

			start++
		}
	}
	return res
}

func GiftShopP2(input string) (res int) {
	prep := strings.Split(input, "\n")[0]
	idsRange := strings.Split(prep, ",")

	for _, idRange := range idsRange {
		ids := strings.Split(idRange, "-")
		start, _ := strconv.Atoi(ids[0])
		end, _ := strconv.Atoi(ids[1])

		for start <= end {

			sId := strconv.Itoa(start)
			if IsInvalidID(sId) {
				res += start
			}
			start++
		}
	}
	return res
}

func IsInvalidID(id string) bool {
	n := len(id)

	if n < 2 {
		return false
	}

	for l := 1; l <= n/2; l++ {
		if n%l != 0 {
			continue
		}

		pattern := id[:l]

		repeatCount := n / l

		repeatedString := strings.Repeat(pattern, repeatCount)

		if id == repeatedString {
			return true
		}
	}

	return false
}
