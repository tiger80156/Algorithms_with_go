package main

import (
	"fmt"
)

type node struct {
	key    int
	parent *node
	left   *node
	right  *node
}

func main() {

	node1 := newNode(5)
	node2 := newNode(3)
	node3 := newNode(7)
	node4 := newNode(2)
	node5 := newNode(4)
	node6 := newNode(8)

	node1.addChild(&node2, &node3)
	node2.addChild(&node4, &node5)
	node3.addChild(nil, &node6)
	node1.traversal()
}

func (n node) traversal() {

	fmt.Println(n.key)

	if n.left != nil {
		n.left.traversal()
	}

	if n.right != nil {
		n.right.traversal()
	}
}

func newNode(k int) node {
	newNode := node{
		key:    k,
		parent: nil,
		left:   nil,
		right:  nil,
	}

	return newNode
}

func (n *node) addChild(leftChild *node, rightChild *node) {
	n.left = leftChild
	n.right = rightChild
	if leftChild != nil {
		leftChild.parent = n
	}
	if rightChild != nil {
		rightChild.parent = n
	}
}
