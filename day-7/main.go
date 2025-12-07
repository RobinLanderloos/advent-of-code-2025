package day7

import (
	"fmt"
	"robinlanderloos/aoc2025/io"
)

func Main() {
	solveP2("day-7/input.txt")
}

func solveP2(path string) {
	grid := getGridP2(path)

	total := traverseAndGetTimelines(grid)

	fmt.Printf("Grid: \r\n%v\r\n", grid)
	fmt.Printf("timelines: %d", total)
}

func traverseAndGetTimelines(grid [][]int) int {
	gridRows := len(grid)
	gridColumns := len(grid[0])

	for row := 0; row < gridRows-1; row++ {
		for column := 0; column < gridColumns; column++ {
			curr := grid[row][column]
			if curr == -1 {
				continue
			}

			next := grid[row+1][column]

			// Splitter
			if next == -1 {
				if column > 0 {
					grid[row+1][column-1] += curr
				}
				if column < gridColumns {
					grid[row+1][column+1] += curr
				}
			} else {
				grid[row+1][column] += curr
			}
		}
	}

	total := 0
	for _, value := range grid[gridRows-1] {
		total += value
	}

	return total
}

func solveP1(path string) {
	grid := getGrid(path)

	splits := traverseAndGetBeamSplits(grid)

	fmt.Printf("Splits: %d", splits)
}

func traverseAndGetBeamSplits(grid [][]string) int {
	gridRows := len(grid)
	gridColumns := len(grid[0])
	splits := 0

	for row := 0; row < gridRows-1; row++ {
		for column := 0; column < gridColumns; column++ {
			if grid[row][column] == "S" || grid[row][column] == "|" {
				// No splitter, update grid
				if grid[row+1][column] == "." {
					grid[row+1][column] = "|"
					continue
				}
				if grid[row+1][column] == "^" {
					splits++
					if column > 0 {
						grid[row+1][column-1] = "|"
					}
					if column < gridColumns {
						grid[row+1][column+1] = "|"
					}
				}
			}
		}
	}

	return splits
}

func printGrid(grid [][]int) {
	gridRows := len(grid)
	gridColumns := len(grid[0])

	for row := 0; row < gridRows-1; row++ {
		for column := 0; column < gridColumns; column++ {
			curr := grid[row][column]
			if curr >= 0 && curr < 10 {
				fmt.Printf(" %d", grid[row][column])
			} else {

				fmt.Print(grid[row][column])
			}
		}
		fmt.Println()
	}
}

func getGridP2(path string) [][]int {
	lines := io.ReadLines(path)
	result := make([][]int, len(lines))

	for row, line := range lines {
		lineSlice := make([]int, len(line))
		for column, char := range line {
			charStr := string(char)
			value := 0
			if charStr == "S" {
				value = 1
			}
			if charStr == "^" {
				value = -1
			}
			lineSlice[column] = value
		}
		result[row] = lineSlice
	}

	return result
}

func getGrid(path string) [][]string {
	lines := io.ReadLines(path)
	result := make([][]string, len(lines))

	for row, line := range lines {
		lineSlice := make([]string, len(line))
		for column, char := range line {
			charStr := string(char)
			lineSlice[column] = charStr
		}
		result[row] = lineSlice
	}

	return result
}
