package main

import (
	"fmt"
)

type arr []int

func main() {
	x := arr{9, 4, 1, 6, 7, 3, 8, 2, 5}
	x.quickSort(0, len(x)-1)
	fmt.Println(x)
}

func (a arr) quickSort(font int, end int) {
	if font < end {
		mid := a.partition(font, end)
		a.quickSort(font, mid-1)
		a.quickSort(mid+1, end)
	}

}

func (a arr) partition(font int, end int) int {
	pivotIndex := end
	curr := font - 1

	for i := font; i < end; i++ {
		if a[i] < a[pivotIndex] {
			curr++
			a[curr], a[i] = a[i], a[curr]
		}
	}
	curr++
	a[curr], a[pivotIndex] = a[pivotIndex], a[curr]
	return curr
}
