package main

type stone struct {
	x      int
	y      int
	player int
}

func NewStone(x int, y int, player int) *stone {

	return &stone{x: x, y: y, player: player}
}

// Array filter with callback
// From here : https://stackoverflow.com/questions/37562873/most-idiomatic-way-to-select-elements-from-an-array-in-golang
func filterStones(stones []stone, test func(stone) bool) (ret []stone) {
	for _, stone := range stones {
		if test(stone) {
			ret = append(ret, stone)
		}
	}
	return
}
