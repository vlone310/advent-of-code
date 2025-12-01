package year2025

import (
	"fmt"
	"strconv"
	"strings"
)

const DIAL_START = 50

func SecretEntrance(input string) (pass int) {
	instructionsSlice := strings.Split(input, "\n")
	dial := DIAL_START

	for _, instruction := range instructionsSlice {
		if len(instruction) == 0 {
			return
		}

		rotation := instruction[0]
		c, _ := strconv.Atoi(instruction[1:])
		rotationCount := c % 100

		if rotation == 'R' {
			dial += rotationCount
		} else {
			dial -= rotationCount
		}

		if dial >= 100 {
			dial -= 100
		} else if dial < 0 {
			dial += 100
		}

		if dial == 0 {
			pass++
		}
	}
	return pass
}

func SecretEntranceP2(input string) (pass int) {
	instructionsSlice := strings.Split(input, "\n")
	dial := DIAL_START

	for _, instruction := range instructionsSlice {
		if len(instruction) == 0 {
			return
		}

		rotation := instruction[0]
		c, _ := strconv.Atoi(instruction[1:])
		rotationCount := c % 100
		passes := c / 100
		prevDial := dial

		if rotation == 'R' {
			dial += rotationCount
		} else {
			dial -= rotationCount
		}

		if dial >= 100 {
			dial -= 100
			passes++
		} else if dial < 0 {
			dial += 100
			if prevDial != 0 {
				passes++
			}
		} else if dial == 0 {
			passes++
		}

		pass += passes

		fmt.Printf("instruction: %s, dial: %d, pass: %d\n", instruction, dial, pass)
	}

	return pass
}
