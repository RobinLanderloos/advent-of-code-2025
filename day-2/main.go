package day2

import (
	"fmt"
	"log"
	"regexp"
	"robinlanderloos/aoc2025/io"
	"strconv"
)

var rangeRegex *regexp.Regexp

type patternFunc func(string) bool
type idRange struct {
	lower int
	upper int
}

func newIdRange(lowerStr, upperStr string) *idRange {
	lower, err := strconv.Atoi(lowerStr)
	if err != nil {
		log.Fatalf("failed to convert lower range to integer: %v", err)
	}

	upper, err := strconv.Atoi(upperStr)
	if err != nil {
		log.Fatalf("failed to convert upper range to integer: %v", err)
	}

	return &idRange{lower: lower, upper: upper}
}

func Day2() {
	rangeRegex = regexp.MustCompile(`(\d+)-(\d+)`)
	// solve("day-2/input.txt", hasRepeatingPatternP1)
	solve("day-2/input.txt", hasRepeatingPatternP2)
}

func solve(path string, patternFunc patternFunc) {
	input := io.ReadLines(path)[0]
	idRanges := getRanges(input)
	result := 0

	for _, r := range idRanges {
		fmt.Printf("Evaluating range: %v\r\n", r)
		for i := r.lower; i <= r.upper; i++ {
			iStr := strconv.Itoa(i)
			if patternFunc(iStr) {
				fmt.Printf("\tInvalid found: %d\r\n", i)
				result += i
			}
		}
	}

	fmt.Printf("Result: %d\r\n", result)
}

func getRanges(input string) []idRange {
	matches := rangeRegex.FindAllStringSubmatch(input, -1)
	idRanges := make([]idRange, 0, len(matches))

	for _, match := range matches {
		idRanges = append(idRanges, *newIdRange(match[1], match[2]))
	}

	return idRanges
}

func hasRepeatingPatternP1(s string) bool {
	length := len(s)

	// No need to evaluate odd length strings
	if length%2 > 0 {
		return false
	}

	// Try all possible pattern lengths from 1 to length/2
	left := s[:length/2]
	right := s[length/2:]

	return left == right
}

func hasRepeatingPatternP2(s string) bool {
	length := len(s)

	// Only evaluate if the pattern can be repeated throughout the whole length
	for patternLength := 1; patternLength <= length/2; patternLength++ {
		if length%patternLength != 0 {
			continue
		}

		hasPattern := true
		currPattern := s[:patternLength]
		currPatternLength := len(currPattern)
		for i := currPatternLength; i < length; i += currPatternLength {
			subStr := s[i : i+currPatternLength]
			if currPattern != subStr {
				hasPattern = false
			}
		}

		// If we haven't set false once, we've got the same pattern throughout the string
		if hasPattern {
			return hasPattern
		}
	}

	return false
}
