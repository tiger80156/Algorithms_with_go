package main

import (
	"fmt"
	"testing"
)

func TestBuildMaxHeap(t *testing.T) {
	x := arr{78, 84, 85, 84, 84, 77, 79, 77, 89}
	x.buildMaxHeap()
	for i := 0; i < len(x); i++ {
		left := x.leftChild(i)
		right := x.rightChild(i)

		if left < len(x) {
			if x[i] < x[left] {
				fmt.Println(i)
				x.print()
				t.Errorf("The max heap tree build fail")
			}
		} else if right < len(x) {
			if x[i] < x[right] {
				x.print()
				t.Errorf("The max heap tree build fail")
			}
		}
	}
}

func TestHeapSort(t *testing.T) {
	x := arr{78, 84, 85, 84, 77}
	x.buildMaxHeap()
	heapSort(x)
	for i := range x {
		if i < len(x)-1 {
			if x[i] > x[i+1] {
				x.print()
				t.Errorf("The sorting result is fault.")
			}
		}
	}
}
