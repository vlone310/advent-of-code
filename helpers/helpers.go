package helpers

import (
	"fmt"
	"os"
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

type Map map[string]string

func (m Map) Get(key string) string {
	return m[key]
}

func (m Map) Set(key, value string) {
	m[key] = value
}

func (m Map) Delete(key string) {
	delete(m, key)
}

func (m Map) Has(key string) bool {
	_, ok := m[key]
	return ok
}
