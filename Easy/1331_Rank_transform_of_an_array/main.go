// https://leetcode.com/problems/rank-transform-of-an-array/
package main

import "sort"

func arrayRankTransform(arr []int) (res []int) {
	if n := len(arr); n > 0 {
		indices := make([]int, n)
		for i := range indices {
			indices[i] = i
		}
		sort.Slice(indices, func(i, j int) bool {
			return arr[indices[i]] < arr[indices[j]]
		})
		res = make([]int, n)
		res[indices[0]] = 1
		for i := 1; i < n; i++ {
			res[indices[i]] = res[indices[i-1]]
			if arr[indices[i]] != arr[indices[i-1]] {
				res[indices[i]]++
			}
		}
	}
	return res
}

func main() {}
