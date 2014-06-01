package game

import (
	"math/rand"
)

type boardHandler struct {

	board *board

}

func BoardHandler() *boardHandler {
	boardHandler := new(boardHandler)

	boardHandler.board = spawnRandom(spawnRandom(Board()))

	return boardHandler
}

func (boardHandler *boardHandler) move(direction Direction) (board *board, score int, changed bool) {

	newBoard, score, changed := move(boardHandler.board, direction)

	if changed {
		boardHandler.board = spawnRandom(newBoard)
		board = boardHandler.board
	}
	return
}

func move(board *board, direction Direction) (newBoard *board, score int, changed bool) {

	newBoard = Board()

	coordinates := generateCoordinates(direction)

	score = 0
	changed = false

	for i := 0; i < BoardSize; i++ {

		row := make([]Cell, 0, BoardSize)

		for j := 0; j < BoardSize; j++ {
			coordinate := coordinates[i * BoardSize + j]
			cell := board.Get(coordinate.x, coordinate.y)
			row = append(row, cell)
		}

		newRow, rowScore, rowChanged := foldRow(row)

		score += rowScore
		changed = changed || rowChanged

		for j := 0; j < BoardSize; j++ {
			coordinate := coordinates[i * BoardSize + j]
			newBoard.Set(coordinate.x, coordinate.y, newRow[j])
		}

	}

	return

}

func spawnRandom(board *board) *board {

	emptyFields := make([]coordinate, 0, BoardSize * BoardSize)

	for x := 0; x < BoardSize; x++ {
		for y := 0; y < BoardSize; y++ {
			cell := board.Get(x, y)

			if cell.Value() == 0 {
				emptyFields = append(emptyFields, coordinate{x, y})
			}
		}
	}

	index := rand.Intn(len(emptyFields))
	value := (rand.Intn(2) + 1) * 2

	coordinate := emptyFields[index]

	board.Set(coordinate.x, coordinate.y, GetCell(value))

	return board

}
