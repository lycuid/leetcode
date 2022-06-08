// https://leetcode.com/problems/validate-binary-search-tree/
package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Aux(root *TreeNode, min, max int) bool {
	return root == nil || (root.Val > min && root.Val < max && Aux(root.Left, min, root.Val) && Aux(root.Right, root.Val, max))
}

func isValidBST(root *TreeNode) bool {
	return Aux(root, math.MinInt, math.MaxInt)
}

func main() {}
