package wordlegrid

import (
	"fmt"
	"strings"
)

var svgMap = map[CellType]string{
	EmptyCellType:   "empty",
	AbsentCellType:  "absent",
	PresentCellType: "present",
	CorrectCellType: "correct",
}

func formatCellID(x, y int, t CellType) string {
	return fmt.Sprintf("%d-%d-%s", y, x, svgMap[t])
}

func formatLetterID(x, y int, t CellType) string {
	id := fmt.Sprintf("%d-%d-letter", y, x)
	if t == AbsentCellType {
		id += "-absent"
	}
	return id
}

func (g *Grid) AsSVG() string {
	lines := strings.Split(svgTemplate, "\n")
	for y := 0; y < 6; y++ {
		for x := 0; x < 5; x++ {
			t := g[y][x].Type
			l := g[y][x].Letter
			cellID := formatCellID(x+1, y+1, t)
			for i, line := range lines {
				if strings.Contains(line, cellID) {
					lines[i] = strings.ReplaceAll(line, "display=\"none\"", "")
					break
				}
			}
			if l == 0 {
				continue
			}
			letterID := formatLetterID(x+1, y+1, t)
			for i, line := range lines {
				if strings.Contains(line, letterID) {
					line = strings.ReplaceAll(line, "display=\"none\"", "")
					lines[i] = strings.ReplaceAll(line, "$", string(l))
					break
				}
			}
		}
	}
	return strings.Join(lines, "\n")
}
