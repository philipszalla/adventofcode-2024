package main

import (
	"github.com/philipszalla/adventofcode-2024/day01"
	"github.com/philipszalla/adventofcode-2024/day02"
	"github.com/philipszalla/adventofcode-2024/day03"
	"github.com/philipszalla/adventofcode-2024/day04"
	"github.com/philipszalla/adventofcode-2024/day05"
	"github.com/philipszalla/adventofcode-2024/day06"
	"github.com/philipszalla/adventofcode-2024/day07"
	"github.com/philipszalla/adventofcode-2024/day08"
	"github.com/philipszalla/adventofcode-2024/day09"
	"github.com/philipszalla/adventofcode-2024/day10"
	"github.com/philipszalla/adventofcode-2024/day11"
	"github.com/philipszalla/adventofcode-2024/day12"
	"github.com/philipszalla/adventofcode-2024/day13"
	"github.com/philipszalla/adventofcode-2024/day15"
	"github.com/philipszalla/adventofcode-2024/day16"
	"github.com/philipszalla/adventofcode-2024/day17"
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

	utils.RunPartWithFile(8, 1, day08.Part1, "day08/puzzle.txt")
	utils.RunPartWithFile(8, 2, day08.Part2, "day08/puzzle.txt")

	utils.RunPartWithFile(9, 1, day09.Part1, "day09/puzzle.txt")
	utils.RunPartWithFile(9, 2, day09.Part2, "day09/puzzle.txt")
	utils.RunPartWithFile(9, 2, day09.Part2b, "day09/puzzle.txt")

	utils.RunPartWithFile(10, 1, day10.Part1, "day10/puzzle.txt")
	utils.RunPartWithFile(10, 2, day10.Part2, "day10/puzzle.txt")

	utils.RunPartWithFile(11, 1, day11.Part1, "day11/puzzle.txt")
	utils.RunPartWithFile(11, 1, day11.Part1b, "day11/puzzle.txt")
	utils.RunPartWithFile(11, 2, day11.Part2, "day11/puzzle.txt")

	utils.RunPartWithFile(12, 1, day12.Part1, "day12/puzzle.txt")

	utils.RunPartWithFile(13, 1, day13.Part1, "day13/puzzle.txt")
	utils.RunPartWithFile(13, 2, day13.Part2, "day13/puzzle.txt")

	utils.RunPartWithFile(15, 1, day15.Part1, "day15/puzzle.txt")

	utils.RunPartWithFile(16, 1, day16.Part1, "day16/puzzle.txt")
	utils.RunPartWithFile(16, 2, day16.Part2, "day16/puzzle.txt")

	utils.RunPartWithFile(17, 1, day17.Part1, "day17/puzzle.txt")
}
