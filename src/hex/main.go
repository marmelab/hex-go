package main

import (
	"bufio"
	"fmt"
	Graph "hex/graph"
	HexGrid "hex/grid"
	"hex/state"
	"hex/tools"
	"os"
)

func main() {

	demoNumber := ""

	fmt.Println("-------------------------------")
	fmt.Println("         AI Golang demo        ")
	fmt.Println("-------------------------------", "")
	fmt.Println("", "Demo number : ")

	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	demoNumber = input.Text()
	fmt.Println("\n", "")

	if demoNumber == "1" {
		demo1A()
		demo1B()
		demo1C()
	}

	if demoNumber == "2" {
		demo2()
	}

	if demoNumber == "3" {
		demo3a()
		demo3b()
	}
}

func demo1A() {
	fmt.Println("Demo number 1 : Detect victory, lost and victory in 1 step ", "Victory detection :")
	matrix := [][]int{
		{0, 1, 0, 0, 2},
		{0, 1, 0, 0, 2},
		{2, 1, 0, 0, 2},
		{0, 1, 2, 2, 0},
		{0, 1, 0, 0, 0},
	}
	grid := HexGrid.GetGridFromMatrix(matrix)
	tools.DisplayStonesAsMatrix(grid)

	graph, endVertexId := Graph.BootstrapGraphPlayer1(matrix)

	bestPath, _ := graph.Shortest(Graph.StartVertexId, endVertexId)
	fmt.Println("")
	fmt.Println("Best path : ", bestPath)
	fmt.Println("is won :", state.IsWon(bestPath))
}

func demo1B() {
	matrix := [][]int{
		{0, 1, 0, 0, 2},
		{0, 1, 0, 0, 2},
		{2, 0, 0, 0, 2},
		{0, 1, 2, 2, 0},
		{0, 1, 0, 0, 0},
	}
	grid := HexGrid.GetGridFromMatrix(matrix)
	tools.DisplayStonesAsMatrix(grid)

	graph, endVertexId := Graph.BootstrapGraphPlayer1(matrix)

	bestPath, _ := graph.Shortest(Graph.StartVertexId, endVertexId)
	fmt.Println("")
	fmt.Println("Best path : ", bestPath)
	fmt.Println("Can win in 1 turn :", state.CanWinInOneTurn(bestPath))
}

func demo1C() {
	matrix := [][]int{
		{0, 1, 0, 0, 2},
		{0, 1, 0, 0, 2},
		{2, 2, 2, 0, 2},
		{0, 1, 2, 2, 0},
		{0, 1, 0, 0, 0},
	}
	grid := HexGrid.GetGridFromMatrix(matrix)
	tools.DisplayStonesAsMatrix(grid)

	graph, endVertexId := Graph.BootstrapGraphPlayer2(matrix)

	bestPath, _ := graph.Shortest(Graph.StartVertexId, endVertexId)
	fmt.Println("")
	fmt.Println("Player 2 Best path : ", bestPath)
	fmt.Println("Loose :", state.IsWon(bestPath))
}

func demo2() {
	fmt.Println("Demo number 2 : Closest to victory")
	matrix := [][]int{
		{0, 1, 0, 0, 2},
		{0, 1, 0, 0, 2},
		{2, 0, 0, 0, 2},
		{0, 1, 2, 2, 0},
		{0, 1, 0, 0, 0},
	}
	grid := HexGrid.GetGridFromMatrix(matrix)
	tools.DisplayStonesAsMatrix(grid)

	graph, endVertexId := Graph.BootstrapGraphPlayer1(matrix)
	graph2, endVertexId := Graph.BootstrapGraphPlayer2(matrix)

	bestPath, _ := graph.Shortest(Graph.StartVertexId, endVertexId)
	bestPath2, _ := graph2.Shortest(Graph.StartVertexId, endVertexId)

	fmt.Println("")
	fmt.Println("Best path P1 : ", bestPath)
	fmt.Println("Best path P2 : ", bestPath2)
	fmt.Println("Player is closest :", state.IsClosestAsOpponent(bestPath, bestPath2))
}

func demo3a() {
	fmt.Println("Demo number 3 : Advice for next move")
	matrix := [][]int{
		{0, 1, 0, 0, 2},
		{0, 1, 0, 0, 2},
		{2, 1, 0, 0, 2},
		{0, 0, 2, 2, 0},
		{0, 1, 0, 0, 0},
	}
	grid := HexGrid.GetGridFromMatrix(matrix)
	tools.DisplayStonesAsMatrix(grid)

	graph, endVertexId := Graph.BootstrapGraphPlayer1(matrix)

	bestPath, _ := graph.Shortest(Graph.StartVertexId, endVertexId)

	fmt.Println("")
	fmt.Println("Graph : ", bestPath)
	advice := state.AdviceForNextMove(bestPath, graph)
	fmt.Println("Advice (ID of the stone) :", advice)

	fmt.Println("Grid with advice :")

	grid.Stones[advice-1].Player = HexGrid.Player1
	tools.DisplayStonesAsMatrix(grid)
}

func demo3b() {
	fmt.Println("Demo number 3 : Advice for next move")
	matrix := [][]int{
		{0, 1, 0, 0, 2},
		{0, 1, 0, 0, 2},
		{2, 2, 0, 0, 2},
		{0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0},
	}
	grid := HexGrid.GetGridFromMatrix(matrix)
	tools.DisplayStonesAsMatrix(grid)

	graph, endVertexId := Graph.BootstrapGraphPlayer2(matrix)

	bestPath, _ := graph.Shortest(Graph.StartVertexId, endVertexId)

	fmt.Println("")
	fmt.Println("Graph : ", bestPath)
	advice := state.AdviceForNextMove(bestPath, graph)
	fmt.Println("Advice P2 (ID of the stone) :", advice)

	fmt.Println("Grid with advice :")

	grid.Stones[advice-1].Player = HexGrid.Player2
	tools.DisplayStonesAsMatrix(grid)
}
