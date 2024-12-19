package day17

import (
	"fmt"
	"strconv"
	"strings"
)

func Part1(lines []string) int {
	regAStr := strings.Split(lines[0], "A: ")[1]
	regBStr := strings.Split(lines[1], "B: ")[1]
	regCStr := strings.Split(lines[2], "C: ")[1]

	regA, _ := strconv.Atoi(regAStr)
	regB, _ := strconv.Atoi(regBStr)
	regC, _ := strconv.Atoi(regCStr)

	instructionsStrings := strings.Split(strings.Split(lines[4], ": ")[1], ",")
	instructions := make([]int, len(instructionsStrings))
	for i, part := range instructionsStrings {
		instructions[i] = int(part[0] - '0')
	}

	output := make([]rune, 0)

	for i := 0; i < len(instructions)-1; i += 2 {
		instruction := instructions[i]
		literal := instructions[i+1]
		var combo int
		if literal == 4 {
			combo = regA
		} else if literal == 5 {
			combo = regB
		} else if literal == 6 {
			combo = regC
		} else {
			combo = literal
		}

		fmt.Printf("pointer: %d instruction: %d literal: %d combo: %d A: %d B: %d C: %d\n", i, instruction, literal, combo, regA, regB, regC)

		switch instruction {
		case 0:
			// adv
			denominator := 1
			for range combo {
				denominator *= 2
			}
			regA = regA / denominator
		case 1:
			// bxl
			regB = regB ^ literal
		case 2:
			// bst
			regB = combo % 8
		case 3:
			// jnz
			if regA != 0 {
				i = literal - 2
			}
		case 4:
			// bxc
			regB = regB ^ regC
		case 5:
			// out
			output = append(output, '0'+rune(combo%8))
		case 6:
			// bdv
			denominator := 1
			for range combo {
				denominator *= 2
			}
			regB = regA / denominator
		case 7:
			// cdv
			denominator := 1
			for range combo {
				denominator *= 2
			}
			regC = regA / denominator
		}
	}

	sum, _ := strconv.Atoi(string(output))

	return sum
}
