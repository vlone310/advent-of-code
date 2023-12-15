package helpers

import (
	"fmt"
	"os"
	"regexp"
)

func ReadFile(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	return string(content)
}

func IsDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func SplitByMultipleSeparators(input string, splitter string) []string {
	r := regexp.MustCompile(splitter)
	return r.Split(input, -1)
}

func IsOutOfBound(index int, lenght int) bool {
	return index < 0 || index >= lenght
}
