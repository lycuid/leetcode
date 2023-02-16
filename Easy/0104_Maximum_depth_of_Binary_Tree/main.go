// https://leetcode.com/problems/maximum-depth-of-binary-tree/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + Max(maxDepth(root.Left), maxDepth(root.Right))
}

func main() {}
