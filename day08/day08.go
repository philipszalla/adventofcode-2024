package day08

type Pos struct {
	x, y int
}

func Part1(lines []string) int {
	width := len(lines[0])
	height := len(lines)
	chars := make(map[rune][]Pos)

	antinodes := make([][]bool, height)
	for i := range antinodes {
		antinodes[i] = make([]bool, width)
	}

	for i := 0; i < width*height; i++ {
		x := i % width
		y := i / height
		char := rune(lines[y][x])

		if char == '.' {
			continue
		}

		currentNode := Pos{x, y}

		_, ok := chars[char]
		if ok {
			// calculate antinodes
			for _, node := range chars[char] {
				deltaX := node.x - currentNode.x
				deltaY := node.y - currentNode.y

				antinode1X := node.x + deltaX
				antinode1Y := node.y + deltaY
				if antinode1X >= 0 && antinode1X < width && antinode1Y >= 0 && antinode1Y < height {
					antinodes[antinode1Y][antinode1X] = true
				}

				antinode2X := currentNode.x - deltaX
				antinode2Y := currentNode.y - deltaY
				if antinode2X >= 0 && antinode2X < width && antinode2Y >= 0 && antinode2Y < height {
					antinodes[antinode2Y][antinode2X] = true
				}
			}

			chars[char] = append(chars[char], currentNode)
		} else {
			chars[char] = []Pos{currentNode}
		}
	}

	sum := 0
	for _, row := range antinodes {
		for _, cell := range row {
			if cell {
				sum++
			}
		}
	}

	return sum
}
