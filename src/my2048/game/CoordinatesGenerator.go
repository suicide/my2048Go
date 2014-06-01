package game

type coordinate struct {
	x int
	y int
}

func generateCoordinates(direction Direction) []coordinate {
	coordinates := make([]coordinate, 0, BoardSize * BoardSize)

	switch direction {
	case Left:
		for x := 0; x < BoardSize; x++ {
			for y := 0; y < BoardSize; y++ {
				coordinates = append(coordinates, coordinate{x, y})
			}
		}
	case Right:
		for x := 0; x < BoardSize; x++ {
			for y := BoardSize - 1; y >= 0; y-- {
				coordinates = append(coordinates, coordinate{x, y})
			}
		}
	case Up:
		for y := 0; y < BoardSize; y++ {
			for x := 0; x < BoardSize; x++ {
				coordinates = append(coordinates, coordinate{x, y})
			}
		}
	case Down:
		for y := 0; y < BoardSize; y++ {
			for x := BoardSize - 1; x >= 0; x-- {
				coordinates = append(coordinates, coordinate{x, y})
			}
		}
	}

	return coordinates
}
