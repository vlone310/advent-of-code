package year2023

import (
	"cmp"
	"math"
	"slices"
	"strconv"
	"strings"
)

type HandEquity struct {
	combination int
	hand        string
	bet         int
}

var CardToEquity map[string]int = map[string]int{
	"2": 1,
	"3": 2,
	"4": 3,
	"5": 4,
	"6": 5,
	"7": 6,
	"8": 7,
	"9": 8,
	"T": 9,
	"J": 10,
	"Q": 11,
	"K": 12,
	"A": 13,
}

var CardToEquityP2 map[string]int = map[string]int{
	"J": 0,
	"2": 1,
	"3": 2,
	"4": 3,
	"5": 4,
	"6": 5,
	"7": 6,
	"8": 7,
	"9": 8,
	"T": 9,
	"Q": 10,
	"K": 11,
	"A": 12,
}

const (
	HighCard int = iota
	OnePair
	TwoPairs
	ThreeOfKind
	FullHouse
	FourOfKind
	FiveOfKind
)

// I suppose it's possible to come up with better approach :)
func CamelCards(input string) (result int) {
	lines := strings.Split(input, "\n")
	parsedHands := []HandEquity{}

	for _, line := range lines {
		lineSplit := strings.Split(line, " ")
		hand := lineSplit[0]
		bet, err := strconv.Atoi(lineSplit[1])
		handMap := map[string]int{}

		if err != nil {
			panic(err)
		}

		for _, card := range hand {
			handMap[string(card)] += 1
		}

		parsedHands = append(parsedHands, HandEquity{
			combination: getCombination(handMap),
			hand:        hand,
			bet:         bet,
		})

	}

	slices.SortFunc(parsedHands, func(a, b HandEquity) int {
		if n := cmp.Compare(a.combination, b.combination); n != 0 {
			return n
		}
		for idx, card := range a.hand {
			eqA := CardToEquity[string(card)]
			eqB := CardToEquity[string(b.hand[idx])]

			if n := cmp.Compare(eqA, eqB); n != 0 {
				return n
			}
		}
		return 1
	})

	for i, hand := range parsedHands {
		result += hand.bet * (i + 1)
	}

	return result
}

func CamelCardsP2(input string) (result int) {
	lines := strings.Split(input, "\n")
	parsedHands := []HandEquity{}

	for _, line := range lines {
		lineSplit := strings.Split(line, " ")
		hand := lineSplit[0]
		bet, err := strconv.Atoi(lineSplit[1])
		handMap := map[string]int{}
		jokersCount := 0

		if err != nil {
			panic(err)
		}

		for _, card := range hand {
			if card == 'J' {
				jokersCount++
			} else {
				handMap[string(card)] += 1
			}
		}
		currMaxVal := 0
		maxKey := ""
		for key, val := range handMap {
			if val > currMaxVal {
				currMaxVal = val
				maxKey = key
			}
		}

		handMap[maxKey] += jokersCount

		parsedHands = append(parsedHands, HandEquity{
			combination: getCombination(handMap),
			hand:        hand,
			bet:         bet,
		})

	}

	slices.SortFunc(parsedHands, func(a, b HandEquity) int {
		if n := cmp.Compare(a.combination, b.combination); n != 0 {
			return n
		}
		for idx, card := range a.hand {
			eqA := CardToEquityP2[string(card)]
			eqB := CardToEquityP2[string(b.hand[idx])]

			if n := cmp.Compare(eqA, eqB); n != 0 {
				return n
			}
		}
		return 1
	})

	for i, hand := range parsedHands {
		result += hand.bet * (i + 1)
	}

	return result
}

// TODO: Refactor
func getCombination(input map[string]int) int {
	currComb := HighCard
	for _, value := range input {
		newComb := HighCard
		if value == 5 {
			newComb = FiveOfKind
		} else if value == 4 {
			newComb = FourOfKind
		} else if (value == 3 && currComb == OnePair) || (value == 2 && currComb == ThreeOfKind) {
			newComb = FullHouse
		} else if value == 3 {
			newComb = ThreeOfKind
		} else if value == 2 && currComb == OnePair {
			newComb = TwoPairs
		} else if value == 2 {
			newComb = OnePair
		} else {
			newComb = HighCard
		}

		currComb = int(math.Max(float64(newComb), float64(currComb)))
	}

	return currComb
}
