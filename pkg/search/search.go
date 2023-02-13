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

func BFS(start string, graph map[string][]string, check func(name string) bool) bool {
	var search_queue []string
	search_queue = append(search_queue, graph[start]...)
	var searched []string
	var person string
	for len(search_queue) != 0 {
		person = search_queue[0]
		search_queue = search_queue[1:]
		if personNotInSearched(person, searched) {
			if check(person) {
				fmt.Println("find: ", person)
				return true
			}
			search_queue = append(search_queue, graph[person]...)
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
