package game

import (
	"bufio"
	"os"
	"fmt"
)


type GameHandler struct {
	boardHandler *boardHandler
}

func (gameHandler *GameHandler) Play() {

	gameHandler.boardHandler = BoardHandler()
	score := 0

	scanner := bufio.NewScanner(os.Stdin)

	render(gameHandler.boardHandler.board, score)

	for {
		fmt.Println("What direction now?")
		scanner.Scan()
		input := scanner.Text()

		direction := mapDirection(input)

		board, moveScore, changed := gameHandler.boardHandler.move(direction)

		if changed {
			score += moveScore
			render(board, score)
		}
	}


}

func mapDirection(input string) Direction{

	var direction Direction

	switch input{
	case "w":
		direction = Up
	case "a":
		direction = Left
	case "s":
		direction = Down
	case "d":
		direction = Right

	}

	return direction

}
