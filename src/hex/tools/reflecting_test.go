package tools

import (
	HexGrid "hex/grid"
	"reflect"
	"testing"
)

func TestAUserCanCheckTypeOfStruct(t *testing.T) {

	stone := HexGrid.NewStone(1, 0, 0, HexGrid.Player1)
	if !IsInstanceOf(stone, &HexGrid.Stone{}) {
		t.Errorf("got %s; want &grid.Stone{}", reflect.TypeOf(stone))
	}

}
