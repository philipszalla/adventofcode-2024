package day18

import (
	"fmt"
	"strconv"
	"strings"
)

var gridSize int8 = 71

type Field struct {
	x, y, fromDir int8
	count         int
}

func Part1(lines []string) int {
	grid := generateGrid(lines, 1024)
	cache := make([]int, len(grid))

	printGrid(grid, cache)

	// queue := list.New()
	// queue.PushBack(Field{0, 0, -1, 0})
	queue := make(chan Field)
	go func() {
		queue <- Field{0, 0, -1, 0}
	}()

	i := 0
	// for queue.Len() > 0 {
	for field := range queue {
		i++
		// element := queue.Front()
		// field := element.Value.(Field)
		// queue.Remove(element)

		// fmt.Printf("Field: %v i: %d\n", field, i)

		go findPath(grid, field.x, field.y, field.fromDir, field.count, cache, queue)

		// if i > 100 {
		// 	break
		// }

		// if (i % 1000000) == 0 {
		// 	printGrid(grid, cache)
		// }
	}

	printGrid(grid, cache)

	return cache[len(cache)-1]
}

var diffX = []int8{0, 1, 0, -1}
var diffY = []int8{-1, 0, 1, 0}

func isNotGoable(grid []bool, x, y int8) bool {
	if x < 0 || x >= gridSize || y < 0 || y >= gridSize {
		return true
	}

	key := int16(y)*int16(gridSize) + int16(x)
	return grid[key]
}

// func findPath(grid []bool, x, y, fromDir, count int, cache []int, queue *list.List) {
func findPath(grid []bool, x, y, fromDir int8, count int, cache []int, queue chan Field) {
	key := int16(y)*int16(gridSize) + int16(x)
	if cache[key] != 0 && cache[key] < count {
		// fmt.Printf("Cached is better: %d\n", cache[key])

		return
	}
	cache[key] = count

	if x == (gridSize-1) && y == (gridSize-1) {
		fmt.Printf("Found goal! %d\n", count)

		return
	}

	// printGrid(grid, cache)

	var dir int8
	for dir = range 4 {
		newX := x + diffX[dir]
		newY := y + diffY[dir]

		// Don't go back
		if ((dir + 2) % 4) == fromDir {
			continue
		}

		if isNotGoable(grid, newX, newY) {
			continue
		}

		// queue.PushBack(Field{newX, newY, dir, count + 1})
		queue <- Field{newX, newY, dir, count + 1}
	}
}

func generateGrid(lines []string, timespan int) []bool {
	grid := make([]bool, int(gridSize)*int(gridSize))

	for i := range timespan {
		coords := strings.Split(lines[i], ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])

		key := int16(y)*int16(gridSize) + int16(x)
		grid[key] = true
	}

	return grid
}

func printGrid(grid []bool, cache []int) {
	for i, blocked := range grid {
		field := '.'
		if blocked {
			field = '#'
		}
		if cache[i] != 0 {
			field = rune('0' + cache[i]%10)
		}

		fmt.Printf("%c", field)

		if ((i + 1) % int(gridSize)) == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
}
