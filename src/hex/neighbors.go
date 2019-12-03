package main

import "fmt"

type neighbor struct {
	name     string
	distance int
}

func NewNeighbor(x int, y int, distance int) *neighbor {
	return &neighbor{fmt.Sprintf("%d,%d", x, y), distance}
}

func getNeighborsForStone(stones []stone, stone stone) []neighbor {
	directions := getDirections(stone.player)
	directionCount := len(directions)

	neighbors := make([]neighbor, directionCount)
	neighborsCount := directionCount

	for _, direction := range directions {
		xNeighbor := stone.x + direction[0]
		yNeighbor := stone.y + direction[1]

		if neighborStone, err := loadStoneByCoord(stones, xNeighbor, yNeighbor); err == nil {

			distance := getDistanceBetweenTwoNeighborsStones(stone, neighborStone)
			if distance >= 0 {
				neighbors = append(neighbors, *NewNeighbor(xNeighbor, yNeighbor, distance))
			} else {
				neighborsCount--
			}
		}
	}
	return neighbors[directionCount:]
}

func getAllNeighborsForStones(stones []stone) [][]neighbor {
	neighbors := make([][]neighbor, len(stones))
	neighbors = neighbors[:0]

	for _, stone := range stones {
		neighbors = append(neighbors, getNeighborsForStone(stones, stone))
	}

	return neighbors
}
