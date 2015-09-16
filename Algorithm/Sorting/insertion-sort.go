package Sorting

func InsertionSort(list []int) []int {

	for newi := 1; newi < len(list); newi++ {
		value := list[newi]
		for move := newi - 1; move >= 0; {
			if value >= list[move] {
				list[move+1] = value
				break
			} else {
				list[move+1] = list[move]
				move--
			}
		}
	}
	return list
}
