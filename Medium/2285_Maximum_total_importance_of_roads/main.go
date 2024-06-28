// https://leetcode.com/problems/maximum-total-importance-of-roads/
package main

import "sort"

func maximumImportance(n int, roads [][]int) (ret int64) {
	degree := make([]int, n)
	for _, edge := range roads {
		degree[edge[0]]++
		degree[edge[1]]++
	}
	sort.Slice(degree, func(i, j int) bool {
		return degree[i] > degree[j]
	})
	for i, deg := range degree {
		ret += int64((n - i) * deg)
	}
	return ret
}

func main() {}
