package year2024

import (
	"log"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/vlone310/advent-of-code/helpers"
)

type Computer struct {
	instructions    []int
	registerA       int
	registerB       int
	registerC       int
	instructionsPtr int
	out             []int
}

func chronospatialComputer(c Computer) Computer {

	for c.instructionsPtr < len(c.instructions)-1 {
		operand := c.instructions[c.instructionsPtr+1]
		instruction := c.instructions[c.instructionsPtr]

		switch instruction {
		// adv
		case 0:
			c.registerA >>= c.comboOperand(operand)
		// bxl
		case 1:
			c.registerB ^= operand
		// bst
		case 2:
			c.registerB = c.comboOperand(operand) & 7
		// jnz
		case 3:
			if c.registerA != 0 {
				c.instructionsPtr = operand
				continue
			}
		// bxc
		case 4:
			c.registerB ^= c.registerC
		// out
		case 5:
			c.out = append(c.out, c.comboOperand(operand)&7)
		// bdv
		case 6:
			c.registerB = c.registerA >> c.comboOperand(operand)
		// cdv
		case 7:
			c.registerC = c.registerA >> c.comboOperand(operand)
		}

		c.instructionsPtr += 2
	}

	return c
}

func chronospatialComputerP2(input string) (a int64) {
	initialC := NewComputer(input)

	for i := len(initialC.instructions) - 1; i >= 0; i-- {
		a <<= 3

		for !slices.Equal(chronospatialComputer(Computer{
			registerA:       int(a),
			registerB:       0,
			registerC:       0,
			instructionsPtr: 0,
			out:             []int{},
			instructions:    initialC.instructions,
		}).out, initialC.instructions[i:]) {
			a++
		}
	}

	return
}

func (c Computer) comboOperand(operand int) int {
	combo := operand

	switch operand {
	case 4:
		combo = c.registerA
	case 5:
		combo = c.registerB
	case 6:
		combo = c.registerC
	}

	return combo
}

func (c Computer) toString() string {
	res := []string{}

	for _, n := range c.out {
		res = append(res, strconv.Itoa(n))
	}

	return strings.Join(res, ",")
}

func NewComputer(input string) Computer {
	var registerA int
	var registerB int
	var registerC int
	split := strings.Split(input, "\n\n")

	re := regexp.MustCompile(`[-]?\d+`)

	registerMatches := re.FindAllString(split[0], -1)

	if len(registerMatches) < 3 {
		log.Fatal("Failed to parse registers")
	}
	registerA = helpers.ParseInt(registerMatches[0])
	registerB = helpers.ParseInt(registerMatches[1])
	registerC = helpers.ParseInt(registerMatches[2])

	program := make([]int, 0)

	programMatches := re.FindAllString(split[1], -1)

	for _, match := range programMatches {
		program = append(program, helpers.ParseInt(match))
	}

	return Computer{
		registerA:    registerA,
		registerB:    registerB,
		registerC:    registerC,
		instructions: program,
		out:          []int{},
	}
}
