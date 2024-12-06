package day06

import (
	"strings"
)

type Pos struct {
	x int
	y int
}
type Pos2 struct {
	x   int
	y   int
	dir int
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
			nextPos = current
		}

		current = nextPos
	}

	return len(watchedFields)
}

func findStart2(lines []string) Pos2 {
	for y, line := range lines {
		x := strings.Index(line, "^")
		if x != -1 {
			return Pos2{x, y, 0}
		}
	}

	return Pos2{-1, -1, -1}
}

func getNextPos(current Pos2, dir int) Pos2 {
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

	return Pos2{-1, -1, -1}
}

func try(lines []string, height int, width int, start Pos2) int {
	current := start

	watchedFields := make(map[Pos2]bool)

	for {
		_, exists := watchedFields[current]
		if exists {
			return 1
		}

		watchedFields[current] = true
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
	height := len(lines)
	width := len(lines[0])
	start := findStart2(lines)

	count := 0
	res := make(chan int)

	for i := 0; i < width*height; i++ {
		x := i % width
		y := i / height

		char := rune(lines[y][x])
		if char == '#' || char == '^' {
			continue
		}

		count++
		go func(lines []string, height int, width int, start Pos2, x, y int) {
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

	sum := 0
	for i := 0; i < count; i++ {
		sum += <-res
	}

	return sum
}
