package day5

import (
	"fmt"
	"log"
	"regexp"
	"robinlanderloos/aoc2025/io"
	"sort"
	"strconv"
)

var rangeRegex *regexp.Regexp

func Main() {
	rangeRegex = regexp.MustCompile(`(\d+)-(\d+)`)
	solveP2("day-5/input.txt")
}

func solveP2(path string) {
	ranges, _ := parseInput(path)
	totalFreshIds := 0

	for _, r := range ranges {
		freshIdsInRange := r[1] - r[0] + 1
		totalFreshIds += freshIdsInRange
	}

	fmt.Printf("Total fresh ids: %d", totalFreshIds)
}

func solveP1(path string) {
	freshCount := 0

	ranges, ids := parseInput(path)

	for _, id := range ids {
		for _, r := range ranges {
			if id >= r[0] && id <= r[1] {
				freshCount++
				continue
			}
		}
	}

	fmt.Printf("Found %d fresh ids", freshCount)
}

func parseInput(path string) ([][]int, []int) {
	ranges := make([][]int, 0)
	ids := make([]int, 0)

	checkingRanges := true
	for line := range io.EnumerateFile(path) {
		// We encountered the middle point of the input
		if line == "" {
			checkingRanges = false
			continue
		}

		if checkingRanges {
			low, high := getRangeFromLine(line)
			r := make([]int, 2)
			r[0] = low
			r[1] = high
			ranges = append(ranges, r)
		} else {
			curr, err := strconv.Atoi(line)
			if err != nil {
				log.Fatalf("could not convert %s to integer", line)
			}
			ids = append(ids, curr)
		}
	}

	ranges = mergeRanges(ranges)

	return ranges, ids
}

func mergeRanges(ranges [][]int) [][]int {
	result := make([][]int, 0)
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	curr := ranges[0]
	for i := 1; i < len(ranges); i++ {
		if rangeOverlaps(curr, ranges[i]) {
			curr = mergeRange(curr, ranges[i])
		} else {
			result = append(result, curr)
			curr = ranges[i]
		}
	}

	if result[len(result)-1][0] != curr[0] && result[len(result)-1][1] != curr[1] {
		result = append(result, curr)
	}

	return result
}

func mergeRange(curr, other []int) []int {
	merged := make([]int, 2)
	lowest := 0
	highest := 0
	if curr[0] <= other[0] {
		lowest = curr[0]
	} else {
		lowest = other[0]
	}

	if curr[1] >= other[1] {
		highest = curr[1]
	} else {
		highest = other[1]
	}

	merged[0] = lowest
	merged[1] = highest

	return merged
}

func rangeOverlaps(curr, other []int) bool {
	return other[0] <= curr[1]+1
}

func getRangeFromLine(line string) (int, int) {
	matches := rangeRegex.FindStringSubmatch(line)

	low, _ := strconv.Atoi(matches[1])
	high, _ := strconv.Atoi(matches[2])

	return low, high
}
