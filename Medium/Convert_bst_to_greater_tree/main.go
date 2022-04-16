// https://leetcode.com/problems/convert-bst-to-greater-tree/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convert(root *TreeNode, base int) int {
	if root == nil {
		return 0
	}
	root.Val += base + convert(root.Right, base)
	return root.Val - base + convert(root.Left, root.Val)
}

func convertBST(root *TreeNode) *TreeNode {
	convert(root, 0)
	return root
}

func main() {}
