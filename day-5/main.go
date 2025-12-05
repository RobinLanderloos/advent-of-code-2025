package day5

import (
	"fmt"
	"log"
	"regexp"
	"robinlanderloos/aoc2025/io"
	"strconv"
)

var rangeRegex *regexp.Regexp

func Main() {
	rangeRegex = regexp.MustCompile(`(\d+)-(\d+)`)
	solve("day-5/input.txt")
}

func solve(path string) {
	lines := io.ReadLines(path)
	freshIds := make(map[int]bool)
	freshCount := 0

	checkingRanges := true
	for _, line := range lines {
		if line == "" {
			checkingRanges = false
			continue
		}

		if checkingRanges {
			fmt.Printf("Adding range '%s' to map", line)
			addRangesFromLineToMap(line, freshIds)
		} else {
			curr, err := strconv.Atoi(line)
			if err != nil {
				log.Fatalf("could not convert %s to integer", line)
			}

			fmt.Printf("Checking if id '%d' in range", curr)
			if freshIds[curr] {
				freshCount++
			}
		}
	}

	fmt.Printf("Found %d fresh ids", freshCount)
}

func addRangesFromLineToMap(line string, ranges map[int]bool) {
	matches := rangeRegex.FindStringSubmatch(line)
	low, err := strconv.Atoi(matches[1])
	if err != nil {
		log.Fatalf("an error occurred while converting the lower bound %s", err.Error())
	}

	high, err := strconv.Atoi(matches[2])
	if err != nil {
		log.Fatalf("an error occurred while converting the higher bound %s", err.Error())
	}

	for i := low; i <= high; i++ {
		fmt.Printf("Setting freshid [%d]", i)
		ranges[i] = true
	}
}
