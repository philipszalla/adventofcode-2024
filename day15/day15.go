package day15

import (
	"strings"
)

var diffX map[rune]int = map[rune]int{
	'^': 0,
	'>': 1,
	'v': 0,
	'<': -1,
}
var diffY map[rune]int = map[rune]int{
	'^': -1,
	'>': 0,
	'v': 1,
	'<': 0,
}

func Part1(lines []string) int {
	warehouse := make([][]rune, 0)
	var height int
	for i, line := range lines {
		if len(line) == 0 {
			height = i
			break
		}

		warehouse = append(warehouse, []rune(line))
	}

	moves := strings.Join(lines[height+1:], "")

	width := len(warehouse[0])

	// fmt.Printf("loaded warehouse\n")

	var curX, curY int
	for i := range width * height {
		x := i % width
		y := i / height

		if warehouse[y][x] == '@' {
			curX = x
			curY = y
			break
		}
	}

	// fmt.Printf("found start at %d|%d\n", curX, curY)

	for _, move := range moves {
		// fmt.Printf("from %d|%d move to %c\n", curX, curY, move)
		// printWarehouse(warehouse)

		if !canMove(warehouse, curX, curY, diffX[move], diffY[move]) {
			continue
		}

		newX := curX + diffX[move]
		newY := curY + diffY[move]
		newField := warehouse[newY][newX]

		warehouse[newY][newX] = '@'
		warehouse[curY][curX] = '.'

		curX = newX
		curY = newY

		for newField == 'O' {
			newX = newX + diffX[move]
			newY = newY + diffY[move]
			newField = warehouse[newY][newX]

			warehouse[newY][newX] = 'O'
		}
	}

	sum := 0
	for i := range width * height {
		x := i % width
		y := i / height

		if warehouse[y][x] != 'O' {
			continue
		}

		sum += 100*y + x
	}

	return sum
}

func canMove(warehouse [][]rune, x, y, diffX, diffY int) bool {
	// fmt.Printf("from %d|%d move to %d|%d\n", x, y, diffX, diffY)

	for {
		x += diffX
		y += diffY

		field := warehouse[y][x]

		// fmt.Printf("found %c at %d|%d\n", field, x, y)

		if field == '.' {
			return true
		}
		if field == '#' {
			return false
		}
	}
}

// func printWarehouse(warehouse [][]rune) {
// 	for _, line := range warehouse {
// 		fmt.Println(string(line))
// 	}
// }
