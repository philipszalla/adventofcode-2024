package day16

func Part1(lines []string) int {
	width := len(lines[0])
	height := len(lines)

	for i := range width * height {
		x := i % width
		y := i / height

		if rune(lines[y][x]) != 'S' {
			continue
		}

		path := make([]int, 1<<18)
		return findPath(lines, width, height, x, y, 1, 0, path)
	}

	return 0
}

const MaxInt = int((^uint(0)) >> 1)

var diffX = []int{0, 1, 0, -1}
var diffY = []int{-1, 0, 1, 0}

func findPath(lines []string, width, height, x, y, dir int, points int, path []int) int {
	if rune(lines[y][x]) == 'E' {
		return points
	}

	key := y<<10 | x<<2 | dir

	// Loop detection
	if path[key] != 0 && path[key] < points {
		return MaxInt
	}
	path[key] = points

	newPoints := MaxInt
	for i := 0; i < 4; i++ {
		// Don't go back
		if i == 2 {
			continue
		}

		newDir := (dir + i) % 4
		newX := x + diffX[newDir]
		newY := y + diffY[newDir]

		if rune(lines[newY][newX]) == '#' {
			continue
		}

		rotations := i
		if rotations == 3 {
			rotations = 1
		}
		newPoints = min(newPoints, findPath(lines, width, height, newX, newY, newDir, points+rotations*1000+1, path))
	}

	return newPoints
}
