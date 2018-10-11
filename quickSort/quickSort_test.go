package main

import "testing"

func TestQuickSort(t *testing.T) {
	x := arr{78, 84, 83, 84, 77, 79, 77, 77, 89}
	x.quickSort(0, len(x)-1)
	for i := range x {
		if i+1 < len(x) {
			if x[i] > x[i+1] {
				t.Errorf("The sorting error occur")
			}
		}
	}
}
