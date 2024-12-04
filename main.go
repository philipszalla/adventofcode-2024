package main

import (
	"github.com/philipszalla/adventofcode-2024/day01"
	"github.com/philipszalla/adventofcode-2024/day02"
	"github.com/philipszalla/adventofcode-2024/day03"
	"github.com/philipszalla/adventofcode-2024/utils"
)

func main() {
	utils.RunPartWithFile(1, 1, day01.Part1, "day01/puzzle.txt")
	utils.RunPartWithFile(1, 2, day01.Part2, "day01/puzzle.txt")

	utils.RunPartWithFile(2, 1, day02.Part1, "day02/puzzle.txt")
	utils.RunPartWithFile(2, 2, day02.Part2, "day02/puzzle.txt")

	utils.RunPartWithFile(3, 1, day03.Part1, "day03/puzzle.txt")
	utils.RunPartWithFile(3, 2, day03.Part2, "day03/puzzle.txt")
}
