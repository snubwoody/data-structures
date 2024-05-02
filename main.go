package main

import "fmt"

/*
	Rules:
	1. All end nodes have to have the same height
	2. Lower values are pushed to the left, higher to the right

*/

func main() {
	btree := BinaryTree{nodes: nil}
	btree.insert(20)
	fmt.Println(btree)
}

//TODO make this an array of 2 instead of a slice
type BTNode struct {
	value int
	left  *BTNode
	right *BTNode
}

type BinaryTree struct {
	nodes []*BTNode
}

func (btree *BinaryTree) insert(value int) {
	node := BTNode{
		value: value,
		left:  nil,
		right: nil,
	}
	btree.nodes = append(btree.nodes, &node)
}
