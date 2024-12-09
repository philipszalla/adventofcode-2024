package day07

import (
	"math"
	"strconv"
	"strings"
)

// func Part1(lines []string) int {
// 	sum := 0

// 	for _, line := range lines {
// 		sum += solveLine(line, combinePart1)
// 	}

//		return sum
//	}
func Part1(lines []string) int {
	return solveLines(lines, combinePart1)
}

func Part2(lines []string) int {
	return solveLines(lines, combinePart2)
}

func solveLines(lines []string, combineFunc func(int, int, []int) bool) int {
	results := make(chan int, len(lines))

	for _, line := range lines {
		go func(line string) {
			results <- solveLine(line, combineFunc)
		}(line)
	}

	sum := 0

	for range lines {
		sum += <-results
	}

	return sum
}

func solveLine(line string, combineFunc func(int, int, []int) bool) int {
	lineParts := strings.Split(line, ": ")

	expectedResult, _ := strconv.Atoi(lineParts[0])

	numStrings := strings.Split(lineParts[1], " ")
	nums := make([]int, len(numStrings))
	for i, numString := range numStrings {
		nums[i], _ = strconv.Atoi(numString)
	}

	if combineFunc(expectedResult, nums[0], nums[1:]) {
		return expectedResult
	} else {
		return 0
	}
}

func combinePart1(expectedResult, current int, nums []int) bool {
	if len(nums) == 0 {
		return expectedResult == current
	}

	return combinePart1(expectedResult, current+nums[0], nums[1:]) || combinePart1(expectedResult, current*nums[0], nums[1:])
}

func combinePart2(expectedResult, current int, nums []int) bool {
	if len(nums) == 0 {
		return expectedResult == current
	}

	concatedNums := current*int(math.Pow10(len(strconv.Itoa(nums[0])))) + nums[0]
	// fmt.Printf("concat %d and %d to %d\n", current, nums[0], concatedNums)

	return combinePart2(expectedResult, current+nums[0], nums[1:]) || combinePart2(expectedResult, current*nums[0], nums[1:]) || combinePart2(expectedResult, concatedNums, nums[1:])
}
