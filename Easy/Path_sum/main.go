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

	left := r.Left != nil && hasPathSum(r.Left, t-r.Val)
	right := r.Right != nil && hasPathSum(r.Right, t-r.Val)

	return left || right
}

func main() {}
