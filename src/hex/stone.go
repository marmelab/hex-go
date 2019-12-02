package main

type stone struct {
	x      int
	y      int
	player int
}

func NewStone(x int, y int, player int) *stone {

	stone := stone{x: x, y: y, player: player}
	return &stone
}
