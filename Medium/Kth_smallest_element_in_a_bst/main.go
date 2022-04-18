// https://leetcode.com/problems/kth-smallest-element-in-a-bst/
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

func kthSmallest(root *TreeNode, k int) int {
	new_root := increasingBST(root)
	for ; k > 1 && new_root != nil; k-- {
		new_root = new_root.Right
	}
	if new_root == nil {
		return 0
	}
	return new_root.Val
}

func main() {}
