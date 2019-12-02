package main

import "fmt"

type node struct {
	name     string
	distance int
}

func NewNode(x int, y int, distance int) *node {

	stone := node{fmt.Sprintf("%d,%d", x, y), distance}
	return &stone
}
