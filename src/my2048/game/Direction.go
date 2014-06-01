package game

type Direction int

const (
	Up Direction = 1 << iota
	Down Direction = 1 << iota
	Left Direction = 1 << iota
	Right Direction = 1 << iota
)
