// https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func solve(nums *[]int, start, end int) *TreeNode {
	if end < start {
		return nil
	}

	var left, right *TreeNode
	pivot := ((end - start) / 2) + ((end - start) % 2) + start

	if pivot > start {
		left = solve(nums, start, pivot-1)
	}
	if pivot < end {
		right = solve(nums, pivot+1, end)
	}

	return &TreeNode{Val: (*nums)[pivot], Left: left, Right: right}
}

func sortedArrayToBST(nums []int) *TreeNode {
	return solve(&nums, 0, len(nums)-1)
}

func main() {}
