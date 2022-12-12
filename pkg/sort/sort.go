package sort

import (
	"fmt"
	"math/rand"
	"sort"
)

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

// TopoSort create a list with valid topological orderings
func TopoSort(m map[string][]string) ([]string, error) {
	var order []string
	seen := make(map[string]bool)
	var check func(item string)

	var findCycle = func(item, val string) bool {
		if _, ex := m[val]; !ex {
			return false
		}
		var vs []string
		vs = append(vs, m[val]...)
		for _, v := range vs {
			if v == item {
				return true
			}
		}
		return false
	}

	check = func(item string) {

		if !seen[item] {
			seen[item] = true
			if _, ex := m[item]; ex {
				for _, it := range m[item] {
					check(it)
				}
			}
			order = append(order, item)
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		var values []string
		if _, ex := m[key]; ex {
			values = append(values, m[key]...)

			for _, val := range values {
				if findCycle(key, val) {
					return nil, fmt.Errorf("find cycle between %s and %s", key, val)
				}
			}
		}

		check(key)
	}
	return order, nil
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys)
	return order
}
