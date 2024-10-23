package wordlegrid

import "fmt"

var emojiMap = map[CellType]string{
	AbsentCellType:  "⬜",
	PresentCellType: "🟨",
	CorrectCellType: "🟩",
}

var darkEmojiMap = map[CellType]string{
	AbsentCellType:  "⬛",
	PresentCellType: "🟨",
	CorrectCellType: "🟩",
}

var contrastEmojiMap = map[CellType]string{
	AbsentCellType:  "⬜",
	PresentCellType: "🟦",
	CorrectCellType: "🟧",
}

var darkContrastEmojiMap = map[CellType]string{
	AbsentCellType:  "⬛",
	PresentCellType: "🟦",
	CorrectCellType: "🟧",
}

func (g *Grid) AsEmoji(dark, contrast bool) (string, error) {
	m := emojiMap
	switch {
	case dark && !contrast:
		m = darkEmojiMap
	case !dark && contrast:
		m = contrastEmojiMap
	case dark && contrast:
		m = darkContrastEmojiMap
	}

	var board string
	for y := 0; y < 6; y++ {
		for x := 0; x < 5; x++ {
			t := g[y][x].Type
			if t == EmptyCellType {
				return "", fmt.Errorf("empty cells are not allowed")
			}
			board += m[t]
		}
		board += "\n"
	}
	return board, nil
}
