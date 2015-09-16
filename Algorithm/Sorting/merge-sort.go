package Sorting

func merge(low, high int) {

}

func MergeSort(low, high int) {
	if low < high {
		mid := (high + low) / 2
		MergeSort(low, mid)
		MergeSort(mid+1, high)
		merge(low, high)
	}
}
