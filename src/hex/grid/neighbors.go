package grid

type Neighbor struct {
	Stone    Stone
	Distance int
}

func NewNeighbor(stone Stone, distance int) *Neighbor {
	return &Neighbor{stone, distance}
}

func GetNeighborsForStone(grid Grid, stone Stone, player int) []Neighbor {
	directions := getDirections()
	var neighbors []Neighbor

	for _, direction := range directions {

		xNeighbor := stone.x + direction[0]
		yNeighbor := stone.y + direction[1]

		if neighborStone, err := loadStoneByCoord(grid, xNeighbor, yNeighbor); err == nil {
			distance := getDistanceByTypeOfStone(player, neighborStone)
			if distance >= 0 {
				neighbors = append(neighbors, *NewNeighbor(neighborStone, distance))
			}
		}
	}

	return neighbors
}
