package Sorting

import (
	"fmt"
)

type List struct {
	list []int
}

func merge(Llist, Rlist []int) (tmplist []int) {
	length := len(Llist) + len(Rlist)
	tmplist = make([]int, length)
	lptr := 0
	rptr := 0
	count := 0
	for count < length {
		if Llist[lptr] <= Rlist[rptr] {
			tmplist[count] = Llist[lptr]
			count++
			lptr++
		} else {
			tmplist[count] = Rlist[rptr]
			count++
			rptr++
		}
	}

	return tmplist

}

func (alist *List) MergeSort(low int, high int) {
	rtlist := make([]int, high)
	if low < high {
		mid := (high + low) / 2
		alist.MergeSort(low, mid)
		Llist := alist.list[low : mid+1]
		alist.MergeSort(mid+1, high)
		Rlist := alist.list[mid+1 : high+1]
		rtlist = merge(Llist, Rlist)
	}

	if len(rtlist) == len(alist.list) {
		fmt.Println(rtlist)
	}

}
