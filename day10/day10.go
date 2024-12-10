package day10

func Part1(lines []string) int {
	width := len(lines[0])
	height := len(lines)

	sum := 0

	for i := 0; i < width*height; i++ {
		x := i % width
		y := i / height

		if rune(lines[y][x]) == '0' {
			// fmt.Printf("Found trailhead at %d,%d\n", x, y)

			foundEnds := make([][]bool, height)
			for i := range foundEnds {
				foundEnds[i] = make([]bool, width)
			}

			sum += traversePaths(lines, width, height, x, y, -1, 0, foundEnds, true)
		}
	}

	return sum
}

func Part2(lines []string) int {
	width := len(lines[0])
	height := len(lines)

	sum := 0

	for i := 0; i < width*height; i++ {
		x := i % width
		y := i / height

		if rune(lines[y][x]) == '0' {
			// fmt.Printf("Found trailhead at %d,%d\n", x, y)

			foundEnds := make([][]bool, height)
			for i := range foundEnds {
				foundEnds[i] = make([]bool, width)
			}

			sum += traversePaths(lines, width, height, x, y, -1, 0, foundEnds, false)
		}
	}

	return sum
}

var deltaX = []int{0, 1, 0, -1}
var deltaY = []int{-1, 0, 1, 0}

func traversePaths(lines []string, width, height, x, y, fromDir, counter int, foundEnds [][]bool, part1 bool) int {
	if !isInGrid(width, height, x, y) {
		return 0
	}

	if int(lines[y][x]-'0') != counter {
		return 0
	}

	if counter == 9 {
		if part1 && foundEnds[y][x] {
			return 0
		}

		foundEnds[y][x] = true
		// fmt.Printf("Found end at %d,%d\n", x, y)

		return 1
	}

	sum := 0
	for i := range deltaX {
		if i == fromDir {
			continue
		}

		sum += traversePaths(lines, width, height, x+deltaX[i], y+deltaY[i], (i+2)%4, counter+1, foundEnds, part1)
	}

	return sum
}

func isInGrid(width, height, x, y int) bool {
	return x >= 0 && x < width && y >= 0 && y < height
}
