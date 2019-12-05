package grid

type Stone struct {
	Id     int
	x      int
	y      int
	Player int
}

func NewStone(id int, x int, y int, player int) *Stone {
	return &Stone{Id: id, x: x, y: y, Player: player}
}
