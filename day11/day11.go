package day11

import (
	"strconv"
	"strings"
)

func Part1(lines []string) int {
	return getStoneCount(lines, 25)
}

func getStoneCount(lines []string, iterations int) int {
	stones := strings.Split(lines[0], " ")

	for iteration := 0; iteration < iterations; iteration++ {
		// fmt.Printf("%v\n", stones)

		for i := 0; i < len(stones); i++ {
			stone := stones[i]
			if stone == "0" {
				stones[i] = "1"
				continue
			}

			numberLen := len(stone)
			if numberLen%2 == 0 {
				stones = append(stones[:i+1], stones[i:]...)

				stones[i] = stone[:numberLen/2]
				num2, _ := strconv.Atoi(stone[numberLen/2:])
				stones[i+1] = strconv.Itoa(num2)
				i++

				continue
			}

			num, _ := strconv.Atoi(stones[i])
			stones[i] = strconv.Itoa(num * 2024)
		}
	}

	return len(stones)
}
