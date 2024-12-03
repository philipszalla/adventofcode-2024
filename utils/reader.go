package utils

import (
	"os"
	"strings"
)

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
