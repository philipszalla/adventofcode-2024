package day06

import (
	"strings"
)

type Pos struct {
	x int16
	y int16
}

func findStart(lines []string) Pos {
	for y, line := range lines {
		x := strings.Index(line, "^")
		if x != -1 {
			return Pos{int16(x), int16(y)}
		}
	}

	return Pos{-1, -1}
}

func findWatchedFields(lines []string, height, width int16, start Pos) [][]bool {
	watchedFields := make([][]bool, height)
	for i := int16(0); i < height; i++ {
		watchedFields[i] = make([]bool, width)
	}

	current := start
	dir := uint8(0)

	for {
		watchedFields[current.y][current.x] = true
		// fmt.Printf("Pos: %o\n", current)

		// get next potential position
		nextPos := getNextPos(current, dir)

		// is next potential position in grid?
		if nextPos.x < 0 || nextPos.x >= width || nextPos.y < 0 || nextPos.y >= height {
			break
		}

		// is next potential position a wall?
		if rune(lines[nextPos.y][nextPos.x]) == '#' {
			// turn right
			dir = (dir + 1) % 4
		} else {
			current = nextPos
		}
	}

	return watchedFields
}

func Part1(lines []string) int {
	height := int16(len(lines))
	width := int16(len(lines[0]))
	start := findStart(lines)

	watchedFields := findWatchedFields(lines, height, width, start)

	sum := 0
	for _, row := range watchedFields {
		for _, cell := range row {
			if cell {
				sum++
			}
		}
	}

	return sum
}

var diffX = [...]int16{0, 1, 0, -1}
var diffY = [...]int16{-1, 0, +1, 0}

func getNextPos(current Pos, dir uint8) Pos {
	return Pos{current.x + diffX[dir], current.y + diffY[dir]}
}

func try(lines []string, height int16, width int16, start Pos) int {
	current := start
	currentDir := uint8(0)

	watchedFields := make([][]uint8, height)
	for i := 0; i < int(height); i++ {
		watchedFields[i] = make([]uint8, width)
	}

	for {
		savedDir := watchedFields[current.y][current.x]
		if savedDir&(1<<currentDir) == (1 << currentDir) {
			return 1
		} else {
			watchedFields[current.y][current.x] |= 1 << currentDir
		}
		// fmt.Printf("Pos: %o\n", current)

		// get next potential position
		nextPos := getNextPos(current, currentDir)

		// is next potential position in grid?
		if nextPos.x < 0 || nextPos.x >= width || nextPos.y < 0 || nextPos.y >= height {
			return 0
		}

		// is next potential position a wall?
		char := rune(lines[nextPos.y][nextPos.x])
		if char == '#' || char == 'O' {
			// turn right
			currentDir = (currentDir + 1) % 4
		} else {
			current = nextPos
		}

	}
}

func Part2(lines []string) int {
	height := int16(len(lines))
	width := int16(len(lines[0]))
	start := findStart(lines)

	count := 0
	res := make(chan int)

	watchedFields := findWatchedFields(lines, height, width, start)
	watchedFields[start.y][start.x] = false

	for y, row := range watchedFields {
		for x, cell := range row {
			if !cell {
				continue
			}

			count++
			go func(lines []string, height int16, width int16, start Pos, x, y int) {
				newLines := make([]string, len(lines))
				copy(newLines, lines)
				for newY := range newLines {
					if y == newY {
						line := []rune(newLines[y])
						line[x] = 'O'
						newLines[y] = string(line)
					} else {
						newLines[newY] = string(newLines[newY])
					}
				}

				res <- try(newLines, height, width, start)
			}(lines, height, width, start, x, y)
		}
	}

	sum := 0
	for i := 0; i < count; i++ {
		sum += <-res
	}

	return sum
}
