package main

import (
	"fmt"
)

type node struct {
	parent     *node
	leftChild  *node
	rightChild *node
	key        int
	color      string
}

type RBT struct {
	root *node
	neer *node
}

func main() {

	// n1.rightChild = &n3
	// n1.leftChild = &n2

	neer := &node{
		color: "b",
	}

	t1 := RBT{
		neer: neer,
		root: neer,
	}

	t1.insertNode(1)
	t1.insertNode(3)
	t1.insertNode(6)
	t1.insertNode(4)
	t1.insertNode(2)
	t1.traverse(t1.root)

	// t1.root.printTree()
	// t1.root.leftChild.printTree()
	// t1.root.rightChild.printTree
	// t1.root.leftChild.printTree()

}

func (t *RBT) leftRotate(n *node) {
	y := n.rightChild

	// Step One: Set the leftChild of y to n right child
	n.rightChild = y.leftChild
	if y.leftChild != t.neer {
		y.leftChild.parent = n
	}

	// Step two: Set node y parent to node x parent
	// Node n is root of tree
	// Set node y to root
	y.parent = n.parent

	if n == t.root {
		t.root = y
	} else {
		if n == n.parent.leftChild {
			n.parent.leftChild = y
		} else {
			n.parent.rightChild = y
		}
	}

	// Step 3 : Set	t1.root.printTree()
	// the n parent	t1.root.leftChild.printTree()
	y.leftChild = n
	n.parent = y
}

func (t *RBT) rightRotate(n *node) {
	y := n.leftChild

	// Step One: Set n left child to y's right child
	n.leftChild = y.rightChild
	if y.rightChild != t.neer {
		y.rightChild.parent = n
	}

	//Step two: Set y to n parent
	y.parent = n.parent

	if n == t.root {
		t.root = y
	} else {
		//Step three: Set n parent node's child to y
		if n == n.parent.leftChild {
			n.parent.leftChild = y
		} else {
			n.parent.rightChild = y
		}
	}
	// Step four: Set n to y right child
	y.rightChild = n
	n.parent = y
}

func (t *RBT) traverse(n *node) {

	if n == t.neer {
		return
	}

	if n.rightChild != t.neer {
		t.traverse(n.rightChild)
	}

	fmt.Println(n.key, n.color)

	if n.leftChild != t.neer {
		t.traverse(n.leftChild)
	}
}

func (n *node) printTree() {
	fmt.Println("LeftChild", n.leftChild)
	fmt.Println("rightChild", n.rightChild)
	fmt.Println("parent", n.parent)
	fmt.Println("values", n.key)
	println()
}

func (t *RBT) insertNode(k int) {
	n := &node{
		key:        k,
		color:      "r",
		rightChild: t.neer,
		leftChild:  t.neer,
	}

	if t.root == t.neer {
		t.root = n
		t.root.parent = t.neer
		t.root.color = "b"
		return
	}

	x := t.root
	y := x

	for x != t.neer {
		y = x
		if n.key > x.key {
			x = x.rightChild
		} else {
			x = x.leftChild
		}
	}

	n.parent = y

	if n.key > y.key {
		y.rightChild = n
	} else {
		y.leftChild = n
	}

	t.colorFixUp(n)
}

func (t *RBT) colorFixUp(n *node) {

	//Case 1 n's parent has color conflict with n
	for n.parent.color == "r" {
		p := n.parent
		a := n.parent.parent

		if n == p.leftChild {
			u := n.parent.parent.rightChild

			if u.color == "r" {
				p.color = "b"
				u.color = "b"
				a.color = "r"

				n = n.parent.parent
			} else {
				// Case 2
				if n == p.rightChild {
					n = p
					t.leftRotate(n)
				}

				// Case 3
				p.color = "b"
				a.color = "r"
				t.rightRotate(a)
			}
		} else {
			u := n.parent.parent.leftChild

			if u.color == "r" {
				p.color = "b"
				u.color = "b"
				a.color = "r"

				n = n.parent.parent
			} else {
				// Case 2
				if n == p.leftChild {
					n = p
					t.rightRotate(n)
				}
				p.color = "b"
				a.color = "r"
				t.leftRotate(a)

			}

		}
	}

}

func (n *node) searchNode(k int) *node {

	if n.leftChild == nil && n.rightChild == nil {
		return nil
	}

	if k > n.key {
		n.leftChild.searchNode(k)
	} else if k < n.key {
		n.rightChild.searchNode(k)
	}
	return n
}
