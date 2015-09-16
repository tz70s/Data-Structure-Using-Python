package Sorting

/*
func swap(a *int, b *int) (*int, *int) {
	*a, *b = *b, *a
	return b, a
}
*/

func Newslice(a []int) []int {
	newslice := make([]int, len(a))
	for k, v := range a {
		newslice[k] = v
	}
	return newslice
}

func Compare(a []int, b []int) bool {
	if len(a) == len(b) {
		for i := 0; i < len(a); i++ {
			if a[i] == b[i] {
				continue
			} else {
				return false
				break
			}
		}
		return true

	} else {
		return false
	}
}
