package Sorting

func merge(list []int, low, high int) {
	mid := (high + low) / 2
	// Creat new array
	Llist := make([]int, mid-low+1)
	Rlist := make([]int, high-mid)

	for i := 0; i < len(Llist); i++ {
		Llist[i] = list[low+i]
	}
	for i := 0; i < len(Rlist); i++ {
		Rlist[i] = list[i+mid+1]
	}

	l, r := 0, 0
	k := low

	for l < len(Llist) && r < len(Rlist) {
		if Llist[l] <= Rlist[r] {
			list[k] = Llist[l]
			l++
		} else {
			list[k] = Rlist[r]
			r++
		}
		k++
	}

	for r < len(Rlist) {
		list[k] = Rlist[r]
		r++
		k++
	}

	for l < len(Llist) {
		list[k] = Llist[l]
		l++
		k++
	}

}

func MergeSort(list []int, low int, high int) {
	if low < high {
		mid := (high + low) / 2
		MergeSort(list, low, mid)
		MergeSort(list, mid+1, high)
		merge(list, low, high)
	}
}
