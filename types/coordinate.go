package types

import "strconv"

type Coordinate struct {
	Row, Col int
}

func NewCoordinate(row, col int) *Coordinate {
	return &Coordinate{
		Row: row,
		Col: col,
	}
}

func NewCoordinateFromStr(rowStr, colStr string) *Coordinate {
	row, _ := strconv.Atoi(rowStr)
	col, _ := strconv.Atoi(colStr)

	return NewCoordinate(row, col)
}
