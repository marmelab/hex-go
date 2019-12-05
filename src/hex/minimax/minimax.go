package minimax

import (
	"hex/grid"
)

func getAllPossibleGrids(stones []grid.Stone, player int) [][]grid.Stone {
	var gridPossibilities [][]grid.Stone
	size := len(stones) * len(stones)

	for i, stone := range stones {

		if stone.Player == grid.Empty {
			stonesCopy := append(make([]grid.Stone, 0, size), stones...)
			stonesCopy[i].Player = player

			gridPossibilities = append(gridPossibilities, stonesCopy)
		}
	}

	return gridPossibilities
}
