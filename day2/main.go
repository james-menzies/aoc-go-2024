package main

import (
	"fmt"
	"github.com/james-menzies/aoc-template-go/internal"
	"strconv"
	"strings"
)

func main() {
	input := internal.ReadInput("day2/input.txt")
	fmt.Printf("First Solution: %v, Second Solution: %v", solution1(input), solution2(input))
}

func parseNums(input []string) ([][]int, error) {
	result := make([][]int, 0)

	for lineNum, line := range input {
		tokens := strings.Fields(line)

		nums := make([]int, 0)
		for tokenNum, token := range tokens {
			tokenAsInt, err := strconv.Atoi(token)
			if err != nil {
				return nil, fmt.Errorf("line #%v pos #%v is not an integer: %v", lineNum, tokenNum, token)
			}

			nums = append(nums, tokenAsInt)
		}

		result = append(result, nums)
	}

	return result, nil
}

func isDiffSafe(a, b int, ascending bool) bool {
	var diff int
	if ascending {
		diff = b - a
	} else {
		diff = a - b
	}

	return diff >= 1 || diff <= 3
}

func isReportSafe(report []int) bool {
	var safetyUsed bool
	var ascending bool = true
	if report[0] > report[1] {
		ascending = false
	}

	var prev int
	for i := 1; i < len(report); i++ {
		if !isDiffSafe(prev, report[i], ascending) {
			if safetyUsed {
				return false
			}

			if




		} else {
			prev = report[i]
		}

	}

	return true
}

func solution1(input []string) any {
	var numberOfSafeReports int
	reports, err := parseNums(input)
	if err != nil {
		panic(fmt.Errorf("could not parse AOC input: %w", err))
	}

	for _, report := range reports {
		if isReportSafe(report) {
			numberOfSafeReports++
			fmt.Printf("report %+v is safe\n", report)
		}
	}

	return numberOfSafeReports
}

func solution2(input []string) any {
	// TODO solve second challenge
	return nil
}
