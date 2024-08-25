// https://leetcode.com/problems/binary-tree-postorder-traversal/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) (ret []int) {
	if root != nil {
		ret = append(ret, postorderTraversal(root.Left)...)
		ret = append(ret, postorderTraversal(root.Right)...)
		ret = append(ret, root.Val)
	}
	return ret
}

func main() {}
