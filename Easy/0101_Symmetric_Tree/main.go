// https://leetcode.com/problems/symmetric-tree/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Aux(left, right *TreeNode) bool {
	return (left == nil && right == nil) ||
		(left != nil && right != nil &&
			left.Val == right.Val &&
			Aux(left.Left, right.Right) && Aux(left.Right, right.Left))
}

func isSymmetric(root *TreeNode) bool {
	return root == nil || Aux(root.Left, root.Right)
}

func main() {}
