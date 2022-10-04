// https://leetcode.com/problems/path-sum/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(r *TreeNode, sum int) bool {
	if r == nil {
		return false
	}
	if r.Left == nil && r.Right == nil {
		return r.Val == sum
	}
	return hasPathSum(r.Left, sum-r.Val) || hasPathSum(r.Right, sum-r.Val)
}

func main() {}
