package main

import (
	"fmt"
	"github.com/james-menzies/aoc-template-go/internal"
	"regexp"
	"strconv"
	"strings"
)

var mulRegex = regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
var doRegex = regexp.MustCompile(`do\(\)`)
var dontRegex = regexp.MustCompile(`don't\(\)`)

func main() {
	input := internal.ReadInput("day3/input.txt")
	fmt.Printf("First Solution: %v, Second Solution: %v", solution1(input), solution2(input))
}

func calculateMulExpression(exp string) int {
	stripped := exp[4 : len(exp)-1]
	tokens := strings.Split(stripped, ",")
	a, _ := strconv.Atoi(tokens[0])
	b, _ := strconv.Atoi(tokens[1])
	return a * b
}

func solution1(input []string) any {

	mulExpressions := make([]string, 0)
	for _, line := range input {
		mulExpressions = append(mulExpressions, mulRegex.FindAllString(line, -1)...)
	}

	var result int
	for _, exp := range mulExpressions {
		result += calculateMulExpression(exp)
	}

	return result
}

func solution2(input []string) any {

	mulExpressions := make([]string, 0)

	for _, line := range input {
		var from int
		var to int
		var currentDo int

		dos := doRegex.FindAllStringIndex(line, -1)
		donts := dontRegex.FindAllStringIndex(line, -1)

		for _, dont := range donts {

		}

	}

	return nil
}
