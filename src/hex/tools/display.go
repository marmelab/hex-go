package tools

import (
	"fmt"
	HexGrid "hex/grid"
)

func DisplayStonesAsMatrix(grid HexGrid.Grid) {
	width := grid.Width
	lines := 0
	for i, stone := range grid.Stones {
		if i%width == 0 {
			fmt.Println()
			lines ++
			for y := 0; y < lines; y++ {
				fmt.Print(" ")
			}
		}
		fmt.Print(stone.Player, " ")
	}
}
