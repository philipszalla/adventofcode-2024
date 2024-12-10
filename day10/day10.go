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

			foundEnds := make([]uint64, height)

			sum += traversePaths(lines, width, height, x, y, -1, 0, func(x, y int) bool {
				// fmt.Printf("%d,%d: %b\n", x, y, foundEnds[y])
				if (foundEnds[y] & (uint64(1) << x)) == uint64(1)<<x {
					return false
				}

				foundEnds[y] |= uint64(1) << x

				return true
			})
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

			sum += traversePaths(lines, width, height, x, y, -1, 0, func(int, int) bool {
				return true
			})
		}
	}

	return sum
}

var deltaX = []int{0, 1, 0, -1}
var deltaY = []int{-1, 0, 1, 0}

func traversePaths(lines []string, width, height, x, y, fromDir, counter int, foundFn func(int, int) bool) int {
	if !isInGrid(width, height, x, y) {
		return 0
	}

	if int(lines[y][x]-'0') != counter {
		return 0
	}

	if counter == 9 {
		if !foundFn(x, y) {
			return 0
		}

		// fmt.Printf("Found end at %d,%d\n", x, y)

		return 1
	}

	sum := 0
	for i := range deltaX {
		if i == fromDir {
			continue
		}

		sum += traversePaths(lines, width, height, x+deltaX[i], y+deltaY[i], (i+2)%4, counter+1, foundFn)
	}

	return sum
}

func isInGrid(width, height, x, y int) bool {
	return x >= 0 && x < width && y >= 0 && y < height
}
