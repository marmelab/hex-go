package main

const EMPTY = 0
const PLAYER1 = 1
const PLAYER2 = 2

func getDirection(player int) [6][2]int {
	if player == PLAYER1 {
		return [6][2]int{{1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}}
	} else if player == PLAYER2 {
		return [6][2]int{{1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}, {1, 1}}
	}
}

func buildFromMatrix(matrix [][]int) []stone {

	length := len(matrix) * len(matrix)
	game := make([]stone, length)
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

func getNeighborsForStone(game []stone, stone stone) []node {

	directions := getDirection(stone.player)
	neighbors := make([]node, 6)

	for _, direction := range directions {
		xNeighbor := stone.x + direction[0]
		yNeighbor := stone.y + direction[1]

		neighborStone := loadStoneByCoord(game, xNeighbor, yNeighbor)

		distance := getDistanceBetweenTwoStones(stone, neighborStone)

		if distance >= 0 {
			neighbors = append(neighbors, *NewNode(xNeighbor, yNeighbor, distance))
		}
	}
	return neighbors
}

func loadStoneByCoord(game []stone, x int, y int) stone {
	return filterStones(game, func(stone stone) bool { return stone.x == x && stone.y == y })[0]
}

func getDistanceBetweenTwoStones(stone1 stone, stone2 stone) int {
	if stone2.player == EMPTY {
		return 1
	} else if stone1.player == stone2.player {
		return 0
	}
	return -1
}
