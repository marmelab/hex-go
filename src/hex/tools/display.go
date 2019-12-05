package tools

import (
	"fmt"
	"hex/grid"
)

func DisplayStonesAsMatrix(stones []grid.Stone, width int) {
	for i, stone := range stones {
		if i%width == 0 {
			fmt.Println()
		}
		fmt.Print(stone.Player, " ")
	}
}
