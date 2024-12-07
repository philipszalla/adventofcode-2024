package main

import (
	"github.com/philipszalla/adventofcode-2024/day01"
	"github.com/philipszalla/adventofcode-2024/day02"
	"github.com/philipszalla/adventofcode-2024/day03"
	"github.com/philipszalla/adventofcode-2024/day04"
	"github.com/philipszalla/adventofcode-2024/day05"
	"github.com/philipszalla/adventofcode-2024/day06"
	"github.com/philipszalla/adventofcode-2024/day07"
	"github.com/philipszalla/adventofcode-2024/utils"
)

func main() {
	utils.RunPartWithFile(1, 1, day01.Part1, "day01/puzzle.txt")
	utils.RunPartWithFile(1, 2, day01.Part2, "day01/puzzle.txt")

	utils.RunPartWithFile(2, 1, day02.Part1, "day02/puzzle.txt")
	utils.RunPartWithFile(2, 2, day02.Part2, "day02/puzzle.txt")

	utils.RunPartWithFile(3, 1, day03.Part1, "day03/puzzle.txt")
	utils.RunPartWithFile(3, 2, day03.Part2, "day03/puzzle.txt")

	utils.RunPartWithFile(4, 1, day04.Part1, "day04/puzzle.txt")
	utils.RunPartWithFile(4, 2, day04.Part2, "day04/puzzle.txt")

	utils.RunPartWithFile(5, 1, day05.Part1, "day05/puzzle.txt")
	utils.RunPartWithFile(5, 2, day05.Part2, "day05/puzzle.txt")

	utils.RunPartWithFile(6, 1, day06.Part1, "day06/puzzle.txt")
	utils.RunPartWithFile(6, 2, day06.Part2, "day06/puzzle.txt")

	utils.RunPartWithFile(7, 1, day07.Part1, "day07/puzzle.txt")
	utils.RunPartWithFile(7, 2, day07.Part2, "day07/puzzle.txt")
}
