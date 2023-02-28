package search

import (
	"fmt"
	"math"
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

func DijkstraAlgo(start string, graph map[string]map[string]uint64) map[int]map[string]uint64 {

	processed := make([]string, 0)

	costs := initCosts(start, graph)
	parents := initParents(start, graph)

	node := findLowestCostNode(costs, processed)
}

func initCosts(start string, graph map[string]map[string]uint64) map[string]uint64 {
	costs := make(map[string]uint64)

	costs[start] = 0
	for key := range graph {
		if key == start {
			continue
		}
		costs[key] = math.MaxUint64
	}
	return costs
}

func initParents(start string, graph map[string]map[string]uint64) map[string]string {
	parents := make(map[string]string)
	for key := range graph {
		if key == start {
			continue
		}
		parents[key] = ""
	}
	return parents
}

func findLowestCostNode(costs map[string]uint64, processed []string) string {
	var lowestCost uint64
	lowestCost = math.MaxUint64
	lowestCostNode := ""

	for node, _ := range costs {
		cost := costs[node]
		if cost < lowestCost && !contains(node, processed) {
			lowestCost = cost
			lowestCostNode = node
		}
	}
	return lowestCostNode
}

func contains(name string, s []string) bool {
	for _, n := range s {
		if n == name {
			return true
		}
	}
	return false
}
