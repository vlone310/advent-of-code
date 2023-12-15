package year2023

import (
	"advent-of-code/helpers"
	"strconv"
	"strings"
)

func Scratchcards(input string) (sum int64) {
	cardsContents := strings.Split(input, "\n")

	for _, cardContent := range cardsContents {
		cardNums := strings.Split(strings.Split(cardContent, ": ")[1], " | ")
		myNums := cardNums[0]
		winningNums := cardNums[1]
		myNumsMap := make(map[int64]bool)
		roundCount := 0

		for _, num := range strings.Split(myNums, " ") {
			numInt, err := strconv.ParseInt(num, 10, 64)

			if err == nil {
				myNumsMap[numInt] = true
			}
		}

		for _, num := range strings.Split(winningNums, " ") {
			numInt, err := strconv.ParseInt(num, 10, 64)

			if err == nil {
				if _, ok := myNumsMap[numInt]; ok {
					if roundCount == 0 {
						roundCount = 1
					} else {
						roundCount *= 2
					}
				}
			}
		}

		sum += int64(roundCount)
	}

	return sum
}

func ScratchcardsP2(input string) (instancesCount int) {
	cardsContents := strings.Split(input, "\n")
	cardInstancesMap := make(map[int64]int)

	for _, cardContent := range cardsContents {
		// parsing card content
		content := strings.Split(cardContent, ": ")
		cardNums := strings.Split(content[1], " | ")
		myNums := cardNums[0]
		winningNums := cardNums[1]
		myNumsMap := make(map[int64]bool)
		copiesWon := 0
		cardId, err := strconv.ParseInt(strings.Split(content[0], " ")[len(strings.Split(content[0], " "))-1], 10, 64)

		if err != nil {
			panic(err)
		}

		// writing my nums to the map
		for _, num := range strings.Split(myNums, " ") {
			numInt, err := strconv.ParseInt(num, 10, 64)

			if err == nil {
				myNumsMap[numInt] = true
			}
		}

		// counting how much copies we won
		for _, num := range strings.Split(winningNums, " ") {
			numInt, err := strconv.ParseInt(num, 10, 64)

			if err == nil {
				if _, ok := myNumsMap[numInt]; ok {
					copiesWon++
				}
			}
		}

		// adding to result sum 1 original card and all copies
		instancesCount += cardInstancesMap[cardId] + 1

		// for each won copy, add to map, 1 * current card instance copies + 1
		for copiesWon > 0 {
			if !helpers.IsOutOfBound(int(cardId)+copiesWon-1, len(cardsContents)) {
				cardInstancesMap[cardId+int64(copiesWon)] += 1 * (cardInstancesMap[cardId] + 1)
			}
			copiesWon--
		}
	}

	return instancesCount
}
