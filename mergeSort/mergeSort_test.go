package main

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	x := arr{4, 3, 4, 1}

	mergeSort(x, 0, len(x)-1)

	for i := 0; i < len(x)-2; i++ {
		if x[i] > x[i+1] {
			fmt.Println(x)
			t.Errorf("The sorting incorrect")
		}
	}

}
func TestMergeSort_withPossitiveNumber(t *testing.T) {
	x := arr{-2, 3, -2, 1, 12, 9, 70, -20, -3, -1, 0, -17}

	mergeSort(x, 0, len(x)-1)

	for i := 0; i < len(x)-2; i++ {
		if x[i] > x[i+1] {
			fmt.Println(x)
			t.Errorf("The sorting incorrect")
		}
	}

}

func TestMergeSort_withRepeatNumber(t *testing.T) {
	x := arr{1, 3, 1, 4, 6, 1, 3, 2}

	mergeSort(x, 0, len(x)-1)
	for i := 0; i < len(x)-2; i++ {
		if x[i] > x[i+1] {
			fmt.Println(x)
			t.Errorf("The sorting incorrect")
		}
	}
}
