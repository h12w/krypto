package main

import (
	"strconv"
)

func (n *node) str(ns nodes, parent op) string {
	if n.children[0] == 0 {
		return strconv.Itoa(n.value)
	}
	op := op(n.value)
	expr := ns[n.children[0]-1].str(ns, op) + op.String() + ns[n.children[1]-1].str(ns, op)
	switch parent {
	case times, divide:
		switch op {
		case plus, minus:
			expr = "(" + expr + ")"
		}
	}
	return expr
}

func (ns nodes) root() *node {
	root := &ns[0]
	for root.parent != 0 {
		root = &ns[root.parent-1]
	}
	return root
}

func (ns nodes) String() string {
	if len(ns) == 0 {
		return ""
	}
	return ns.root().str(ns, plus)
}

func (op op) String() string {
	switch op {
	case plus:
		return "+"
	case minus:
		return "-"
	case times:
		return "*"
	case divide:
		return "/"
	}
	return ""
}
