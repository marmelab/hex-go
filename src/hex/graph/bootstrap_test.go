package graph

import (
	"github.com/RyanCarrier/dijkstra"
	"hex/tools"
	"reflect"
	"testing"
)

func TestBoostrappingOfGraph(t *testing.T) {
	matrix := [][]int{
		{0, 1},
		{2, 1},
	}

	got, endVertexId := BootstrapGraphPlayer1(matrix)

	endVertexIdexpected := 5

	if !tools.IsInstanceOf(got, &dijkstra.Graph{}) {
		t.Errorf("got %s; want dijkstra.Graph{}", reflect.TypeOf(got))
	}

	if endVertexId != endVertexIdexpected {
		t.Errorf("got %d; want %d", endVertexId, endVertexIdexpected)
	}
}
