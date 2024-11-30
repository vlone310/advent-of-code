package year2023

import (
	"slices"
	"strconv"
	"strings"

	"github.com/vlone310/advent-of-code/helpers"
)

type Lens struct {
	label       string
	focalLength int
}

func LensLibrary(input string) (sum int) {
	steps := strings.Split(input, ",")

	for _, step := range steps {
		stepSum := 0
		for _, el := range step {
			ascii := int(el)

			stepSum = (stepSum + ascii) * 17 % 256

		}
		sum += stepSum
	}

	return
}

func LensLibraryP2(input string) (sum int) {
	HASHMAP := map[int][]Lens{}

	for _, instruction := range strings.Split(input, ",") {
		instructionSplit := helpers.SplitByMultipleSeparators(instruction, `=|-`)
		isInsertOperation := strings.Contains(instruction, "=")
		label := instructionSplit[0]
		focalLengthStr := instructionSplit[1]
		hash := LensLibrary(label)

		// Performing delete operation
		if !isInsertOperation {
			for i, lens := range HASHMAP[hash] {
				if lens.label == label {
					HASHMAP[hash] = append(HASHMAP[hash][:i], HASHMAP[hash][i+1:]...)
					break
				}
			}
			// Performing insert operation
		} else {
			focalLength, err := strconv.Atoi(focalLengthStr)
			if err != nil {
				panic(err)
			}

			indexOfLabel := slices.IndexFunc(HASHMAP[hash], func(lens Lens) bool {
				return lens.label == label
			})

			// Check if lens with same label exists, if yes, replace it with new one
			if indexOfLabel != -1 {
				HASHMAP[hash][indexOfLabel] = Lens{label, focalLength}
				// If lens with same label doesn't exist, append new one to the end of slice
			} else {
				HASHMAP[hash] = append(HASHMAP[hash], Lens{label, focalLength})

			}

		}
	}

	for key, lenses := range HASHMAP {
		for i, lens := range lenses {
			sum += (key + 1) * (i + 1) * lens.focalLength
		}
	}

	return
}
