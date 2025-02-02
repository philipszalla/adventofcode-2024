package day11

import (
	"fmt"
	"strconv"
	"strings"
)

func Part1(lines []string) int {
	return getStoneCount(lines, 25)
}

func Part1b(lines []string) int {
	return getStoneCountRecursive(lines, 25, 12)
}

func Part2(lines []string) int {
	return getStoneCountRecursive(lines, 75, 17)
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

var ITERATION_BITS = 7

func getStoneCountRecursive(lines []string, iterations int, cacheSize int) int {
	if iterations > 1<<(ITERATION_BITS)-1 {
		panic(fmt.Sprintf("only supports up to %d iterations", 1<<(ITERATION_BITS+1)-1))
	}

	stoneStrings := strings.Split(lines[0], " ")

	cache := make(map[int]int, 1<<cacheSize)

	sum := 0
	for _, str := range stoneStrings {
		stone, _ := strconv.Atoi(str)
		sum += countStones(stone, iterations, cache)
	}

	// fmt.Printf("cache size: %d\n", len(cache))

	return sum
}

func countStones(stone int, iteration int, cache map[int]int) int {
	key := stone<<ITERATION_BITS | iteration
	cached, ok := cache[key]
	if ok {
		return cached
	}

	if iteration == 0 {
		return 1
	}

	var sum int
	if stone == 0 {
		sum = countStones(1, iteration-1, cache)
	} else {
		num := stone
		numberLen := 0
		for num > 0 {
			num /= 10
			numberLen++
		}

		if numberLen%2 == 1 {
			sum = countStones(stone*2024, iteration-1, cache)
		} else {
			factor := 1
			halfLen := numberLen / 2
			for i := 0; i < halfLen; i++ {
				factor *= 10
			}

			a := stone / factor
			b := stone - a*factor

			sum = countStones(a, iteration-1, cache) + countStones(b, iteration-1, cache)
		}
	}

	cache[key] = sum

	return sum
}
