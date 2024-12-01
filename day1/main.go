package main

import (
	"errors"
	"fmt"
	"github.com/james-menzies/aoc-template-go/internal"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input := internal.ReadInput("day1/input.txt")
	fmt.Printf("First Solution: %v, Second Solution: %v", solution1(input), solution2(input))
}

func parseNums(input []string) (leftCol []int, rightCol []int, err error) {
	leftCol = make([]int, 0)
	rightCol = make([]int, 0)

	for lineNum, line := range input {
		tokens := strings.Fields(line)
		if len(tokens) != 2 {
			return nil, nil, fmt.Errorf("line #%v does not contain exactly two tokens: %v", lineNum, line)
		}

		leftNum, lErr := strconv.Atoi(tokens[0])
		rightNum, rErr := strconv.Atoi(tokens[1])

		if err := errors.Join(lErr, rErr); err != nil {
			return nil, nil, fmt.Errorf("line #%v does not contain numerals: %v", lineNum, line)
		}

		leftCol = append(leftCol, leftNum)
		rightCol = append(rightCol, rightNum)
	}

	return leftCol, rightCol, nil
}

func solution1(input []string) any {
	var result int
	leftCol, rightCol, err := parseNums(input)
	if err != nil {
		panic(fmt.Errorf("error parsing input: %w", err))
	}
	slices.Sort(leftCol)
	slices.Sort(rightCol)

	for i := 0; i < len(leftCol); i++ {
		var difference int
		if leftCol[i] < rightCol[i] {
			difference = rightCol[i] - leftCol[i]
		} else {
			difference = leftCol[i] - rightCol[i]
		}

		result += difference
	}

	return result
}

func solution2(input []string) any {
	var result int
	leftCol, rightCol, err := parseNums(input)
	if err != nil {
		panic(fmt.Errorf("error parsing input: %w", err))
	}

	rightColOccurrences := make(map[int]int)
	for _, num := range rightCol {
		existing, ok := rightColOccurrences[num]
		if !ok {
			rightColOccurrences[num] = 1
		} else {
			rightColOccurrences[num] = existing + 1
		}
	}

	for _, num := range leftCol {
		existing, ok := rightColOccurrences[num]
		if ok {
			result += num * existing
		}
	}

	return result
}
