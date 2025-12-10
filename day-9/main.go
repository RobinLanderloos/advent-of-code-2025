package day9

import (
	"fmt"
	"robinlanderloos/aoc2025/io"
	"robinlanderloos/aoc2025/types"
	"strings"
)

func Main() {
	// solveP1("day-9/input.txt")
	solveP2("day-9/example-input.txt")
}

func solveP2(path string) {
	coordinates := []*types.Coordinate{}
	redTiles := map[types.Coordinate]bool{}

	for _, coord := range coordinates {
		redTiles[*coord] = true
	}
}

func solveP1(path string) {
	coordinates := []*types.Coordinate{}
	for line := range io.EnumerateFile(path) {
		coords := strings.Split(line, ",")
		coord := types.NewCoordinateFromStr(coords[1], coords[0])
		coordinates = append(coordinates, coord)
	}

	largestRect := 0
	for i := 0; i < len(coordinates)-1; i++ {
		for j := i + 1; j < len(coordinates); j++ {
			size := getRectangleSize(coordinates[i], coordinates[j])
			if size > largestRect {
				largestRect = size
			}
		}
	}

	fmt.Printf("Largest size found: %v", largestRect)
}

func getRectangleSize(start, end *types.Coordinate) int {
	xDiff := start.Col - end.Col
	yDiff := start.Row - end.Row
	sizeX := max(xDiff, -xDiff) + 1
	sizeY := max(yDiff, -yDiff) + 1

	return sizeX * sizeY
}
