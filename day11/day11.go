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

func Part1b(lines []string) int {
	return getStoneCountParallel(lines, 25)
}

func getStoneCount(lines []string, iterations int) int {
	stoneStrings := strings.Split(lines[0], " ")
	stones := make([]int, len(stoneStrings))
	for i, str := range stoneStrings {
		stones[i], _ = strconv.Atoi(str)
	}

	for iteration := 0; iteration < iterations; iteration++ {
		start := time.Now()

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

		end := time.Now()
		elapsed := end.Sub(start)
		fmt.Printf("Iteration: %2d, stone count: %12d, elapsed time: %s\n", iteration, len(stones), elapsed)
	}

	return len(stones)
}

type Tuple struct {
	a, b int
}

func getStoneCountParallel(lines []string, iterations int) int {
	stoneStrings := strings.Split(lines[0], " ")
	stones := make([]int, len(stoneStrings))
	for i, str := range stoneStrings {
		stones[i], _ = strconv.Atoi(str)
	}

	for iteration := 0; iteration < iterations; iteration++ {
		start := time.Now()

		ch := make(chan Tuple)

		for _, stone := range stones {
			go func(stone int) {
				if stone == 0 {
					ch <- Tuple{1, -1}
					return
				}

				str := strconv.Itoa(stone)
				numberLen := len(str)
				if numberLen%2 == 0 {
					a, _ := strconv.Atoi(str[:numberLen/2])
					b, _ := strconv.Atoi(str[numberLen/2:])

					ch <- Tuple{a, b}
					return
				}

				ch <- Tuple{stone * 2024, -1}
			}(stone)
		}

		newStones := make([]int, len(stones)*2)
		i := 0
		for range stones {
			tuple := <-ch
			newStones[i] = tuple.a
			i++
			if tuple.b != -1 {
				newStones[i] = tuple.b
				i++
			}
		}
		stones = newStones[:i]

		end := time.Now()
		elapsed := end.Sub(start)
		fmt.Printf("Iteration: %2d, stone count: %12d, elapsed time: %s\n", iteration, len(stones), elapsed)
	}

	return len(stones)
}
