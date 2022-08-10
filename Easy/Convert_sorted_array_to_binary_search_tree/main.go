// https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/
package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func sortedArrayToBST(nums []int) (root *TreeNode) {
	if mid := len(nums) / 2; len(nums) > 0 {
		root = &TreeNode{
			Val:   nums[mid],
			Left:  sortedArrayToBST(nums[:mid]),
			Right: sortedArrayToBST(nums[mid+1:]),
		}
	}
	return root
}

func main() {}
