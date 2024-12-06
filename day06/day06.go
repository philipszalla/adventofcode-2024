package day06

import (
	"strings"
)

type Pos struct {
	x int
	y int
}
type Pos2 struct {
	x   int16
	y   int16
	dir uint8
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

func findWatchedFields(lines []string, height, width int) map[Pos]bool {
	watchedFields := make(map[Pos]bool, height*width)

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
			nextPos = current
		}

		current = nextPos
	}

	return watchedFields
}

func Part1(lines []string) int {
	height := len(lines)
	width := len(lines[0])

	watchedFields := findWatchedFields(lines, height, width)

	return len(watchedFields)
}

func findStart2(lines []string) Pos2 {
	for y, line := range lines {
		x := strings.Index(line, "^")
		if x != -1 {
			return Pos2{int16(x), int16(y), 0}
		}
	}

	return Pos2{-1, -1, 0}
}

func getNextPos(current Pos2, dir uint8) Pos2 {
	switch dir {
	case 0:
		return Pos2{current.x, current.y - 1, dir}
	case 1:
		return Pos2{current.x + 1, current.y, dir}
	case 2:
		return Pos2{current.x, current.y + 1, dir}
	case 3:
		return Pos2{current.x - 1, current.y, dir}
	}

	return Pos2{-1, -1, 0}
}

func try(lines []string, height int16, width int16, start Pos2) int {
	current := start

	watchedFields := make([][]uint8, height)
	for i := 0; i < int(height); i++ {
		watchedFields[i] = make([]uint8, width)
	}

	for {
		savedDir := watchedFields[current.y][current.x]
		if savedDir&(1<<current.dir) == (1 << current.dir) {
			return 1
		} else {
			watchedFields[current.y][current.x] |= 1 << current.dir
		}
		// fmt.Printf("Pos: %o\n", current)

		// get next potential position
		nextPos := getNextPos(current, current.dir)

		// is next potential position in grid?
		if nextPos.x < 0 || nextPos.x >= width || nextPos.y < 0 || nextPos.y >= height {
			return 0
		}

		// is next potential position a wall?
		char := rune(lines[nextPos.y][nextPos.x])
		if char == '#' || char == 'O' {
			// turn right
			dir := (current.dir + 1) % 4
			nextPos = Pos2{current.x, current.y, dir}
		}

		current = nextPos
	}
}

func Part2(lines []string) int {
	height := int16(len(lines))
	width := int16(len(lines[0]))
	start := findStart2(lines)

	count := 0
	res := make(chan int)

	watchedFields := findWatchedFields(lines, int(height), int(width))
	delete(watchedFields, Pos{int(start.x), int(start.y)})

	for field := range watchedFields {
		count++
		go func(lines []string, height int16, width int16, start Pos2, x, y int) {
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
		}(lines, height, width, start, field.x, field.y)
	}

	sum := 0
	for i := 0; i < count; i++ {
		sum += <-res
	}

	return sum
}
