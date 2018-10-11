package main

import (
	"fmt"
)

type arr []int

func main() {
	x := arr{78, 84, 85, 83, 84, 77, 79, 77, 77, 89}
	mergeSort(x, 0, len(x)-1)
	x.print()
}

func mergeSort(a arr, font int, end int) {

	if font < end {
		mid := int((font + end) / 2)
		mergeSort(a, font, mid)
		mergeSort(a, mid+1, end)
		merge(a, font, mid, end)
	}

}

func merge(a arr, font int, mid int, end int) {
	mergeArr := a[font : end+1]
	leftSub := make(arr, mid-font+1)
	copy(leftSub, a[font:mid+1])
	leftSub = append(leftSub, 1000000)

	rightSub := make(arr, end-mid)
	copy(rightSub, a[mid+1:end+1])
	rightSub = append(rightSub, 1000000)

	rp := 0
	lp := 0

	for i := 0; i < end+1; i++ {
		if leftSub[lp] > rightSub[rp] {
			mergeArr[i] = rightSub[rp]
			rp++
		} else if rightSub[rp] > leftSub[lp] {
			mergeArr[i] = leftSub[lp]
			lp++
		} else if rightSub[rp] == leftSub[lp] && leftSub[lp] != 1000000 {
			mergeArr[i] = leftSub[lp]
			lp++
		}
	}

}

func (a arr) print() {
	for i, arritem := range a {
		fmt.Println(i, arritem)
	}
}
