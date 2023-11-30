package year2022

import (
	"strings"
)

func RockPaperScissors(input string) float64 {
	scores := map[string]float64{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}
	opToMyHandMap := map[string]string{
		"A": "X",
		"B": "Y",
		"C": "Z",
	}
	winnings := map[string]string{
		"X": "Z",
		"Y": "X",
		"Z": "Y",
	}

	var scoreSum float64

	roundsList := strings.Split(input, "\n")

	for _, round := range roundsList {
		var hands = strings.Split(round, " ")
		if len(hands) != 2 {
			continue
		}
		opHand := opToMyHandMap[hands[0]]
		myHand := hands[1]

		scoreSum += scores[myHand]

		// handle draw round
		if opHand == myHand {
			scoreSum += 3
			// handle my win
		} else if opHand == winnings[myHand] {
			scoreSum += 6
			// handle my loss
		}
	}

	return scoreSum
}

func RockPaperScissorsP2(input string) float64 {
	scores := map[string]float64{
		"A": 1,
		"B": 2,
		"C": 3,
	}

	wins := map[string]string{
		"A": "C",
		"B": "A",
		"C": "B",
	}

	loses := map[string]string{
		"A": "B",
		"B": "C",
		"C": "A",
	}

	var scoreSum float64

	roundsList := strings.Split(input, "\n")
	
	for _, round := range roundsList {
		var in = strings.Split(round, " ")
		if len(in) != 2 {
			continue
		}
		opHand := in[0]
		expectedRoundResult := in[1]

		// I need to lose
		if expectedRoundResult == "X" {
			myHand := wins[opHand]
			scoreSum += scores[myHand]
			// draw
		} else if expectedRoundResult == "Y" {
			scoreSum += scores[opHand] + 3
			// I win
		} else {
			myHand := loses[opHand]
			scoreSum += scores[myHand] + 6
		}
	}

	return scoreSum
}