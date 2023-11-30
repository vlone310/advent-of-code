package helpers

import (
	"fmt"
	"os"
)

func ReadFile(path string) (string) {
	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	return string(content);
}