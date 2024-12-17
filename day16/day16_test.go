package day16_test

import (
	"testing"

	"github.com/philipszalla/adventofcode-2024/day16"
	"github.com/philipszalla/adventofcode-2024/utils"
)

var example1 = `###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############`

var example2 = `#################
#...#...#...#..E#
#.#.#.#.#.#.#.#.#
#.#.#.#...#...#.#
#.#.#.#.###.#.#.#
#...#.#.#.....#.#
#.#.#.#.#.#####.#
#.#...#.#.#.....#
#.#.#####.#.###.#
#.#.#.......#...#
#.#.###.#####.###
#.#.#...#.....#.#
#.#.#.#####.###.#
#.#.#.........#.#
#.#.#.#########.#
#S#.............#
#################`

func TestPart1(t *testing.T) {
	utils.TestPart(t, day16.Part1, example1, 7036)
	utils.TestPart(t, day16.Part1, example2, 11048)
}

func TestPart2(t *testing.T) {
	utils.TestPart(t, day16.Part2, example1, 45)
	utils.TestPart(t, day16.Part2, example2, 64)
}
