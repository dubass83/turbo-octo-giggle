package search

import (
	"fmt"
	"sort"
	"math"
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

func BFS(start string, graph map[string][]string, check func(name string) bool) bool {
	var searchQueue []string
	searchQueue = append(searchQueue, graph[start]...)
	var searched []string
	var person string
	for len(searchQueue) != 0 {
		person = searchQueue[0]
		searchQueue = searchQueue[1:]
		if personNotInSearched(person, searched) {
			if check(person) {
				fmt.Println("find: ", person)
				return true
			}
			searchQueue = append(searchQueue, graph[person]...)
			searched = append(searched, person)
		}
	}
	return false
}

func personNotInSearched(name string, serched []string) bool {
	for _, n := range serched {
		if n == name {
			return false
		}
	}
	return true
}

func DijkstraAlgo(start string, graph map[string]map[string]int64) map[int]map[string]int64 {

}

initCosts(start string, graph map[string]map[string]int64) map[string]int64{
  costs = make(map[string]int64)

  costs[start] = 0
  for key := range graph{
	if key == start {
		continue
	}
	costs[key] = math.Inf(1)
  }
  return costs
}

func contains(name string, s []string) bool {
	for _, n := range s {
		if n == name {
			return true
		}
		return false
	}
}
