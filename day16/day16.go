package day16

func Part1(lines []string) int {
	x, y := getStart(lines)

	path := make([]int, 1<<18)
	return findPath(lines, x, y, 1, 0, path)
}

func Part2(lines []string) int {
	x, y := getStart(lines)

	// Get max movements from part 1
	loops := make([]int, 1<<18)
	part1 := findPath(lines, x, y, 1, 0, loops)

	maxRotations := part1 / 1000
	maxSteps := part1 % 1000

	// Part 2

	loops = make([]int, 1<<18)
	path := make([]int, 0, maxSteps)
	fields := findPath2(lines, x, y, 1, 0, maxSteps, 0, maxRotations, loops, path)

	distinctFields := make(map[int]bool, len(fields))
	for _, field := range fields {
		distinctFields[field] = true
	}

	return len(distinctFields)
}

const MaxInt = int((^uint(0)) >> 1)

var diffX = []int{0, 1, 0, -1}
var diffY = []int{-1, 0, 1, 0}

func findPath(lines []string, x, y, dir int, points int, path []int) int {
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
		newPoints = min(newPoints, findPath(lines, newX, newY, newDir, points+rotations*1000+1, path))
	}

	return newPoints
}

func findPath2(lines []string, x, y, dir int, steps, maxSteps, rotations, maxRotations int, loops []int, path []int) []int {
	path = append(path, y<<8|x)

	if steps > maxSteps || rotations > maxRotations {
		return []int{}
	}

	if rune(lines[y][x]) == 'E' {
		return path
	}

	key := y<<10 | x<<2 | dir

	// Loop detection
	points := rotations*1000 + steps
	if loops[key] != 0 && loops[key] < points {
		return []int{}
	}
	loops[key] = points

	newPath := []int{}
	branches := 0
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

		branchPath := path
		if branches > 0 {
			branchPath = make([]int, len(path), cap(path))
			copy(branchPath, path)
		}
		branches++

		rots := i
		if rots == 3 {
			rots = 1
		}

		branchPath = findPath2(lines, newX, newY, newDir, steps+1, maxSteps, rotations+rots, maxRotations, loops, branchPath)
		if len(branchPath) > 0 {
			newPath = append(newPath, branchPath...)
		}
	}

	return newPath
}

func getStart(lines []string) (int, int) {
	for y, line := range lines {
		for x, field := range line {
			if field == 'S' {
				return x, y
			}
		}
	}
	return -1, -1
}
