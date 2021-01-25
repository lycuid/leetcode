// https://leetcode.com/problems/symmetric-tree/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isMirrored(left, right *TreeNode) bool {
	if (left == nil && right != nil) || (left != nil && right == nil) {
		return false
	}
	if left == nil && right == nil {
		return true
	}
	return left.Val == right.Val && isMirrored(left.Left, right.Right) && isMirrored(left.Right, right.Left)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirrored(root.Left, root.Right)
}

func main() {}
