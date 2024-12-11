package day11_test

import (
	"testing"

	"github.com/philipszalla/adventofcode-2024/day11"
	"github.com/philipszalla/adventofcode-2024/utils"
)

var example = `125 17`

func TestPart1(t *testing.T) {
	utils.TestPart(t, day11.Part1, example, 55312)
}
