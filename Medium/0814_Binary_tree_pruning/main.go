// https://leetcode.com/problems/binary-tree-pruning/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pruneTree(root *TreeNode) *TreeNode {
	if root != nil {
		root.Left, root.Right = pruneTree(root.Left), pruneTree(root.Right)
		if root.Val != 1 && root.Left == nil && root.Right == nil {
			return nil
		}
	}
	return root
}

func main() {}
