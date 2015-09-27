package Sorting

import (
	"fmt"
)

func merge(list []int, low, high int) {
	mid := (high + low) / 2
	// Creat new array
	Llist := make([]int, mid-low+2)
	Rlist := make([]int, high-mid+1)

	for i := 0; i < len(Llist); i++ {
		Llist[i] = list[low]
		//fmt.Print(Llist[i])
	}
	fmt.Println()
	for i := 0; i < len(Rlist); i++ {
		Rlist[i] = list[i+mid]
		//fmt.Print(Rlist[i])
	}

	l, r := 0, 0

	total := high - low + 1
	k := 0

	for k < total {
		if Llist[l] <= Rlist[r] && l < len(Llist) && r < len(Rlist) {
			list[k] = Llist[l]
			l += 1
			k += 1
		} else if Rlist[r] < Llist[l] && r < len(Rlist) && l < len(Llist) {
			list[k] = Rlist[r]
			r += 1
			k += 1
		} else if l == len(Llist) {
			list[k] = Rlist[r]
			r += 1
			k += 1
		} else if r == len(Rlist) {
			list[k] = Llist[l]
			l += 1
			k += 1
		}
	}

}

func MergeSort(list []int, low int, high int) {
	if low < high {
		mid := (high + low) / 2
		MergeSort(list, low, mid)
		//fmt.Println(list[low : mid+1])
		MergeSort(list, mid+1, high)
		//fmt.Println(list[mid+1 : high+1])
		merge(list, low, high)
	}
}
