package day05

import (
	"slices"
	"strconv"
	"strings"
)

type Order struct {
	a string
	b string
}

func Part1(lines []string) int {
	// orders := make([]Order, 0)
	orders := make(map[string][]string)

	sum := 0
	isConfigSection := true
	for _, line := range lines {
		if line == "" {
			isConfigSection = false
			continue
		}

		if isConfigSection {
			nums := strings.Split(line, "|")
			// orders = append(orders, Order{nums[0], nums[1]})
			_, ok := orders[nums[0]]
			if ok {
				orders[nums[0]] = append(orders[nums[0]], nums[1])
			} else {
				orders[nums[0]] = make([]string, 1)
				orders[nums[0]][0] = nums[1]
			}
		} else {
			nums := strings.Split(line, ",")

			rightOrder := true
			for i, num := range nums {
				order, ok := orders[num]
				if ok {
					for _, toCheck := range order {
						if slices.Contains(nums[:i], toCheck) {
							rightOrder = false
							break
						}
					}

					if !rightOrder {
						break
					}
				}
			}

			if rightOrder {
				numValue, _ := strconv.Atoi(nums[len(nums)/2])
				sum += numValue
			}
		}
	}

	return sum
}

func Part2(lines []string) int {
	// orders := make([]Order, 0)
	orders := make(map[string][]string)

	sum := 0
	isConfigSection := true
	for _, line := range lines {
		if line == "" {
			isConfigSection = false
			continue
		}

		if isConfigSection {
			nums := strings.Split(line, "|")
			// orders = append(orders, Order{nums[0], nums[1]})
			_, ok := orders[nums[0]]
			if ok {
				orders[nums[0]] = append(orders[nums[0]], nums[1])
			} else {
				orders[nums[0]] = make([]string, 1)
				orders[nums[0]][0] = nums[1]
			}
		} else {
			nums := strings.Split(line, ",")

			rightOrder := true
			for i := 0; i < len(nums); i++ {
				num := nums[i]

				order, ok := orders[num]
				if ok {
					for _, toCheck := range order {
						wrongIndex := slices.Index(nums[:i], toCheck)
						if wrongIndex != -1 {
							rightOrder = false

							// correct order
							newNums := make([]string, len(nums))
							copy(newNums, nums)
							newNums = append(newNums[:wrongIndex], newNums[wrongIndex+1:]...)
							newNums = append(newNums, toCheck)
							nums = newNums

							i--
						}
					}
				}
			}

			if !rightOrder {
				numValue, _ := strconv.Atoi(nums[len(nums)/2])
				sum += numValue
			}
		}
	}

	return sum
}
