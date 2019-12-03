package main

import (
	"errors"
	"fmt"
)

// @todo: Move this file in his own package

const EMPTY = 0
const PLAYER1 = 1
const PLAYER2 = 2

func getDirections(player int) [6][2]int {
	return [6][2]int{{0, -1}, {1, -1}, {1, 0}, {-1, 0}, {-1, 1}, {1, 1}}
}

func buildFromMatrix(matrix [][]int) []stone {

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

func getNeighborsForStone(stones []stone, stone stone) [6]node {
	directions := getDirections(stone.player)
	neighbors := [6]node{}

	for i, direction := range directions {
		xNeighbor := stone.x + direction[0]
		yNeighbor := stone.y + direction[1]

		if neighborStone, err := loadStoneByCoord(stones, xNeighbor, yNeighbor); err != nil {
			fmt.Println("Error occured: ", err)
		} else {
			distance := getDistanceBetweenTwoStones(stone, neighborStone)
			if distance >= 0 {
				neighbors[i] = *NewNode(xNeighbor, yNeighbor, distance)
			}
		}
	}

	return neighbors
}

func loadStoneByCoord(stones []stone, x int, y int) (stone, error) {
	for _, stone := range stones {
		if stone.x == x && stone.y == y {
			return stone, nil
		}
	}
	return stone{}, errors.New("can't find stone at this position")
}

func getDistanceBetweenTwoStones(stone1 stone, stone2 stone) int {
	if stone2.player == EMPTY {
		return 1
	} else if stone1.player == stone2.player {
		return 0
	} else {
		return -1
	}
}
