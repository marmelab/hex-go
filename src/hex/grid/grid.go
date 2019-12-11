package grid

import (
	"errors"
	"math"
)

const Empty = 0
const Player1 = 1
const Player2 = 2

const DistanceOwned = 0
const DistanceEmpty = 1
const DistanceOpponent = -1

type Grid struct {
	Stones []Stone
	Width  int
}

func NewGrid(stones []Stone, width int) *Grid {
	return &Grid{Stones: stones, Width: width}
}

func getDirections() [6][2]int {
	return [6][2]int{{0, -1}, {1, -1}, {1, 0}, {-1, 0}, {-1, 1}, {0, 1}}
}

func GetGridFromArray(array []int) Grid {
	length := len(array)
	width := int(math.Sqrt(float64(length)))

	stones := make([]Stone, length)
	stones = stones[:0]

	grid := NewGrid(stones, width)

	y := 0
	x := 0
	count := 0
	for _, value := range array {

		count++
		player := getPlayerValue(value)
		grid.Stones = append(grid.Stones, *NewStone(count, x, y, player))

		if x == width-1 {
			x = 0
			y++
		} else {
			x++
		}
	}

	return *grid
}

func GetGridFromMatrix(matrix [][]int) Grid {

	width := len(matrix)
	stones := make([]Stone, width*width)
	stones = stones[:0]

	grid := NewGrid(stones, width)
	count := 0
	for y, line := range matrix {
		for x, value := range line {
			player := getPlayerValue(value)
			count++
			grid.Stones = append(grid.Stones, *NewStone(count, x, y, player))
		}
	}

	return *grid
}

func getPlayerValue(value int) int {
	if value == Player1 {
		return Player1
	}
	if value == Player2 {
		return Player2
	}

	return Empty
}

func loadStoneByCoord(grid Grid, x int, y int) (Stone, error) {
	for _, stone := range grid.Stones {
		if stone.x == x && stone.y == y {
			return stone, nil
		}
	}
	return Stone{}, errors.New("can't find Stone at this position")
}

func getDistanceByTypeOfStone(player int, stone2 Stone) int {
	if stone2.Player == Empty {
		return DistanceEmpty
	}

	if player == stone2.Player {
		return DistanceOwned
	}

	return DistanceOpponent
}
