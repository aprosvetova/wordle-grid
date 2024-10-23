package wordlegrid

import (
	"fmt"
)

func ParseBoardState(attempts []string, solution string) (*Grid, error) {
	if len(solution) != 5 {
		return nil, fmt.Errorf("solution must be 5 characters long")
	}
	g := NewGrid()
	for y, attempt := range attempts {
		if len(attempt) != 5 {
			return nil, fmt.Errorf("attempt must be 5 characters long")
		}
		row := evaluateRow(attempt, solution)
		for x, cell := range row {
			if err := g.SetCell(x+1, y+1, cell, rune(attempt[x])); err != nil {
				return nil, err
			}
		}
	}
	return g, nil
}

func evaluateRow(guess, solution string) [5]CellType {
	var result [5]CellType
	var solutionRemaining [5]bool

	for i := 0; i < 5; i++ {
		if guess[i] == solution[i] {
			result[i] = CorrectCellType
			solutionRemaining[i] = false
		} else {
			solutionRemaining[i] = true
		}
	}

	for i := 0; i < 5; i++ {
		if result[i] == CorrectCellType {
			continue
		}
		for j := 0; j < 5; j++ {
			if solutionRemaining[j] && guess[i] == solution[j] {
				result[i] = PresentCellType
				solutionRemaining[j] = false
				break
			}
		}
	}

	for i := 0; i < 5; i++ {
		if result[i] == EmptyCellType {
			result[i] = AbsentCellType
		}
	}

	return result
}
