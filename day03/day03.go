package day03

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/philipszalla/adventofcode-2024/utils"
)

func Part1(filename string) int {
	lines := utils.ReadFile(filename)
	str := strings.Join(lines, "")

	re := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
	muls := re.FindAllStringSubmatch(str, -1)

	sum := 0

	for _, mul := range muls {
		a, _ := strconv.Atoi(mul[1])
		b, _ := strconv.Atoi(mul[2])
		sum += a * b
	}

	fmt.Printf("day03.Part1 %d\n", sum)
	return sum
}

func Part2(filename string) int {
	lines := utils.ReadFile(filename)
	str := strings.Join(lines, "")

	for {
		dont_pos := strings.Index(str, "don't()")
		if dont_pos == -1 {
			break
		}
		do_pos := strings.Index(str[dont_pos:], "do()")
		if do_pos == -1 {
			str = str[:dont_pos]
			break
		} else {
			str = str[:dont_pos] + str[dont_pos+do_pos+4:]
		}
	}

	re := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
	muls := re.FindAllStringSubmatch(str, -1)

	sum := 0

	for _, mul := range muls {
		a, _ := strconv.Atoi(mul[1])
		b, _ := strconv.Atoi(mul[2])
		sum += a * b
	}

	fmt.Printf("day03.Part2 %d\n", sum)
	return sum
}
