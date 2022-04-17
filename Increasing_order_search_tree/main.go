// https://leetcode.com/problems/increasing-order-search-tree/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	var old_root, r *TreeNode
	if root == nil {
		goto EXIT
	}
	root.Right = increasingBST(root.Right)
	if root.Left == nil {
		goto EXIT
	}
	old_root = root
	root = increasingBST(root.Left)
	old_root.Left = nil
	for r = root; r.Right != nil; r = r.Right {
	}
	r.Right = old_root
EXIT:
	return root
}

func main() {}
