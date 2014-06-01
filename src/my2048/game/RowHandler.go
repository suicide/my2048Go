package game

import (
	"reflect"
)

func foldRow(row []Cell) (newRow []Cell, score int, changed bool) {

	newRow = make([]Cell, 0, BoardSize)
	var previousCell Cell = nil
	score = 0

	for _, cell := range row {
		if cell.Value() == 0 {
			continue
		}

		if cell == previousCell {
			//merge
			newCell := GetCell(cell.Value() * 2)
			newRow = append(newRow, newCell)
			score += newCell.Value()

			previousCell = nil
		} else {
			if previousCell != nil {
				newRow = append(newRow, previousCell)
			}
			previousCell = cell
		}
	}

	if previousCell != nil {
		newRow = append(newRow, previousCell)
	}

	length := len(newRow)
	for i := 0; i < BoardSize - length; i++ {
		newRow = append(newRow, GetCell(0))
	}

	changed = !reflect.DeepEqual(row, newRow)

	return
}
