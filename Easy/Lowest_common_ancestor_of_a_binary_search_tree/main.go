// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root != nil && root.Val != p.Val && root.Val != q.Val {
		left := lowestCommonAncestor(root.Left, p, q)
		right := lowestCommonAncestor(root.Right, p, q)
		if left == nil && right == nil {
			return nil
		} else if left != nil && right == nil {
			return left
		} else if left == nil && right != nil {
			return right
		}
	}
	return root
}

func main() {}
