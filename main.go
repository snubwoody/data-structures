package main

import "fmt"

/*
	Binary Search Tree

	Rules:
	1. All end nodes have to have the same height
	2. Lower values are pushed to the left, higher to the right

*/

func main() {
	btree := BST{root: nil}
	btree.insert(20)
	fmt.Println(btree)
}

//TODO make this an array of 2 instead of a slice
type BSTNode struct {
	value int
	left  *BSTNode
	right *BSTNode
}

type BST struct {
	root *BSTNode
}

func (btree *BST) insert(value int) {
	if btree.root == nil {
		btree.root = &BSTNode{
			value: value,
			left:  nil,
			right: nil,
		}
		return
	}
}
