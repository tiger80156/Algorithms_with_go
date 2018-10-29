package main

import (
	"testing"
)

func TestRBT_leftRotate(t *testing.T) {

	n1 := node{
		key:   3,
		color: "r",
	}

	n2 := node{
		parent: &n1,
		key:    2,
		color:  "b",
	}

	n3 := node{
		parent: &n1,
		key:    5,
		color:  "b",
	}

	n1.rightChild = &n3
	n1.leftChild = &n2

	t1 := RBT{
		root: &n1,
	}

	t1.leftRotate(&n1)

	if n3.parent != nil {
		t.Errorf("The root parent is wrong")
	}

	if n3.leftChild != &n1 {
		// fmt.Println(n3.leftChild.key)
		t.Errorf("The leftChild is wrong")
	}

	if n3.leftChild.leftChild != &n2 {
		t.Errorf("The leftChild's leftChild is wrong")
	}

	if n3.rightChild != nil {
		t.Errorf("The root right child should be nil")
	}

	if n3.leftChild.rightChild != nil {
		t.Errorf("The leftChild's rightChild should be nil")
	}

	if n3.leftChild.parent != &n3 {
		t.Errorf("The parent is incorrect")
	}

	if n3.leftChild.leftChild.parent != t1.root.leftChild {
		t.Errorf("The parent is incorrect")
	}

}
