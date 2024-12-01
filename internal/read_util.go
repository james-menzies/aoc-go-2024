package internal

import (
	"bufio"
	"fmt"
	"os"
)

// ReadInput takes a filename, and returns a slice of strings where each string is a line in the input
func ReadInput(filename string) []string {
	result := make([]string, 0)

	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Errorf("could not read AOC input data: %w", err))
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Errorf("could not read AOC input date: %w", err))
	}
	return result
}
