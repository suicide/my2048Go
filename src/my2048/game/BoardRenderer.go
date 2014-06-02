package game

import "fmt"

const separator = "-------------------------------------"

func render(board *board, score int) {

	fmt.Println(separator)

	for x := 0; x < BoardSize; x++ {

		fmt.Printf("| %6d | %6d | %6d | %6d |\n", board.Get(x, 0).Value(), board.Get(x, 1).Value(), board.Get(x, 2).Value(), board.Get(x, 3).Value())
		fmt.Println(separator)

	}

	fmt.Println("Score: ", score)

}
