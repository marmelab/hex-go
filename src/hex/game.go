package main

const EMPTY = 0
const PLAYER1 = 1
const PLAYER2 = 2


func buildFromMatrix(matrix [][]int) []stone {

	length := len(matrix) * len(matrix)
	game := make([]stone, length )
	for x, line := range matrix {
		for y, value := range line {
			player := EMPTY
			if value == PLAYER1 || value == PLAYER2 {
				player = value
			}

			NewStone(x, y, player)
		}
	}

	return game
}
