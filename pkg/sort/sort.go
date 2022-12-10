package sort

import "math/rand"

// QuikSort function is using an algorithm for sorting the elements
// of a collection in an organized way.
func QuikSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	left, right := 0, len(a)-1
	pivot := rand.Int() % len(a)
	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}
	a[left], a[right] = a[right], a[left]

	QuikSort(a[:left])
	QuikSort(a[left+1:])
	return a
}
