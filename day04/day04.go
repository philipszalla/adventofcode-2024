package day04

var searchString = "XMAS"
var searchString2 = "SAMX"

func Part1(lines []string) int {
	sum := 0

	height := len(lines)
	width := len(lines[0])

	// loop over all characters
	for i := 0; i < height*width; i++ {
		y := i / height
		x := i % width

		// loop over directions right, right-bottom, bottom, left-bottom
		for dir := 0; dir < 4; dir++ {
			factorX := 1
			if dir == 2 { // b
				factorX = 0
			}
			if dir == 3 { // lb
				factorX = -1
			}

			factorY := 1
			if dir == 0 { // r
				factorY = 0
			}

			str := ""

			// loop over search string length
			for k := 0; k < len(searchString); k++ {
				charX := x + k*factorX
				charY := y + k*factorY

				if charX < 0 || charX > width-1 || charY < 0 || charY > height-1 {
					break
				}
				str += string(lines[charY][charX])
			}

			if str == searchString || str == searchString2 {
				sum += 1
				// fmt.Printf("Found str at (%d,%d) in dir %d: %s\n", x, y, dir, str)
			}
		}
	}

	return sum
}

func Part2(lines []string) int {
	sum := 0

	height := len(lines)
	width := len(lines[0])

	// loop over all characters
	for i := 0; i < height*width; i++ {
		y := i / height
		x := i % width

		if x < 1 || x > width-2 || y < 1 || y > height-2 {
			continue
		}

		if string(lines[y][x]) != "A" {
			continue
		}

		charTL := string(lines[y-1][x-1])
		charBR := string(lines[y+1][x+1])
		charTR := string(lines[y-1][x+1])
		charBL := string(lines[y+1][x-1])

		if ((charTL == "M" && charBR == "S") || (charTL == "S" && charBR == "M")) &&
			((charTR == "M" && charBL == "S") || (charTR == "S" && charBL == "M")) {
			sum += 1
			// fmt.Printf("Found str at (%d,%d)\n", x, y)
		}
	}

	return sum
}
