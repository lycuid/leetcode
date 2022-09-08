// https://leetcode.com/problems/path-sum/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(r *TreeNode, t int) bool {
	if r == nil {
		return false
	}
	if r.Val == t && r.Left == nil && r.Right == nil {
		return true
	}
	return hasPathSum(r.Left, t-r.Val) || hasPathSum(r.Right, t-r.Val)
}

func main() {}
