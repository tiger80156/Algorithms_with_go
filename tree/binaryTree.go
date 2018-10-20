package main

import (
	"fmt"
	"os"
)

type node struct {
	key    int
	parent *node
	left   *node
	right  *node
}

type tree struct {
	root *node
}

func main() {

	node1 := newNode(5)
	node2 := newNode(3)
	node3 := newNode(7)
	node4 := newNode(2)
	node5 := newNode(4)
	node6 := newNode(6)
	node7 := newNode(9)

	node1.addChild(&node2, &node3)
	node2.addChild(&node4, &node5)
	node3.addChild(&node6, &node7)

	t := tree{
		root: &node1,
	}

	t.root.traversal()
	sn, err := t.root.search(3)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(sn.successor())

}

func (n node) traversal() {

	if n.left != nil {
		n.left.traversal()
	}

	fmt.Println(n.key)

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

func (n *node) search(key int) (*node, error) {

	if n.key == key {
		return n, nil
	}

	if n.left == nil && n.right == nil {
		return nil, fmt.Errorf("The tree node didn't exist.")
	}

	if key < n.key {
		return n.left.search(key)
	} else {
		return n.right.search(key)
	}

}

func (n *node) print() {
	if n.parent != nil {
		fmt.Println("Parent: ", n.parent.key)
	}
	if n.right != nil {
		fmt.Println("right child: ", n.right.key)
	}
	if n.right != nil {
		fmt.Println("left child: ", n.left.key)
	}
}

func (n *node) successor() *node {
	// Case 1 - the right tree isn't empty
	// Find the left most node
	if n.right != nil {
		return n.right.leftmost()
	}

	// Case 2 - the right tree is empty
	// Find the node ancestor
	y := n.parent
	for y != nil && n == y.right {
		n = y
		y = y.parent
	}

	return y
}

func (n *node) leftmost() *node {
	for n.left != nil {
		n = n.left
	}

	return n
}
