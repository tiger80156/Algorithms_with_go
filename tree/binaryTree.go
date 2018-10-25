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
	t := tree{}

	t.insertion(&node1)
	t.insertion(&node2)
	t.insertion(&node3)
	t.insertion(&node4)
	t.insertion(&node5)
	t.insertion(&node6)
	t.insertion(&node7)

	t.deleteNode(3)
	t.deleteNode(4)
	t.root.traversal()

}

func (n node) traversal() {
	// The inorder traversal of BST
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

// Find minimum
func (n *node) leftmost() *node {
	for n.left != nil {
		n = n.left
	}
	return n
}

// Find Maximum
func (n *node) rightmost() *node {
	if n.right != nil {
		n = n.right.rightmost()
	}
	return n
}

func (n *node) predecessor() *node {
	if n.left != nil {
		return n.left.rightmost()
	}

	y := n.parent
	for y != nil && n == y.left {
		n = y
		y = n.parent
	}
	return y
}

func (t *tree) insertion(n *node) {
	y := t.root
	// Decide the parent
	if y == nil {
		t.root = n
		return
	}
	x := y

	for y != nil {
		x = y
		if n.key < x.key {
			y = y.left
		} else if n.key > x.key {
			y = y.right
		}
	}
	n.parent = x

	// Decide n is to left or right child
	if n.key > x.key {
		x.right = n
	} else {
		x.left = n
	}
}

func (t *tree) deleteNode(k int) {
	n, err := t.root.search(k)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if n.left == nil {
		t.transplant(n, n.right)
	} else if n.right == nil {
		t.transplant(n, n.left)
	} else {
		y := n.right.leftmost()

		if y.parent != n {
			t.transplant(n, y)
			y.right = n.right
			y.right.parent = y
		}

		t.transplant(n, y)
		y.left = n.left
		y.left.parent = y

	}

}

func (t *tree) transplant(old *node, new *node) {
	if old.parent == nil {
		t.root = new
		return
	}

	if old.parent.left == old {
		old.parent.left = new
	} else {
		old.parent.right = new
	}

	new.parent = old.parent
}
