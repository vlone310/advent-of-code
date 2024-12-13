package helpers

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
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

func CopyMap[S comparable, T any](m map[S]T) map[S]T {
	res := make(map[S]T)
	for k, v := range m {
		res[k] = v
	}
	return res
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func Timer() func() {
	start := time.Now()
	return func() {
		fmt.Printf("took %v\n", time.Since(start))
	}
}

func ParseInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return 0
	}
	return val
}
