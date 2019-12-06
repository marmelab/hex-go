package tools

import (
	"fmt"
	HexGrid "hex/grid"
)

func DisplayStonesAsMatrix(stones []HexGrid.Stone, width int) {
	for i, stone := range stones {
		if i%width == 0 {
			fmt.Println()
		}
		fmt.Print(stone.Player, " ")
	}
}
