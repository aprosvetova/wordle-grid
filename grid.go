package wordlegrid

import (
	"fmt"
	"unicode"
)

type CellType int

const (
	EmptyCellType CellType = iota
	AbsentCellType
	PresentCellType
	CorrectCellType
)

type Cell struct {
	Type   CellType
	Letter rune
}

type Grid [6][5]Cell

func NewGrid() *Grid {
	return &Grid{}
}

func (g *Grid) SetCell(x, y int, t CellType, letter rune) error {
	if x < 1 || x > 5 {
		return fmt.Errorf("column out of bounds")
	}
	if y < 1 || y > 6 {
		return fmt.Errorf("row out of bounds")
	}
	if t != EmptyCellType && letter == 0 {
		return fmt.Errorf("letter required for non-empty cells")
	}
	if (letter < 'a' || letter > 'z') && (letter < 'A' || letter > 'Z') {
		return fmt.Errorf("invalid letter")
	}

	letter = unicode.ToUpper(letter)
	g[y-1][x-1] = Cell{t, letter}

	return nil
}
