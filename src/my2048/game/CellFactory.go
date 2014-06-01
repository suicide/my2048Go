package game

func GetCell(value int) Cell {

	return valueCell{value}

}

type valueCell struct {
	value int
}

func (cell valueCell) Value() int {
	return cell.value
}
