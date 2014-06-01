package game

type board struct {
	cells [BoardSize * BoardSize]Cell
}

const BoardSize = 4

func Board() *board {
	board := new(board)

	for i := 0; i < BoardSize * BoardSize; i++ {
		board.cells[i] = GetCell(0)
	}

	return board
}

func (board board) Get(x,y int) Cell {
	return board.cells[x * BoardSize + y]
}

func (board *board) Set(x,y int, cell Cell) {
	board.cells[x * BoardSize + y] = cell
}
