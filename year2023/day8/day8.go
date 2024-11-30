package year2023

import (
	"strings"

	"github.com/vlone310/advent-of-code/helpers"
)

func HauntedWasteland(input string) (steps int) {
	inputSplit := strings.Split(input, "\n")

	instructions := inputSplit[0]
	networkNodesMap := map[string][]string{}

	for _, networkNodeStr := range inputSplit[2:] {
		nodesSplit := strings.Split(networkNodeStr, " = ")
		mainNode := nodesSplit[0]
		nodes := helpers.SplitByMultipleSeparators(nodesSplit[1][1:len(nodesSplit[1])-1], `, `)
		networkNodesMap[mainNode] = nodes
	}

	currNode := "AAA"
	currInstructionIdx := 0

	for currNode != "ZZZ" {
		instruction := instructions[currInstructionIdx]

		nodes := networkNodesMap[currNode]

		if instruction == 'L' {
			currNode = nodes[0]
		} else {
			currNode = nodes[1]
		}

		if currInstructionIdx == len(instructions)-1 {
			currInstructionIdx = 0
		} else {
			currInstructionIdx++
		}

		steps++
	}

	return steps
}

func HauntedWastelandP2(input string) int {
	inputSplit := strings.Split(input, "\n")

	instructions := inputSplit[0]
	networkNodesMap := map[string][]string{}
	startingNodes := []string{}

	for _, networkNodeStr := range inputSplit[2:] {
		nodesSplit := strings.Split(networkNodeStr, " = ")
		mainNode := nodesSplit[0]
		nodes := helpers.SplitByMultipleSeparators(nodesSplit[1][1:len(nodesSplit[1])-1], `, `)
		if mainNode[len(mainNode)-1] == 'A' {
			startingNodes = append(startingNodes, mainNode)
		}
		networkNodesMap[mainNode] = nodes
	}

	answer := 1

	for _, startingNode := range startingNodes {
		curSteps := 0
		curNode := startingNode
		currInstructionIndex := 0

		for curNode[len(curNode)-1] != 'Z' {
			instruction := instructions[currInstructionIndex]

			nodes := networkNodesMap[curNode]

			if instruction == 'L' {
				curNode = nodes[0]
			} else {
				curNode = nodes[1]
			}

			if currInstructionIndex == len(instructions)-1 {
				currInstructionIndex = 0
			} else {
				currInstructionIndex++
			}

			curSteps++
		}

		answer = (answer * curSteps) / gcd(answer, curSteps)
	}

	return answer
}

func gcd(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
