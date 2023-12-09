// https://leetcode.com/problems/binary-tree-inorder-traversal/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) (inorder []int) {
	if root != nil {
		inorder = append(inorder, inorderTraversal(root.Left)...)
		inorder = append(inorder, root.Val)
		inorder = append(inorder, inorderTraversal(root.Right)...)
	}
	return inorder
}

func main() {}
