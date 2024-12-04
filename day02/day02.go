package day02

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/philipszalla/adventofcode-2024/utils"
)

func reportIsSafe(levels []int) bool {
	started := false
	is_up := false

	is_safe := true

	for i := 1; i < len(levels); i++ {
		prev := levels[i-1]
		curr := levels[i]

		if !started {
			started = true
			is_up = prev < curr
		} else {
			if is_up != (prev < curr) {
				is_safe = false
				break
			}
		}

		diff := max(prev, curr) - min(prev, curr)
		if diff < 1 || diff > 3 {
			is_safe = false
			break
		}
	}

	return is_safe
}

func Part1(filename string) int {
	lines := utils.ReadFile(filename)

	reports := make(chan int, len(lines))

	for _, line := range lines {
		go func(line string) {
			numbers := strings.Split(line, " ")
			levels := make([]int, len(numbers))
			for i, number := range numbers {
				levels[i], _ = strconv.Atoi(number)
			}

			if reportIsSafe(levels) {
				reports <- 1
			} else {
				reports <- 0
			}
		}(line)
	}

	sum := 0
	for range lines {
		sum += <-reports
	}

	fmt.Printf("day02.Part1 %d\n", sum)
	return sum
}

func Part2(filename string) int {
	lines := utils.ReadFile(filename)

	reports := make(chan int, len(lines))

	for _, line := range lines {
		go func(line string) {
			numbers := strings.Split(line, " ")
			levels := make([]int, len(numbers))
			for i, number := range numbers {
				levels[i], _ = strconv.Atoi(number)
			}

			is_safe := reportIsSafe(levels)
			if !is_safe {
				// Problem Dampener
				for index := 0; index < len(levels); index++ {
					levels2 := make([]int, len(levels))
					copy(levels2, levels) // IMPORTANT
					levels2 = append(levels2[:index], levels2[index+1:]...)

					is_safe = reportIsSafe(levels2)
					if is_safe {
						break
					}
				}
			}

			if is_safe {
				reports <- 1
			} else {
				reports <- 0
			}
		}(line)
	}

	sum := 0
	for range lines {
		sum += <-reports
	}

	fmt.Printf("day02.Part2 %d\n", sum)
	return sum
}
