package day12

func Part1(lines []string) int {
	width := len(lines[0])
	height := len(lines)

	processedFields := make([]bool, width*height)

	sum := 0
	for i := 0; i < width*height; i++ {
		x := i % width
		y := i / height

		area, edges := getRegion(lines, width, height, x, y, processedFields)
		sum += area * edges
	}

	return sum
}

var diffX = []int{0, 1, 0, -1}
var diffY = []int{-1, 0, 1, 0}

func getRegion(lines []string, width, height, x, y int, processedFields []bool) (int, int) {
	key := y*height + x
	if processedFields[key] {
		return 0, 0
	}

	area := 1
	edges := 0

	plant := lines[y][x]
	processedFields[key] = true

	for i := range 4 {
		newX := x + diffX[i]
		newY := y + diffY[i]

		if newX < 0 || newX >= width || newY < 0 || newY >= height {
			edges++
			continue
		}

		newPlant := lines[newY][newX]
		if newPlant != plant {
			edges++
			continue
		}

		newArea, newEdges := getRegion(lines, width, height, newX, newY, processedFields)
		area += newArea
		edges += newEdges
	}

	return area, edges
}
