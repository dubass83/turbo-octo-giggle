package search

import "testing"

func TestBinaryInt(t *testing.T) {
	givenList := []int{-1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	expectedResult := 7
	actualResult, err := BinaryInt(givenList, 7)
	if err != nil {
		t.Errorf("this is unxpected error: %v", err)
	}
	if expectedResult != actualResult {
		t.Errorf("expectedResult: %d not equil actualResult: %d", expectedResult, actualResult)
	}
}

func TestBFS(t *testing.T) {
	var graph = make(map[string][]string)
	graph["you"] = []string{"alice", "bob", "claire"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"thom", "jonny"}
	graph["anuj"] = []string{}
	graph["peggy"] = []string{}
	graph["thom"] = []string{}
	graph["jonny"] = []string{}
	checkPersonM := func(name string) bool {
		return name[len(name)-1] == 'm'
	}
	checkPersonX := func(name string) bool {
		return name[len(name)-1] == 'x'
	}
	if !BFS("you", graph, checkPersonM) {
		t.Errorf("unexpected")
	}
	if BFS("you", graph, checkPersonX) {
		t.Errorf("expected")
	}
}
