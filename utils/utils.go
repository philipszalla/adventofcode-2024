package utils

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

type partFn func([]string) int

func RunPartWithFile(day int, part int, fn partFn, filename string) {
	lines := ReadFile(filename)
	RunPart(day, part, fn, lines)
}

func RunPart(day int, part int, fn partFn, lines []string) {
	start := time.Now()

	result := fn(lines)

	end := time.Now()
	elapsed := end.Sub(start)

	fmt.Printf("Finished Day %2d Part %d! Result: %d, Elapsed time: %s\n", day, part, result, elapsed)
}

func TestPart(t *testing.T, fn partFn, content string, expectedResult int) {
	lines := strings.Split(content, "\n")

	result := fn(lines)

	if result != expectedResult {
		t.Fatalf("Expected result is %d. But got %d", expectedResult, result)
	}
}

func ReadFile(filename string) []string {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	content := string(bytes)
	content = strings.ReplaceAll(content, "\r", "")
	lines := strings.Split(content, "\n")

	return lines
}
