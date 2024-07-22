// https://leetcode.com/problems/sort-the-people/
package main

import "sort"

func sortPeople(names []string, heights []int) []string {
	indices, ret := make([]int, len(names)), make([]string, len(names))
	for i := range indices {
		indices[i] = i
	}
	sort.Slice(indices, func(i, j int) bool {
		return heights[indices[i]] > heights[indices[j]]
	})
	for i, index := range indices {
		ret[i] = names[index]
	}
	return ret
}

func main() {}
