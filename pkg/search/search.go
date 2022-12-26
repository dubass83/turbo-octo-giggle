package search

import (
	"fmt"
	"sort"
)

func BinaryInt(sl []int, val int) (find int, err error) {
	sort.Ints(sl)
	low := 0
	high := len(sl) - 1
	if sl[low] > val || sl[high] < val {
		err = fmt.Errorf("no such value in the list: %d", val)
		return -1, err
	}
	for low <= high {
		mid := (low + high) / 2
		if sl[mid] == val {
			find = mid
			return find, nil
		}
		if sl[mid] < val {
			low = mid + 1
			continue
		}
		if sl[mid] > val {
			high = mid - 1
		}
	}
	err = fmt.Errorf("no such value in the list: %d", val)
	return -1, err
}
