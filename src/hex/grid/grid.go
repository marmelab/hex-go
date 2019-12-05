package grid

import (
	"errors"
)

const Empty = 0
const Player1 = 1
const Player2 = 2

const DistanceOwned = 0
const DistanceEmpty = 1
const DistanceOpponent = -1

func getDirections() [6][2]int {
	return [6][2]int{{0, -1}, {1, -1}, {1, 0}, {-1, 0}, {-1, 1}, {0, 1}}
}

func GetStonesFromMatrix(matrix [][]int) []Stone {

	length := len(matrix)
	stones := make([]Stone, length*length)
	stones = stones[:0]

	count := 0
	for y, line := range matrix {
		for x, value := range line {

			// @todo: Refactor this (Make a function ?)
			player := Empty
			if value == Player1 || value == Player2 {
				player = value
			}

			count++
			stones = append(stones, *NewStone(count, x, y, player))
		}
	}

	return stones
}

func loadStoneByCoord(stones []Stone, x int, y int) (Stone, error) {
	for _, stone := range stones {
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
