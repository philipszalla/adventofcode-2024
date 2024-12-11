package day11

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func Part1(lines []string) int {
	return getStoneCount(lines, 25)
}

func getStoneCount(lines []string, iterations int) int {
	stoneStrings := strings.Split(lines[0], " ")
	stones := make([]int, len(stoneStrings))
	for i, str := range stoneStrings {
		stones[i], _ = strconv.Atoi(str)
	}

	start := time.Now()

	for iteration := 0; iteration < iterations; iteration++ {
		end := time.Now()
		elapsed := end.Sub(start)
		fmt.Printf("Elapsed time: %s, Iteration: %d, stone count: %d\n", elapsed, iteration, len(stones))
		start = time.Now()

		for i := 0; i < len(stones); i++ {
			stone := stones[i]
			if stone == 0 {
				stones[i] = 1
				continue
			}

			str := strconv.Itoa(stone)
			numberLen := len(str)
			if numberLen%2 == 0 {
				stones = append(stones[:i+1], stones[i:]...)

				stones[i], _ = strconv.Atoi(str[:numberLen/2])
				stones[i+1], _ = strconv.Atoi(str[numberLen/2:])
				i++

				continue
			}

			stones[i] *= 2024
		}
	}

	return len(stones)
}
