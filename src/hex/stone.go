package main

type Stone struct {
	id     int
	x      int
	y      int
	player int
}

func NewStone(id int, x int, y int, player int) *Stone {
	return &Stone{id: id, x: x, y: y, player: player}
}
