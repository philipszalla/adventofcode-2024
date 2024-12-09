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

type BlockSpan struct {
	index, size int
}

func Part2(lines []string) int {
	line := lines[0]
	blocks := make([]int, 0)
	freeBlocks := make([]BlockSpan, 0)
	fileBlocks := make([]BlockSpan, 0)

	// fmt.Printf("line: %s\n", line)
	// fmt.Printf("line length: %d\n", len(line))

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

		if value == -1 {
			freeBlocks = append(freeBlocks, BlockSpan{len(blocks), count})
		} else {
			fileBlocks = append(fileBlocks, BlockSpan{len(blocks), count})
		}

		for j := 0; j < count; j++ {
			blocks = append(blocks, value)
		}
	}

	// fmt.Printf("freeBlocks: %v\n", freeBlocks)
	// fmt.Printf("fileBlocks: %v\n", fileBlocks)
	// fmt.Printf("blocks: %v\n", blocks)
	// fmt.Printf("block count: %d\n", len(blocks))

	// sort blocks
	for i := len(fileBlocks) - 1; i >= 0; i-- {
		fileBlock := fileBlocks[i]

		for j := 0; j < len(freeBlocks); j++ {
			freeBlock := freeBlocks[j]

			if fileBlock.size > freeBlock.size || fileBlock.index <= freeBlock.index {
				continue
			}

			for k := 0; k < fileBlock.size; k++ {
				blocks[freeBlock.index+k] = blocks[fileBlock.index+k]
				blocks[fileBlock.index+k] = -1
			}

			if freeBlock.size == fileBlock.size {
				freeBlocks = append(freeBlocks[:j], freeBlocks[j+1:]...)
			} else {
				freeBlocks[j].index += fileBlock.size
				freeBlocks[j].size -= fileBlock.size
			}

			break
		}

		// fmt.Printf("blocks: %v\n", blocks)
	}

	// fmt.Printf("sorted blocks: %v\n", blocks)

	sum := 0
	for i, fileIndex := range blocks {
		if fileIndex == -1 {
			continue
		}

		sum += i * fileIndex
	}

	return sum
}
