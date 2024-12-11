package day11

import (
	"strconv"
	"strings"
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

	// fmt.Printf("Stones: %v\n", stones)

	cache := make(map[int][]int)

	for iteration := 0; iteration < iterations; iteration++ {
		// start := time.Now()

		newStones := make([]int, len(stones)*2)
		i := 0

		for _, stone := range stones {
			tuple, ok := cache[stone]

			if ok {
				copy(newStones[i:], tuple)
				i += len(tuple)
				continue
			}

			if stone == 0 {
				newStones[i] = 1
				i++
				continue
			}

			str := strconv.Itoa(stone)
			numberLen := len(str)
			if numberLen%2 == 0 {
				newStones[i], _ = strconv.Atoi(str[:numberLen/2])
				newStones[i+1], _ = strconv.Atoi(str[numberLen/2:])

				cache[stone] = newStones[i : i+2]

				i += 2
				continue
			}

			newStones[i] = stone * 2024
			cache[stone] = newStones[i : i+1]
			i++
		}

		stones = newStones[:i]

		// fmt.Printf("Stones: %v\n", stones)

		// end := time.Now()
		// elapsed := end.Sub(start)
		// fmt.Printf("Iteration: %2d, stone count: %12d, elapsed time: %s\n", iteration, len(stones), elapsed)
	}

	return len(stones)
}
