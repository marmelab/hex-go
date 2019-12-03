package main

type stone struct {
	x      int
	y      int
	player int
}

func NewStone(x int, y int, player int) *stone {
	return &stone{x: x, y: y, player: player}
}

// @todo: Improve this function to reduce memory usage
func (stone stone) getId() int {
	return stone.x + ',' + stone.y
}
