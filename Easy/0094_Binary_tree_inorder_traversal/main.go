// https://leetcode.com/problems/binary-tree-inorder-traversal/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) (ret []int) {
	if root != nil {
		ret = append(ret, inorderTraversal(root.Left)...)
		ret = append(ret, root.Val)
		ret = append(ret, inorderTraversal(root.Right)...)
	}
	return ret
}

func main() {}
