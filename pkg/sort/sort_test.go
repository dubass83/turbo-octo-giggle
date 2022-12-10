package sort

import "testing"

func TestQuikSort(t *testing.T) {
	testSlice := []int{9, 1, 2, 4, 5, 3, 6, 8, 7}
	actualSlice := QuikSort(testSlice)
	expectedSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	test2Slice := []int{9, 1, 2, 4, 5, 3, 6, 8, 7}

	if !compareSlice(actualSlice, expectedSlice) {
		t.Errorf("Slice is not sorted well!\n"+
			"actualSlice:\t%v\nexpectedSlice:\t%v", actualSlice, expectedSlice)
	}

	if compareSlice(actualSlice, test2Slice) {
		t.Errorf("Something wrong with compareSlice function\n"+
			"actualSlice:\t%v\ntestSlice:\t%v", actualSlice, testSlice)
	}
}

func compareSlice(foo, bar []int) bool {
	if len(foo) != len(bar) {
		return false
	}
	for i := range foo {
		if foo[i] != bar[i] {
			return false
		}
	}
	return true
}
