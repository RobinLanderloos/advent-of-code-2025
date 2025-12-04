package day3

import (
	"fmt"
	"log"
	"robinlanderloos/aoc2025/io"
	"strconv"
	"strings"
)

type getHighestJoltageFunc func(string) int

func Main() {
	solve("day-3/input.txt", getHighestJoltageForLineP2)
}

func solve(path string, gethighgetHighestJoltageFunc getHighestJoltageFunc) {
	totalJoltage := 0
	for line := range io.EnumerateFile(path) {
		joltage := gethighgetHighestJoltageFunc(line)
		fmt.Printf("Found joltage %d for line %s\r\n", joltage, line)
		totalJoltage += joltage
	}
	fmt.Printf("Found total joltage: %d", totalJoltage)
}

func getHighestJoltageForLineP2(line string) int {
	lineLength := len(line)

	joltageStr := make([]string, 12)
	innerIndex := 0
	for i := range 12 {
		highest := 0
		highestIndex := 0
		for j := innerIndex; j <= lineLength-(12-i); j++ {
			charInt := int(line[j] - '0')
			if charInt > highest {
				highest = charInt
				highestIndex = j
			}
		}
		joltageStr[i] = strconv.Itoa(highest)
		innerIndex = highestIndex + 1
		highest = 0
	}

	joltageInt, err := strconv.Atoi(strings.Join(joltageStr, ""))

	if err != nil {
		log.Fatalf("an error occurred while converting the joltage: %s", err.Error())
	}

	return joltageInt
}

func getHighestJoltageForLineP1(line string) int {
	lineLength := len(line)

	highestLeft := 0
	highestRight := 0
	joltage := make([]string, 2)
	for i, c := range line {
		charInt := int(c - '0')
		if charInt > highestLeft && i < lineLength-1 {
			highestLeft = charInt
			joltage[0] = string(c)
			highestRight = 0
			joltage[1] = "0"
		} else if charInt > highestRight {
			highestRight = charInt
			joltage[1] = string(c)
		}
	}

	joltageInt, err := strconv.Atoi(strings.Join(joltage, ""))

	if err != nil {
		log.Fatalf("an error occurred while converting the joltage: %s", err.Error())
	}

	return joltageInt
}
