package sort

import (
	"fmt"
	"testing"
)

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

func TestTopoSort(t *testing.T) {
	var prereqs = map[string][]string{
		"algorithms": {"data structures"},
		"calculus":   {"linear algebra"},
		"compilers": {
			"data structures",
			"formal languages",
			"computer organization",
		},
		"data structures":       {"discrete math"},
		"databases":             {"data structures"},
		"discrete math":         {"intro to programming"},
		"formal languages":      {"discrete math"},
		"networks":              {"operating systems"},
		"operating systems":     {"data structures", "computer organization"},
		"programming languages": {"data structures", "computer organization"},
	}

	var expectedResult = []string{
		"intro to programming",
		"discrete math",
		"data structures",
		"algorithms",
		"linear algebra",
		"calculus",
		"formal languages",
		"computer organization",
		"compilers",
		"databases",
		"operating systems",
		"networks",
		"programming languages",
	}

	actualResult := topoSort(prereqs)
	if !compareSlice(actualResult, expectedResult) {
		for i, val := range actualResult {
			fmt.Printf("%d:\t%s\n", i+1, val)
		}
		t.Errorf("Slice is not sorted well!\n"+
			"actualResult:\t%v\nexpectedResult:\t%v", actualResult, expectedResult)
	}

	actualResult = TopoSort(prereqs)
	if !compareSlice(actualResult, expectedResult) {
		for i, val := range actualResult {
			fmt.Printf("%d:\t%s\n", i+1, val)
		}
		t.Errorf("Slice is not sorted well!\n"+
			"actualResult:\t%v\nexpectedResult:\t%v", actualResult, expectedResult)
	}

}

func compareSlice[V int | string](foo, bar []V) bool {
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
