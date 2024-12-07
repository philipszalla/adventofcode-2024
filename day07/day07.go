package day07

import (
	"fmt"
	"strconv"
	"strings"
)

func Part1(lines []string) int {
	sum := 0

	for _, line := range lines {
		sum += solveLine(line)
	}

	return sum
}

func solveLine(line string) int {
	lineParts := strings.Split(line, ": ")

	expectedResult, _ := strconv.Atoi(lineParts[0])

	numStrings := strings.Split(lineParts[1], " ")
	nums := make([]int, len(numStrings))
	for i, numString := range numStrings {
		nums[i], _ = strconv.Atoi(numString)
	}

	fmt.Printf("expected: %d nums: %o\n", expectedResult, nums)

	if combineRecurse(expectedResult, nums[0], nums[1:]) {
		return expectedResult
	} else {
		return 0
	}
}

func combineRecurse(expectedResult, current int, nums []int) bool {
	if len(nums) == 0 {
		return expectedResult == current
	}

	return combineRecurse(expectedResult, current+nums[0], nums[1:]) || combineRecurse(expectedResult, current*nums[0], nums[1:])

}
