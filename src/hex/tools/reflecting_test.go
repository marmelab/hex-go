package tools

import (
	"hex/grid"
	"reflect"
	"testing"
)

func TestAUserCanCheckTypeOfStruct(t *testing.T) {

	stone := grid.NewStone(1, 0, 0, grid.Player1)
	if !IsInstanceOf(stone, &grid.Stone{}) {
		t.Errorf("got %s; want &grid.Stone{}", reflect.TypeOf(stone))
	}

}
