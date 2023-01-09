// https://leetcode.com/problems/binary-tree-preorder-traversal/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) (ret []int) {
	if root != nil {
		ret = append(ret, root.Val)
		ret = append(ret, preorderTraversal(root.Left)...)
		ret = append(ret, preorderTraversal(root.Right)...)
	}
	return ret
}

func main() {}
