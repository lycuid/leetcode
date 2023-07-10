// https://leetcode.com/problems/minimum-depth-of-binary-tree/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	if root.Left == nil {
		return 1 + minDepth(root.Right)
	} else if root.Right == nil {
		return 1 + minDepth(root.Left)
	}
	return 1 + Min(minDepth(root.Left), minDepth(root.Right))
}

func main() {}
