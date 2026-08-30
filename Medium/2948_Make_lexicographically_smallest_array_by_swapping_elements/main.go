// https://leetcode.com/problems/make-lexicographically-smallest-array-by-swapping-elements/
package main

import "sort"

func lexicographicallySmallestArray(nums []int, limit int) []int {
	indices := make([]int, len(nums))
	for i := range indices {
		indices[i] = i
	}
	sort.Slice(indices, func(i, j int) bool {
		return nums[indices[i]] < nums[indices[j]]
	})
	parent := make([]int, len(indices))
	children := make([][]int, len(indices))
	for i := range indices {
		parent[indices[i]] = indices[i]
		if j := i - 1; j >= 0 && nums[indices[j]]+limit >= nums[indices[i]] {
			parent[indices[i]] = parent[indices[j]]
		}
		pi := parent[indices[i]]
		children[pi] = append(children[pi], indices[i])
	}
	for p := range children {
		sort.Ints(children[p])
	}
	res := make([]int, len(nums))
	for i := range indices {
		pi := parent[indices[i]]
		res[children[pi][0]] = nums[indices[i]]
		children[pi] = children[pi][1:]
	}
	return res
}

func main() {}
