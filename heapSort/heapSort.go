package main

import (
	"fmt"
	"os"
)

type arr []int

func main() {
	x := arr{78, 84, 85, 83, 84, 77, 79, 77, 77, 89}
	x.buildMaxHeap()
	heapSort(x)
	x.print()
}

func (a arr) parnet(index int) int {
	if index < 0 {
		fmt.Println("Error parant")
		os.Exit(2)
	}
	return a[(index-1)/2]
}

func (a arr) leftChild(index int) int {
	return index*2 + 1
}

func (a arr) rightChild(index int) int {
	return (index + 1) * 2
}

func (a arr) print() {
	for i, item := range a {
		fmt.Println(i, item)
	}
}

func (a arr) buildMinHeap() {
	for i := len(a)/2 - 1; i >= 0; i-- {
		if a.leftChild(i) < len(a) || a.rightChild(i) < len(a) {
			minHeapify(a, i)
		}
	}
}

func (a arr) buildMaxHeap() {
	for i := len(a)/2 - 1; i >= 0; i-- {
		if a.leftChild(i) < len(a) || a.rightChild(i) < len(a) {
			maxHeapify(a, i)
		}
	}
}

func maxHeapify(a arr, i int) {

	leftChild := a.leftChild(i)
	rightChild := a.rightChild(i)
	largest := i

	if leftChild < len(a) {
		if a[leftChild] > a[largest] {
			largest = leftChild
		}
	}

	if rightChild < len(a) {
		if a[rightChild] > a[largest] {
			largest = rightChild
		}
	}

	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		maxHeapify(a, largest)
	}
}

func minHeapify(a arr, i int) {
	leftChild := a.leftChild(i)
	rightChild := a.rightChild(i)
	smallest := i

	if leftChild < len(a) {
		if a[leftChild] < a[smallest] {
			smallest = leftChild
		}
	}

	if rightChild < len(a) {
		if a[rightChild] < a[smallest] {
			smallest = rightChild
		}
	}

	if smallest != i {
		a[i], a[smallest] = a[smallest], a[i]
		minHeapify(a, smallest)
	}

}

func heapSort(a arr) {
	for i := 1; i < len(a); i++ {
		a[0], a[len(a)-i] = a[len(a)-i], a[0]
		a[:len(a)-i].buildMaxHeap()
	}
}
