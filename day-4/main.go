package day4

import (
	"fmt"
	"robinlanderloos/aoc2025/io"
)

func Day4() {
	solveP2("day-4/input.txt")
}

func solveP2(path string) {
	grid := getGrid(path)
	removedRolls := 0

	noRollsRemoved := false
	for noRollsRemoved == false {
		rollsRemoved := 0
		for i := range grid {
			for j := range grid[i] {
				// Only check rolls of paper
				if grid[i][j] != "@" {
					continue
				}
				paperNeighboursCount := getNeighbouringRollsOfPaper(grid, i, j)
				if paperNeighboursCount < 4 {
					rollsRemoved++
					grid[i][j] = "."
				}
			}
		}

		removedRolls += rollsRemoved
		if rollsRemoved == 0 {
			noRollsRemoved = true
		}
	}

	fmt.Print(removedRolls)
}

func solveP1(path string) {
	grid := getGrid(path)
	accessibleRolls := 0

	for i := range grid {
		for j := range grid[i] {
			// Only check rolls of paper
			if grid[i][j] != "@" {
				continue
			}
			paperNeighboursCount := getNeighbouringRollsOfPaper(grid, i, j)
			if paperNeighboursCount < 4 {
				accessibleRolls++
			}
		}
	}

	fmt.Print(accessibleRolls)
}

func getNeighbouringRollsOfPaper(grid [][]string, x, y int) int {
	count := 0

	lowX := 0
	highX := 0
	lowY := 0
	highY := 0

	if x > 0 {
		lowX = -1
	}

	if x < len(grid[0])-1 {
		highX = 1
	}

	if y > 0 {
		lowY = -1
	}

	if y < len(grid[0])-1 {
		highY = 1
	}

	for i := lowX; i <= highX; i++ {
		for j := lowY; j <= highY; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if grid[x+i][y+j] == "@" {
				count++
			}
		}
	}

	return count
}

func getGrid(path string) [][]string {
	lines := io.ReadLines(path)
	grid := make([][]string, len(lines))

	for i, line := range lines {
		lineRunes := make([]string, len(line))
		for i, c := range line {
			lineRunes[i] = string(c)
		}
		grid[i] = lineRunes
	}

	return grid
}
