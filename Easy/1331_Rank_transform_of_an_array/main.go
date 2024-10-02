// https://leetcode.com/problems/rank-transform-of-an-array/
package main

import "sort"

func arrayRankTransform(arr []int) (rank []int) {
	if n := len(arr); n > 0 {
		rank, indices := make([]int, n), make([]int, n)
		for i := range indices {
			indices[i] = i
		}
		sort.Slice(indices, func(i, j int) bool {
			if arr[indices[i]] == arr[indices[j]] {
				return indices[i] < indices[j]
			}
			return arr[indices[i]] < arr[indices[j]]
		})
		rank[indices[0]] = 1
		for i := 1; i < n; i++ {
			rank[indices[i]] = rank[indices[i-1]]
			if arr[indices[i]] != arr[indices[i-1]] {
				rank[indices[i]]++
			}
		}
		return rank
	}
	return nil
}

func main() {}
