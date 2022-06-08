// https://leetcode.com/problems/binary-tree-inorder-traversal/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorder(root *TreeNode, arr *[]int) {
	if root != nil {
		inorder(root.Left, arr)
		(*arr) = append(*arr, root.Val)
		inorder(root.Right, arr)
	}
}

func inorderTraversal(root *TreeNode) (ret []int) {
	inorder(root, &ret)
	return ret
}

func main() {}
