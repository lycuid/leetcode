// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root != nil && root != p && root != q {
		left := lowestCommonAncestor(root.Left, p, q)
		right := lowestCommonAncestor(root.Right, p, q)
		if left == nil && right != nil {
			return right
		}
		if right == nil && left != nil {
			return left
		}
		if left == nil && right == nil {
			return nil
		}
	}
	return root
}

func main() {}
