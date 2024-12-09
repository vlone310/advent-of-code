package year2024

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func diskFragmenter(input string) (res int) {
	layout := []string{}

	// compute layout of files and free spaces (.)
	for i, instruction := range input {
		if byte(instruction) == '\n' {
			continue
		}

		fileID := strconv.Itoa(i / 2)
		isFile := i%2 == 0

		if isFile {
			layout = append(layout, slices.Repeat([]string{fileID}, int(instruction-'0'))...)
		} else {
			layout = append(layout, slices.Repeat([]string{"."}, int(instruction-'0'))...)
		}
	}

	// using two pointers to rearrange files in disk
	i := 0
	j := len(layout) - 1

	for i < j {
		if layout[i] != "." {
			i++
			continue
		}

		if layout[j] == "." {
			j--
			continue
		}

		layout[i], layout[j] = layout[j], layout[i]
	}

	// calculate the result
	idx := 0
	for idx < len(layout) && layout[idx] != "." {
		num, _ := strconv.Atoi(layout[idx])
		res += num * idx
		idx++
	}

	return
}

func diskFragmenterP2(input string) (res int) {
	layout := []string{}

	// compute layout of files and free spaces (.)
	for i, instruction := range input {
		if byte(instruction) == '\n' || byte(instruction) == '0' {
			continue
		}

		fileID := i / 2
		isFile := i%2 == 0

		if isFile {
			layout = append(layout, fmt.Sprintf("%c_%v", instruction, fileID))
		} else {
			layout = append(layout, fmt.Sprintf("%c_%v", instruction, "."))
		}
	}

	// using two pointers to rearrange files in disk
	i := 0
	j := len(layout) - 1
	for j > 0 {
		instructionI := strings.Split(layout[i], "_")
		instructionJ := strings.Split(layout[j], "_")

		if i == j {
			i = 0
			j--
			continue
		}

		if instructionI[1] != "." {
			i++
			continue
		}

		if instructionJ[1] == "." {
			j--
			continue
		}

		emptySlots, _ := strconv.Atoi(instructionI[0])
		filesLen, _ := strconv.Atoi(instructionJ[0])

		if emptySlots < filesLen {
			i++
			continue
		}

		tempLayout := []string{}
		tempLayout = append(tempLayout, layout[:i]...)
		tempLayout = append(tempLayout, layout[j])
		if emptySlots-filesLen > 0 {
			tempLayout = append(tempLayout, fmt.Sprintf("%v_%v", emptySlots-filesLen, "."))
		}
		tempLayout = append(tempLayout, layout[i+1:j]...)
		tempLayout = append(tempLayout, fmt.Sprintf("%v_%v", filesLen, "."))
		tempLayout = append(tempLayout, layout[j+1:]...)
		layout = tempLayout
		i = 0
	}

	// calculate the result
	curIdx := 0
	for _, instruciton := range layout {
		instructionSplit := strings.Split(instruciton, "_")
		instructionCount, _ := strconv.Atoi(instructionSplit[0])
		if instructionSplit[1] == "." {
			curIdx += instructionCount
			continue
		}

		instructionId, _ := strconv.Atoi(instructionSplit[1])
		for i := 0; i < instructionCount; i++ {
			res += (curIdx + i) * instructionId
		}
		curIdx += instructionCount
	}

	return
}
