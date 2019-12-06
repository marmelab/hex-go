package minimax

import (
	HexGrid "hex/grid"
)

func getAllPossibleGrids(originalGrid HexGrid.Grid, player int) []HexGrid.Grid {
	var gridPossibilities []HexGrid.Grid
	width := originalGrid.Width
	size := width * width

	for i, stone := range originalGrid.Stones {
		if stone.Player == HexGrid.Empty {
			stonesCopy := append(make([]HexGrid.Stone, 0, size), originalGrid.Stones...)
			stonesCopy[i].Player = player
			newGrid := HexGrid.NewGrid(stonesCopy, width)
			gridPossibilities = append(gridPossibilities, *newGrid)
		}
	}

	return gridPossibilities
}
