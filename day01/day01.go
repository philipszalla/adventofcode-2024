package day01

import (
	"slices"
	"strconv"
	"strings"
)

func Part1(lines []string) int {
	list_a := make([]int, len(lines))
	list_b := make([]int, len(lines))
	for index, line := range lines {
		parts := strings.Split(line, "   ")
		list_a[index], _ = strconv.Atoi(parts[0])
		list_b[index], _ = strconv.Atoi(parts[1])
	}

	slices.Sort(list_a)
	slices.Sort(list_b)

	sum := 0
	for i := 0; i < len(list_a); i++ {
		sum += max(list_a[i], list_b[i]) - min(list_a[i], list_b[i])
	}

	return sum
}

func Part2(lines []string) int {
	list_a := make([]int, len(lines))
	map_b := make(map[int]int)
	for index, line := range lines {
		parts := strings.Split(line, "   ")
		list_a[index], _ = strconv.Atoi(parts[0])

		b, _ := strconv.Atoi(parts[1])
		_, ok := map_b[b]
		if ok {
			map_b[b] += 1
		} else {
			map_b[b] = 1
		}
	}

	sum := 0
	for _, a := range list_a {
		count, ok := map_b[a]
		if ok {
			sum += a * count
		}
	}

	return sum
}
