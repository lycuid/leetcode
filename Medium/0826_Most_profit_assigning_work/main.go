// https://leetcode.com/problems/most-profit-assigning-work/
package main

import "sort"

func maxProfitAssignment(difficulty []int, profit []int, worker []int) (max_profit int) {
	indices := make([]int, len(difficulty))
	for i := range indices {
		indices[i] = i
	}
	sort.Slice(indices, func(i, j int) bool {
		return profit[indices[i]] > profit[indices[j]]
	})
	sort.Slice(worker, func(i, j int) bool {
		return worker[i] > worker[j]
	})
	for _, ability := range worker {
		for len(indices) > 0 && difficulty[indices[0]] > ability {
			indices = indices[1:]
		}
		if len(indices) == 0 {
			break
		}
		max_profit += profit[indices[0]]
	}
	return max_profit
}

func main() {}
