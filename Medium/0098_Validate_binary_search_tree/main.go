// https://leetcode.com/problems/validate-binary-search-tree/
package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Solve(root *TreeNode, min, max int) bool {
	return root == nil || root.Val > min && root.Val < max &&
		Solve(root.Left, min, root.Val) && Solve(root.Right, root.Val, max)
}

func isValidBST(root *TreeNode) bool {
	return Solve(root, math.MinInt, math.MaxInt)
}

func main() {}
