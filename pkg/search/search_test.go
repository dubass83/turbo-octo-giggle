package search

import (
	"encoding/json"
	"testing"
)

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

func TestDijstraAlgo(t *testing.T) {
	var graph = make(map[string]map[string]uint64)
	graph["start"] = make(map[string]uint64)
	graph["start"]["a"] = 6
	graph["start"]["b"] = 2

	graph["a"] = make(map[string]uint64)
	graph["a"]["fin"] = 1

	graph["b"] = make(map[string]uint64)
	graph["b"]["a"] = 3
	graph["b"]["fin"] = 5

	graph["fin"] = make(map[string]uint64)

	var testRes []map[string]uint64
	testRes0 := map[string]uint64{"start": 0}
	testRes1 := map[string]uint64{"b": 2}
	testRes2 := map[string]uint64{"a": 5}
	testRes3 := map[string]uint64{"fin": 6}

	testRes = append(testRes, testRes0, testRes1, testRes2, testRes3)

	res := DijkstraAlgo("start", graph)

	myJSONres, _ := json.MarshalIndent(res, "", "    ")
	myJSONtestRes, _ := json.MarshalIndent(testRes, "", "    ")

	if !compare(res, testRes) {
		t.Errorf("Unexpected result!:\n%s\n\n expected is:\n%s", myJSONres, myJSONtestRes)
	}

}

func compare(foo, bar []map[string]uint64) bool {
	if len(foo) != len(bar) {
		return false
	}
	for i := range foo {
		for key, _ := range foo[i] {
			if foo[i][key] != bar[i][key] {
				return false
			}
		}
	}
	return true
}
