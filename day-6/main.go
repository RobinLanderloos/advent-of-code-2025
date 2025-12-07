package day6

import (
	"fmt"
	"regexp"
	"robinlanderloos/aoc2025/io"
	"strconv"
)

type problem struct {
	numbers   []int
	operation string
	result    int
}

func newProblem(size int) *problem {
	return &problem{
		numbers: make([]int, size),
	}
}

func (problem *problem) solve() {
	fmt.Printf("Solving %s for %v\r\n", problem.operation, problem.numbers)
	switch problem.operation {
	case "*":
		result := 1
		for _, number := range problem.numbers {
			result *= number
		}
		problem.result = result
		return
	case "+":
		result := 0
		for _, number := range problem.numbers {
			result += number
		}
		problem.result = result
		return
	}

	panic("unknown operator found")
}

type transformFunc func([]string) []*problem

var extractNumbersFromLineRegex *regexp.Regexp
var extractOperationsFromLineRegex *regexp.Regexp

func Main() {
	extractNumbersFromLineRegex = regexp.MustCompile(`(\d+)`)
	extractOperationsFromLineRegex = regexp.MustCompile(`([*+])`)

	solve("day-6/input.txt", transformInputP2)
}

func solve(path string, transformFunc transformFunc) {
	lines := io.ReadLines(path)

	problems := transformFunc(lines)

	grandTotal := 0
	for _, problem := range problems {
		fmt.Printf("Problem: %v\r\n", problem)
		grandTotal += problem.result
	}

	fmt.Printf("Grand total: %d", grandTotal)
}

func transformInputP2(lines []string) []*problem {
	result := make([]*problem, 0)

	problem := newProblem(0)
	problemNumberIndex := 0
	for column := len(lines[0]) - 1; column >= 0; column-- {
		numberStr := getNumberFromColumn(lines, column)
		number, _ := strconv.Atoi(numberStr)
		problem.numbers = append(problem.numbers, number)
		problemNumberIndex++
		// Check if we have the operation, this is also the last column for any numbers
		operation := string(lines[len(lines)-1][column])
		if operation != " " {
			problem.operation = operation
			problem.solve()
			// When we find the operation, we'll also skip the next column as it'll always be empty
			column--

			// Create new problem and add current problem to result
			result = append(result, problem)
			problem = newProblem(0)
			problemNumberIndex = 0
		}
	}

	return result

}

func getNumberFromColumn(lines []string, column int) string {
	number := ""
	// Read to row before operations row
	for row := 0; row < len(lines)-1; row++ {
		curr := string(lines[row][column])
		if curr != " " {
			number += curr
		}
	}
	return number
}

func transformInputP1(lines []string) []*problem {
	problems := make([]*problem, 0)

	for row, line := range lines {
		// Operations row, all problems should've been instantiated
		if row == len(lines)-1 {
			operations := extractOperationsFromLineRegex.FindAllStringSubmatch(line, -1)
			for operationCol, operation := range operations {
				problems[operationCol].operation = operation[0]
				problems[operationCol].solve()
			}
			continue
		}

		numbers := extractNumbersFromLineRegex.FindAllStringSubmatch(line, -1)
		for column, numberStr := range numbers {
			var problem *problem
			if len(problems) <= column {
				problem = newProblem(len(lines) - 1)
				problems = append(problems, problem)
			} else {
				problem = problems[column]
			}
			convertedNumber, _ := strconv.Atoi(numberStr[0])
			problem.numbers[row] = convertedNumber
		}
	}

	return problems
}
