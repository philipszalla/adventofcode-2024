package day06

import (
	"strings"
)

type Pos struct {
	x int
	y int
}

func findStart(lines []string) Pos {
	for y, line := range lines {
		x := strings.Index(line, "^")
		if x != -1 {
			return Pos{x, y}
		}
	}

	return Pos{-1, -1}
}

func Part1(lines []string) int {
	height := len(lines)
	width := len(lines[0])

	watchedFields := make(map[Pos]bool)

	current := findStart(lines)
	dir := 0

	for {
		watchedFields[current] = true
		// fmt.Printf("Pos: %o\n", current)

		// get next potential position
		nextPos := current
		switch dir {
		case 0:
			nextPos = Pos{current.x, current.y - 1}
		case 1:
			nextPos = Pos{current.x + 1, current.y}
		case 2:
			nextPos = Pos{current.x, current.y + 1}
		case 3:
			nextPos = Pos{current.x - 1, current.y}
		}

		// is next potential position in grid?
		if nextPos.x < 0 || nextPos.x >= width || nextPos.y < 0 || nextPos.y >= height {
			break
		}

		// is next potential position a wall?
		if rune(lines[nextPos.y][nextPos.x]) == '#' {
			// turn right
			dir = (dir + 1) % 4
			switch dir {
			case 0:
				nextPos = Pos{current.x, current.y - 1}
			case 1:
				nextPos = Pos{current.x + 1, current.y}
			case 2:
				nextPos = Pos{current.x, current.y + 1}
			case 3:
				nextPos = Pos{current.x - 1, current.y}
			}

			// is next position in grid?
			if nextPos.x < 0 || nextPos.x >= width || nextPos.y < 0 || nextPos.y >= height {
				break
			}
		}

		current = nextPos
	}

	return len(watchedFields)
}
