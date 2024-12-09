package day09

func Part1(lines []string) int {
	line := lines[0]
	blocks := make([]int, 0)
	freeBlocks := make([]int, 0)

	// fmt.Printf("line: %s\n", line)

	// parse blocks
	for i, char := range line {
		count := int(char - '0')

		value := -1
		if i%2 == 0 {
			value = i / 2
		}

		// ignore last blocks, when free
		if value == -1 && i == len(line)-1 {
			break
		}

		for j := 0; j < count; j++ {
			if value == -1 {
				freeBlocks = append(freeBlocks, len(blocks))
			}

			blocks = append(blocks, value)
		}
	}

	// fmt.Printf("blocks: %v\n", blocks)

	// sort blocks
	for _, freeBlockIndex := range freeBlocks {
		// get last filled block
		var value, i int
		for i = len(blocks) - 1; i >= 0; i-- {
			value = blocks[i]

			if value != -1 {
				break
			}
		}

		if freeBlockIndex >= len(blocks) {
			break
		}
		blocks[freeBlockIndex] = value
		blocks = blocks[:i]
	}

	// fmt.Printf("sorted blocks: %v\n", blocks)

	sum := 0
	for i, fileIndex := range blocks {
		sum += i * fileIndex
	}

	return sum
}
