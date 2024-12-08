package day08

type Pos struct {
	x, y int
}

func Part1(lines []string) int {
	return iteratePuzzle(lines, antinodesPart1)
}

func antinodesPart1(a, b Pos, width, height int) []Pos {
	antinodes := make([]Pos, 0)

	deltaX := b.x - a.x
	deltaY := b.y - a.y

	antinode1 := Pos{b.x + deltaX, b.y + deltaY}
	if antinode1.x >= 0 && antinode1.x < width && antinode1.y >= 0 && antinode1.y < height {
		antinodes = append(antinodes, antinode1)
	}

	antinode2 := Pos{a.x - deltaX, a.y - deltaY}
	if antinode2.x >= 0 && antinode2.x < width && antinode2.y >= 0 && antinode2.y < height {
		antinodes = append(antinodes, antinode2)
	}

	return antinodes
}

func Part2(lines []string) int {
	return iteratePuzzle(lines, antinodesPart2)
}

func antinodesPart2(a, b Pos, width, height int) []Pos {
	antinodes := make([]Pos, 0)

	deltaX := b.x - a.x
	deltaY := b.y - a.y

	for i := 0; true; i++ {
		antinode := Pos{b.x + deltaX*i, b.y + deltaY*i}
		if antinode.x < 0 || antinode.x >= width || antinode.y < 0 || antinode.y >= height {
			break
		}

		antinodes = append(antinodes, antinode)
	}

	for i := 0; true; i++ {
		antinode := Pos{b.x - deltaX*i, b.y - deltaY*i}
		if antinode.x < 0 || antinode.x >= width || antinode.y < 0 || antinode.y >= height {
			break
		}

		antinodes = append(antinodes, antinode)
	}

	return antinodes
}

func iteratePuzzle(lines []string, antinodeFn func(Pos, Pos, int, int) []Pos) int {
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
				newAntinodes := antinodeFn(currentNode, node, width, height)
				for _, newAntinode := range newAntinodes {
					antinodes[newAntinode.y][newAntinode.x] = true
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
