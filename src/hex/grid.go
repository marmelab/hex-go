package main

import (
	"errors"
)

// @todo: Move this file in his own package

const EMPTY = 0
const PLAYER1 = 1
const PLAYER2 = 2

func getDirections(player int) [6][2]int {
	return [6][2]int{{0, -1}, {1, -1}, {1, 0}, {-1, 0}, {-1, 1}, {1, 1}}
}

func getStonesFromMatrix(matrix [][]int) []stone {

	length := len(matrix)
	stones := make([]stone, length*length)
	stones = stones[:0]

	for y, line := range matrix {
		for x, value := range line {

			// @todo: Refactor this (Make a function ?)
			player := EMPTY
			if value == PLAYER1 || value == PLAYER2 {
				player = value
			}

			stones = append(stones, *NewStone(x, y, player))
		}
	}

	return stones
}

func loadStoneByCoord(stones []stone, x int, y int) (stone, error) {
	for _, stone := range stones {
		if stone.x == x && stone.y == y {
			return stone, nil
		}
	}
	return stone{}, errors.New("can't find stone at this position")
}

func getDistanceBetweenTwoNeighborsStones(stone1 stone, stone2 stone) int {
	if stone2.player == EMPTY {
		return 1
	}

	if stone1.player == stone2.player {
		return 0
	}

	return -1
}
