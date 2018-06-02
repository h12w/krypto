package main

type (
	node struct {
		value    int
		parent   int8
		children [2]int8
	}
	nodes []node
	op    int
)

const (
	plus = iota
	minus
	times
	divide
	maxOp
)
